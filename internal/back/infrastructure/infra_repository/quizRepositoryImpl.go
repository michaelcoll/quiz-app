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

package infra_repository

import (
	"context"
	"database/sql"

	"github.com/school-by-hiit/quiz-app/internal/back/domain/model"
	"github.com/school-by-hiit/quiz-app/internal/back/domain/repository"
	"github.com/school-by-hiit/quiz-app/internal/back/infrastructure/db"
	"github.com/school-by-hiit/quiz-app/internal/back/infrastructure/sqlc"
)

type QuizDBRepository struct {
	repository.QuizRepository

	c *sql.DB
	q *sqlc.Queries
}

func New() *QuizDBRepository {
	connection := db.Connect(false, "data")
	db.New(connection).Migrate()

	return &QuizDBRepository{q: sqlc.New(connection), c: connection}
}

func (r *QuizDBRepository) Close() {
	r.c.Close()
}

func (r *QuizDBRepository) Create(ctx context.Context, quiz model.Quiz) error {

	err := r.q.CreateOrReplaceQuiz(ctx, sqlc.CreateOrReplaceQuizParams{
		Sha1:     quiz.Sha1,
		Name:     quiz.Name,
		Filename: quiz.Filename,
		Version:  1,
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

			var valid int64
			if answer.Valid {
				valid = 1
			} else {
				valid = 0
			}

			err := r.q.CreateOrReplaceAnswer(ctx, sqlc.CreateOrReplaceAnswerParams{
				Sha1:    answer.Sha1,
				Content: answer.Content,
				Valid:   valid,
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
