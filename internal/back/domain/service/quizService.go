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

package service

import (
	"context"
	"fmt"

	"github.com/fatih/color"

	"github.com/school-by-hiit/quiz-app/internal/back/domain/model"
	"github.com/school-by-hiit/quiz-app/internal/back/domain/repository"
)

type QuizService struct {
	r repository.QuizRepository
}

func New(r repository.QuizRepository) QuizService {
	return QuizService{r: r}
}

func (s *QuizService) Sync(ctx context.Context, repoUrl string, token string) error {

	quizzes, err := s.ScanGitRepo(repoUrl, token)
	if err != nil {
		return err
	}

	for _, quiz := range quizzes {
		err := s.saveQuiz(ctx, quiz)
		if err != nil {
			return err
		}
	}

	fmt.Printf("%s Repo %s synced.\n", color.GreenString("✓"), color.BlueString(repoUrl))

	return nil
}

func (s *QuizService) saveQuiz(ctx context.Context, quiz model.Quiz) error {

	err := s.r.Create(ctx, quiz)
	if err != nil {
		return err
	}

	return nil
}
