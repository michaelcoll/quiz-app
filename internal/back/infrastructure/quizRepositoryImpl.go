/*
 * Copyright (c) 2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
	pm "github.com/michaelcoll/quiz-app/internal/back/presentation"
)

type QuizDBRepository struct {
	domain.QuizRepository

	w *ConnectionWrapper
}

func NewQuizRepository(w *ConnectionWrapper) *QuizDBRepository {
	return &QuizDBRepository{w: w}
}

func isAdmin(userId string) bool {
	return userId == ""
}

func (r *QuizDBRepository) FindFullBySha1(ctx context.Context, sha1 string, userId string) (*domain.Quiz, error) {
	entities, err := r.w.queries().FindQuizFullBySha1(ctx, sqlc.FindQuizFullBySha1Params{
		Sha1: sha1,
		ID:   userId,
	})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, pm.Errorf(http.StatusNotFound, "quiz with sha1: %s was not found.", sha1)
	}

	quiz := domain.Quiz{}

	for _, entity := range entities {
		if quiz.Sha1 == "" {
			quiz.Sha1 = entity.QuizSha1
			quiz.Filename = entity.QuizFilename
			quiz.Name = entity.QuizName
			quiz.Active = entity.QuizActive
			quiz.Version = entity.QuizVersion
			quiz.Duration = entity.QuizDuration
			quiz.CreatedAt = entity.QuizCreatedAt
			quiz.Questions = map[string]domain.QuizQuestion{}
		}

		if _, found := quiz.Questions[entity.QuestionSha1]; !found {
			newQuestion := domain.QuizQuestion{
				Sha1:         entity.QuestionSha1,
				Position:     entity.QuestionPosition,
				Content:      entity.QuestionContent,
				Code:         entity.QuestionCode.String,
				CodeLanguage: entity.QuestionCodeLanguage.String,
				Answers:      map[string]domain.QuizQuestionAnswer{},
			}
			quiz.Questions[entity.QuestionSha1] = newQuestion
		}

		quiz.Questions[entity.QuestionSha1].Answers[entity.AnswerSha1] = domain.QuizQuestionAnswer{
			Sha1:    entity.AnswerSha1,
			Content: entity.AnswerContent,
			Valid:   entity.AnswerValid,
		}
	}

	return &quiz, nil
}

func (r *QuizDBRepository) FindAllActive(ctx context.Context, userId string, limit uint16, offset uint16) ([]*domain.Quiz, error) {
	entities, err := r.w.queries().FindAllActiveQuiz(ctx, sqlc.FindAllActiveQuizParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	domainsMap := make(map[string]*domain.Quiz)

	for _, entity := range entities {
		if _, found := domainsMap[entity.Sha1]; !found {
			domainsMap[entity.Sha1] = &domain.Quiz{
				Sha1:     entity.Sha1,
				Name:     entity.Name,
				Duration: entity.Duration,
				Classes:  map[uuid.UUID]string{},
			}
		}

		if entity.ClassName != "" {
			domainsMap[entity.Sha1].Classes[entity.ClassUuid] = entity.ClassName
		}

		if isAdmin(userId) {
			domainsMap[entity.Sha1].Filename = entity.Filename
			domainsMap[entity.Sha1].Version = entity.Version
			domainsMap[entity.Sha1].Active = entity.Active
			domainsMap[entity.Sha1].CreatedAt = entity.CreatedAt
		}
	}

	var domains []*domain.Quiz

	for _, d := range domainsMap {
		domains = append(domains, d)
	}

	return domains, nil
}

func (r *QuizDBRepository) CountAllActive(ctx context.Context, userId string) (uint32, error) {

	if isAdmin(userId) {
		count, err := r.w.queries().CountAllActiveQuiz(ctx)
		if err != nil {
			return 0, err
		}

		return uint32(count), nil
	}

	count, err := r.w.queries().CountAllActiveQuizForUser(ctx, userId)
	if err != nil {
		return 0, err
	}

	return uint32(count), nil

}

func (r *QuizDBRepository) FindLatestVersionByFilename(ctx context.Context, filename string) (*domain.Quiz, error) {

	quiz, err := r.w.queries().FindQuizByFilenameAndLatestVersion(ctx, filename)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return r.toQuiz(quiz), nil
}

func (r *QuizDBRepository) Create(ctx context.Context, quiz *domain.Quiz) error {

	err := r.w.queries().CreateOrReplaceQuiz(ctx, sqlc.CreateOrReplaceQuizParams{
		Sha1:      quiz.Sha1,
		Name:      quiz.Name,
		Filename:  quiz.Filename,
		Version:   quiz.Version,
		Duration:  quiz.Duration,
		CreatedAt: quiz.CreatedAt,
	})
	if err != nil {
		return err
	}

	for _, question := range quiz.Questions {
		err := r.w.queries().CreateOrReplaceQuestion(ctx, sqlc.CreateOrReplaceQuestionParams{
			Sha1:         question.Sha1,
			Position:     int64(question.Position),
			Content:      question.Content,
			Code:         sql.NullString{String: question.Code, Valid: true},
			CodeLanguage: sql.NullString{String: question.CodeLanguage, Valid: true},
		})
		if err != nil {
			return err
		}

		err = r.w.queries().LinkQuestion(ctx, sqlc.LinkQuestionParams{
			QuizSha1:     quiz.Sha1,
			QuestionSha1: question.Sha1,
		})
		if err != nil {
			return err
		}

		for _, answer := range question.Answers {
			err := r.w.queries().CreateOrReplaceAnswer(ctx, sqlc.CreateOrReplaceAnswerParams{
				Sha1:    answer.Sha1,
				Content: answer.Content,
				Valid:   answer.Valid,
			})
			if err != nil {
				return err
			}

			err = r.w.queries().LinkAnswer(ctx, sqlc.LinkAnswerParams{
				QuestionSha1: question.Sha1,
				AnswerSha1:   answer.Sha1,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *QuizDBRepository) ActivateOnlyVersion(ctx context.Context, filename string, version int) error {
	err := r.w.queries().ActivateOnlyVersion(ctx, sqlc.ActivateOnlyVersionParams{
		Filename: filename,
		Version:  version,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *QuizDBRepository) FindAllSessions(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16) ([]*domain.Session, error) {
	if isAdmin(userId) {
		sessions, err := r.w.queries().FindAllSessions(ctx, sqlc.FindAllSessionsParams{
			QuizActive: quizActive,
			Limit:      int64(limit),
			Offset:     int64(offset),
		})
		if err != nil {
			return nil, err
		}

		return r.toSessionArray(sessions), nil
	}

	sessions, err := r.w.queries().FindAllSessionsForUser(ctx, sqlc.FindAllSessionsForUserParams{
		QuizActive: quizActive,
		UserID:     userId,
		Limit:      int64(limit),
		Offset:     int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toSessionArray(sessions), nil
}

func (r *QuizDBRepository) CountAllSessions(ctx context.Context, quizActive bool, userId string) (uint32, error) {
	if isAdmin(userId) {
		count, err := r.w.queries().CountAllSessions(ctx, quizActive)
		if err != nil {
			return 0, err
		}

		return uint32(count), nil
	}

	count, err := r.w.queries().CountAllSessionsForUser(ctx, sqlc.CountAllSessionsForUserParams{
		QuizActive: quizActive,
		UserID:     userId,
	})
	if err != nil {
		return 0, err
	}

	return uint32(count), nil
}

func (r *QuizDBRepository) StartSession(ctx context.Context, userId string, quizSha1 string) (uuid.UUID, error) {
	sessionUuid := uuid.New()

	err := r.w.queries().CreateOrReplaceSession(ctx, sqlc.CreateOrReplaceSessionParams{
		Uuid:     sessionUuid,
		QuizSha1: quizSha1,
		UserID:   userId,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return sessionUuid, nil
}

func (r *QuizDBRepository) AddSessionAnswer(ctx context.Context, sessionUuid uuid.UUID, questionSha1 string, answerSha1 string, checked bool) error {

	err := r.w.queries().CreateOrReplaceSessionAnswer(ctx, sqlc.CreateOrReplaceSessionAnswerParams{
		SessionUuid:  sessionUuid,
		QuestionSha1: questionSha1,
		AnswerSha1:   answerSha1,
		Checked:      checked,
	})
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" || err.Error() == "session is over" {
			return domain.Errorf(domain.InvalidArgument, err.Error())
		}
		return err
	}

	return nil
}

func (r *QuizDBRepository) FindAllQuizSessions(ctx context.Context, userId string, classId string, limit uint16, offset uint16) ([]*domain.QuizSession, error) {
	if isAdmin(userId) {
		quizSessions, err := r.w.queries().FindAllQuizSessions(ctx, sqlc.FindAllQuizSessionsParams{
			ClassId: classId,
			Limit:   int64(limit),
			Offset:  int64(offset),
		})
		if err != nil {
			return nil, err
		}

		return r.toQuizSessionArray(quizSessions, userId, true), nil
	}

	quizSessions, err := r.w.queries().FindAllQuizSessionsForUser(ctx, sqlc.FindAllQuizSessionsParams{
		UserId: userId,
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toQuizSessionArray(quizSessions, userId, false), nil
}

func (r *QuizDBRepository) FindQuizSessionByUuid(ctx context.Context, sessionUuid uuid.UUID) (*domain.QuizSessionDetail, error) {

	details, err := r.w.queries().FindQuizSessionByUuid(ctx, sessionUuid)
	if err != nil {
		return nil, err
	}

	if len(details) == 0 {
		return nil, pm.Errorf(http.StatusNotFound, "session with uuid: %s was not found.", sessionUuid)
	}

	sessionDetail := domain.QuizSessionDetail{}

	for _, entity := range details {
		if sessionDetail.QuizSha1 == "" {
			sessionDetail.SessionId = entity.SessionUuid
			sessionDetail.UserId = entity.UserID
			sessionDetail.RemainingSec = entity.RemainingSec
			sessionDetail.QuizSha1 = entity.QuizSha1
			sessionDetail.Name = entity.QuizName
			sessionDetail.QuizDuration = entity.QuizDuration
			sessionDetail.Questions = map[string]domain.QuizQuestion{}

			if sessionDetail.RemainingSec == 0 {
				sessionDetail.Result = &domain.SessionResult{
					GoodAnswer:  entity.Results,
					TotalAnswer: entity.CheckedAnswers,
				}
			}

		}

		if _, found := sessionDetail.Questions[entity.QuestionSha1]; !found {
			newQuestion := domain.QuizQuestion{
				Sha1:         entity.QuestionSha1,
				Position:     entity.QuestionPosition,
				Content:      entity.QuestionContent,
				Code:         entity.QuestionCode.String,
				CodeLanguage: entity.QuestionCodeLanguage.String,
				Answers:      map[string]domain.QuizQuestionAnswer{},
			}
			sessionDetail.Questions[entity.QuestionSha1] = newQuestion
		}

		answerValid := false
		if sessionDetail.RemainingSec == 0 {
			answerValid = entity.AnswerValid
		}

		sessionDetail.Questions[entity.QuestionSha1].Answers[entity.AnswerSha1] = domain.QuizQuestionAnswer{
			Sha1:    entity.AnswerSha1,
			Content: entity.AnswerContent,
			Checked: entity.AnswerChecked,
			Valid:   answerValid,
		}
	}

	return &sessionDetail, nil
}
