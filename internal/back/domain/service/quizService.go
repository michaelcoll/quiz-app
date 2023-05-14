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
	"strconv"

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

func (s *QuizService) Close() {
	s.r.Close()
}

func (s *QuizService) Sync(ctx context.Context, repoUrl string, token string, verbose bool) error {

	quizzes, err := s.ScanGitRepo(repoUrl, token)
	if err != nil {
		return err
	}

	var syncStats model.SyncStats
	for _, quiz := range quizzes {
		stats, err := s.saveQuiz(ctx, quiz, verbose)
		if err != nil {
			return err
		}

		syncStats = addStats(syncStats, stats)
	}

	if verbose {
		if syncStats.Created > 0 || syncStats.Updated > 0 {
			fmt.Printf("%s Repo synced (%s quiz(zes) created, %s quiz(zes) updated)\n",
				color.GreenString("✓"),
				color.BlueString(strconv.Itoa(syncStats.Created)),
				color.BlueString(strconv.Itoa(syncStats.Updated)))
		} else {
			fmt.Printf("%s Repo synced %s\n",
				color.GreenString("✓"),
				color.BlueString(color.New(color.FgHiBlack).Sprintf(" — no changes")))
		}
	}

	return nil
}

func (s *QuizService) saveQuiz(ctx context.Context, quiz model.Quiz, verbose bool) (model.SyncStats, error) {

	latestQuiz, err := s.r.FindLatestVersionByFilename(ctx, quiz.Filename)
	if err != nil {
		if verbose {
			fmt.Printf("%s Creating quiz %s\n",
				color.GreenString("✓"),
				quiz.Filename)
		}

		err = s.r.Create(ctx, quiz)
		if err != nil {
			return model.SyncStats{}, err
		}

		return model.SyncStats{
			Updated: 0,
			Created: 1,
		}, nil
	} else if latestQuiz.Sha1 != quiz.Sha1 {
		quiz.Version = latestQuiz.Version + 1

		if verbose {
			fmt.Printf("%s Updating quiz %s to version %d\n",
				color.GreenString("✓"),
				quiz.Filename, quiz.Version)
		}

		err = s.r.Create(ctx, quiz)
		if err != nil {
			return model.SyncStats{}, err
		}

		err := s.r.ActivateOnlyVersion(ctx, quiz.Filename, quiz.Version)
		if err != nil {
			return model.SyncStats{}, err
		}

		return model.SyncStats{
			Updated: 1,
			Created: 0,
		}, nil
	} else {
		return model.SyncStats{
			Updated: 0,
			Created: 0,
		}, nil
	}
}

func addStats(stat1 model.SyncStats, stat2 model.SyncStats) model.SyncStats {
	return model.SyncStats{
		Created: stat1.Created + stat2.Created,
		Updated: stat1.Updated + stat2.Updated,
	}
}
