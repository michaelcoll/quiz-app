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

	q *sqlc.Queries
}

func NewQuizRepository(c *sql.DB) *QuizDBRepository {
	return &QuizDBRepository{q: sqlc.New(c)}
}

func (r *QuizDBRepository) FindFullBySha1(ctx context.Context, sha1 string, userId string) (*domain.Quiz, error) {
	entities, err := r.q.FindQuizFullBySha1(ctx, sqlc.FindQuizFullBySha1Params{
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
			quiz.Version = int(entity.QuizVersion)
			quiz.Duration = int(entity.QuizDuration)
			quiz.CreatedAt = entity.QuizCreatedAt
			quiz.Questions = map[string]domain.QuizQuestion{}
		}

		if _, found := quiz.Questions[entity.QuestionSha1]; !found {
			newQuestion := domain.QuizQuestion{
				Sha1:    entity.QuestionSha1,
				Content: entity.QuestionContent,
				Answers: map[string]domain.QuizQuestionAnswer{},
			}
			quiz.Questions[entity.QuestionSha1] = newQuestion
		} else {
			quiz.Questions[entity.QuestionSha1].Answers[entity.AnswerSha1] = domain.QuizQuestionAnswer{
				Sha1:    entity.AnswerSha1,
				Content: entity.AnswerContent,
				Valid:   entity.AnswerValid,
			}
		}
	}

	return &quiz, nil
}

func (r *QuizDBRepository) FindAllActive(ctx context.Context, userId string, limit uint16, offset uint16) ([]*domain.Quiz, error) {
	quizzes, err := r.q.FindAllActiveQuiz(ctx, sqlc.FindAllActiveQuizParams{
		ID:     userId,
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toQuizArray(quizzes, userId == ""), nil
}

func (r *QuizDBRepository) CountAllActive(ctx context.Context, userId string) (uint32, error) {
	count, err := r.q.CountAllActiveQuiz(ctx, userId)
	if err != nil {
		return 0, err
	}

	return uint32(count), nil
}

func (r *QuizDBRepository) FindLatestVersionByFilename(ctx context.Context, filename string) (*domain.Quiz, error) {

	quiz, err := r.q.FindQuizByFilenameAndLatestVersion(ctx, filename)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return r.toQuiz(quiz), nil
}

func (r *QuizDBRepository) Create(ctx context.Context, quiz *domain.Quiz) error {

	err := r.q.CreateOrReplaceQuiz(ctx, sqlc.CreateOrReplaceQuizParams{
		Sha1:      quiz.Sha1,
		Name:      quiz.Name,
		Filename:  quiz.Filename,
		Version:   int64(quiz.Version),
		Duration:  int64(quiz.Duration),
		CreatedAt: quiz.CreatedAt,
	})
	if err != nil {
		return err
	}

	for _, question := range quiz.Questions {
		err := r.q.CreateOrReplaceQuestion(ctx, sqlc.CreateOrReplaceQuestionParams{
			Sha1:    question.Sha1,
			Content: question.Content,
		})
		if err != nil {
			return err
		}

		err = r.q.LinkQuestion(ctx, sqlc.LinkQuestionParams{
			QuizSha1:     quiz.Sha1,
			QuestionSha1: question.Sha1,
		})
		if err != nil {
			return err
		}

		for _, answer := range question.Answers {
			err := r.q.CreateOrReplaceAnswer(ctx, sqlc.CreateOrReplaceAnswerParams{
				Sha1:    answer.Sha1,
				Content: answer.Content,
				Valid:   answer.Valid,
			})
			if err != nil {
				return err
			}

			err = r.q.LinkAnswer(ctx, sqlc.LinkAnswerParams{
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
	err := r.q.ActivateOnlyVersion(ctx, sqlc.ActivateOnlyVersionParams{
		Filename: filename,
		Version:  int64(version),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *QuizDBRepository) toQuiz(entity sqlc.Quiz) *domain.Quiz {
	return &domain.Quiz{
		Sha1:      entity.Sha1,
		Filename:  entity.Filename,
		Name:      entity.Name,
		Version:   int(entity.Version),
		Duration:  int(entity.Duration),
		Active:    entity.Active,
		CreatedAt: entity.CreatedAt,
	}
}

func (r *QuizDBRepository) toQuizArray(entities []sqlc.FindAllActiveQuizRow, isAdmin bool) []*domain.Quiz {
	domains := make([]*domain.Quiz, len(entities))

	for i, entity := range entities {
		if isAdmin {
			domains[i] = &domain.Quiz{
				Sha1:      entity.Sha1,
				Filename:  entity.Filename,
				Name:      entity.Name,
				Version:   int(entity.Version),
				Duration:  int(entity.Duration),
				Active:    entity.Active,
				CreatedAt: entity.CreatedAt,
			}
		} else {
			domains[i] = &domain.Quiz{
				Sha1:     entity.Sha1,
				Name:     entity.Name,
				Duration: int(entity.Duration),
			}
		}
	}

	return domains
}

func (r *QuizDBRepository) FindAllSessions(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16) ([]*domain.Session, error) {
	if len(userId) > 0 {
		sessions, err := r.q.FindAllSessionsForUser(ctx, sqlc.FindAllSessionsForUserParams{
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

	sessions, err := r.q.FindAllSessions(ctx, sqlc.FindAllSessionsParams{
		QuizActive: quizActive,
		Limit:      int64(limit),
		Offset:     int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toSessionArray(sessions), nil
}

func (r *QuizDBRepository) FindAllQuizSessions(ctx context.Context, userId string, limit uint16, offset uint16) ([]*domain.QuizSession, error) {

	if userId == "" {
		quizSessions, err := r.q.FindAllQuizSessions(ctx, sqlc.FindAllQuizSessionsParams{
			Limit:  int64(limit),
			Offset: int64(offset),
		})
		if err != nil {
			return nil, err
		}

		return r.toQuizSessionArray(quizSessions, true), nil
	}

	quizSessions, err := r.q.FindAllQuizSessionsForUser(ctx, sqlc.FindAllQuizSessionsForUserParams{
		UserID: userId,
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toQuizSessionArray(quizSessions, false), nil
}

func (r *QuizDBRepository) CountAllSessions(ctx context.Context, quizActive bool, userId string) (uint32, error) {
	if len(userId) > 0 {
		count, err := r.q.CountAllSessionsForUser(ctx, sqlc.CountAllSessionsForUserParams{
			QuizActive: quizActive,
			UserID:     userId,
		})
		if err != nil {
			return 0, err
		}

		return uint32(count), nil
	}

	count, err := r.q.CountAllSessions(ctx, quizActive)
	if err != nil {
		return 0, err
	}

	return uint32(count), nil
}

func (r *QuizDBRepository) toSession(entity sqlc.SessionView) *domain.Session {

	d := domain.Session{
		Id:           entity.Uuid,
		QuizSha1:     entity.QuizSha1,
		QuizName:     entity.QuizName,
		QuizActive:   entity.QuizActive,
		UserId:       entity.UserID,
		UserName:     entity.UserName,
		RemainingSec: entity.RemainingSec,
	}

	if entity.RemainingSec == 0 {
		var goodAnswer int
		if entity.Results.Valid {
			goodAnswer = int(entity.Results.Float64)
		}
		d.Result = &domain.SessionResult{
			GoodAnswer:  goodAnswer,
			TotalAnswer: int(entity.CheckedAnswers),
		}
	}

	return &d
}

func (r *QuizDBRepository) toSessionArray(entities []sqlc.SessionView) []*domain.Session {
	domains := make([]*domain.Session, len(entities))

	for i, entity := range entities {
		domains[i] = r.toSession(entity)
	}

	return domains
}

func (r *QuizDBRepository) StartSession(ctx context.Context, userId string, quizSha1 string) (uuid.UUID, error) {
	sessionUuid := uuid.New()

	err := r.q.CreateOrReplaceSession(ctx, sqlc.CreateOrReplaceSessionParams{
		Uuid:     sessionUuid,
		QuizSha1: quizSha1,
		UserID:   userId,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return sessionUuid, nil
}

func (r *QuizDBRepository) AddSessionAnswer(ctx context.Context, sessionUuid uuid.UUID, userId string, questionSha1 string, answerSha1 string, checked bool) error {

	err := r.q.CreateOrReplaceSessionAnswer(ctx, sqlc.CreateOrReplaceSessionAnswerParams{
		SessionUuid:  sessionUuid,
		UserID:       userId,
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

func (r *QuizDBRepository) toQuizSession(entity sqlc.QuizSessionView, isAdmin bool) *domain.QuizSession {

	d := domain.QuizSession{
		QuizSha1: entity.QuizSha1,
		Name:     entity.QuizName,
		Duration: int(entity.QuizDuration),
	}

	if isAdmin {
		d.Filename = entity.QuizFilename
		d.Version = int(entity.QuizVersion)
		d.CreatedAt = entity.QuizCreatedAt
	}

	userSession := domain.UserSession{
		SessionId:    entity.SessionUuid,
		UserId:       entity.UserID,
		UserName:     entity.UserName,
		RemainingSec: entity.RemainingSec,
	}
	if entity.RemainingSec == 0 {
		userSession.Result = &domain.SessionResult{
			GoodAnswer:  entity.Results,
			TotalAnswer: entity.CheckedAnswers,
		}
	}
	if userSession.UserId != "" {
		userSessions := make([]*domain.UserSession, 1)
		d.UserSessions = userSessions
		d.UserSessions[0] = &userSession
	}

	return &d
}

func (r *QuizDBRepository) toQuizSessionArray(entities []sqlc.QuizSessionView, isAdmin bool) []*domain.QuizSession {

	m := make(map[string]*domain.QuizSession)
	for _, entity := range entities {
		quizSession := r.toQuizSession(entity, isAdmin)
		if existingSession, found := m[quizSession.QuizSha1]; found {
			existingSession.UserSessions = append(existingSession.UserSessions, quizSession.UserSessions...)
		} else {
			m[quizSession.QuizSha1] = quizSession
		}
	}

	domains := make([]*domain.QuizSession, len(m))
	i := 0
	for _, session := range m {
		domains[i] = session
		i++
	}

	return domains
}
