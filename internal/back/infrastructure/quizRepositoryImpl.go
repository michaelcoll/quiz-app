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

func (r *QuizDBRepository) FindBySha1(ctx context.Context, sha1 string) (*domain.Quiz, error) {
	entity, err := r.q.FindQuizBySha1(ctx, sha1)
	if err != nil {
		return nil, err
	}

	return r.toQuiz(entity), nil
}

func (r *QuizDBRepository) FindFullBySha1(ctx context.Context, sha1 string) (*domain.Quiz, error) {
	entities, err := r.q.FindQuizFullBySha1(ctx, sha1)
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
			quiz.Active = intToBool(entity.QuizActive)
			quiz.Version = int(entity.QuizVersion)
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
				Valid:   intToBool(entity.AnswerValid),
			}
		}
	}

	return &quiz, nil
}

func (r *QuizDBRepository) FindAllActive(ctx context.Context, limit uint16, offset uint16) ([]*domain.Quiz, error) {
	quizzes, err := r.q.FindAllActiveQuiz(ctx, sqlc.FindAllActiveQuizParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toQuizArray(quizzes), nil
}

func (r *QuizDBRepository) CountAllActive(ctx context.Context) (uint32, error) {
	count, err := r.q.CountAllActiveQuiz(ctx)
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
		Sha1:     quiz.Sha1,
		Name:     quiz.Name,
		Filename: quiz.Filename,
		Version:  int64(quiz.Version),
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
				Valid:   boolToInt(answer.Valid),
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
		Active:    intToBool(entity.Active),
		CreatedAt: entity.CreatedAt,
	}
}

func (r *QuizDBRepository) toQuizArray(entities []sqlc.Quiz) []*domain.Quiz {
	domains := make([]*domain.Quiz, len(entities))

	for i, entity := range entities {
		domains[i] = r.toQuiz(entity)
	}

	return domains
}
