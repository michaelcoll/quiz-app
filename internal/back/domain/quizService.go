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

package domain

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type QuizService struct {
	r QuizRepository
}

func NewQuizService(r QuizRepository) QuizService {
	return QuizService{r: r}
}

func (s *QuizService) FindFullBySha1(ctx context.Context, sha1 string, userId string) (*Quiz, error) {
	quiz, err := s.r.FindFullBySha1(ctx, sha1, userId)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (s *QuizService) FindAllActive(ctx context.Context, userId string, limit uint16, offset uint16) ([]*Quiz, uint32, error) {
	quizzes, err := s.r.FindAllActive(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.r.CountAllActive(ctx, userId)
	if err != nil {
		return nil, 0, err
	}

	return quizzes, count, nil
}

func (s *QuizService) Sync(ctx context.Context) error {

	quizzes, err := s.ScanGitRepo()
	if err != nil {
		return err
	}

	var syncStats SyncStats
	for _, quiz := range quizzes {
		stats, err := s.SaveQuiz(ctx, quiz)
		if err != nil {
			return err
		}

		syncStats = addStats(syncStats, stats)
	}

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

	return nil
}

func (s *QuizService) SaveQuiz(ctx context.Context, quiz *Quiz) (*SyncStats, error) {

	verbose := viper.GetBool("verbose")

	latestQuiz, err := s.r.FindLatestVersionByFilename(ctx, quiz.Filename)
	if err != nil {
		return nil, err
	}
	if latestQuiz == nil {
		if verbose {
			fmt.Printf("%s Creating quiz %s\n",
				color.GreenString("✓"),
				quiz.Filename)
		}

		err = s.r.Create(ctx, quiz)
		if err != nil {
			return nil, err
		}

		return &SyncStats{
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
			return nil, err
		}

		err := s.r.ActivateOnlyVersion(ctx, quiz.Filename, quiz.Version)
		if err != nil {
			return nil, err
		}

		return &SyncStats{
			Updated: 1,
			Created: 0,
		}, nil
	} else {
		return &SyncStats{
			Updated: 0,
			Created: 0,
		}, nil
	}
}

func addStats(stat1 SyncStats, stat2 *SyncStats) SyncStats {
	return SyncStats{
		Created: stat1.Created + stat2.Created,
		Updated: stat1.Updated + stat2.Updated,
	}
}

func (s *QuizService) FindAllSessions(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16) ([]*Session, uint32, error) {
	sessions, err := s.r.FindAllSessions(ctx, quizActive, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.r.CountAllSessions(ctx, quizActive, userId)
	if err != nil {
		return nil, 0, err
	}

	return sessions, count, nil
}

func (s *QuizService) StartSession(ctx context.Context, userId string, quizSha1 string) (uuid.UUID, error) {
	return s.r.StartSession(ctx, userId, quizSha1)
}

func (s *QuizService) AddSessionAnswer(ctx context.Context, sessionUuid uuid.UUID, userId string, questionSha1 string, answerSha1 string, checked bool) error {
	return s.r.AddSessionAnswer(ctx, sessionUuid, userId, questionSha1, answerSha1, checked)
}

func (s *QuizService) FindAllQuizSessions(ctx context.Context, userId string, limit uint16, offset uint16) ([]*QuizSession, uint32, error) {
	quizzes, err := s.r.FindAllQuizSessions(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.r.CountAllActive(ctx, userId)
	if err != nil {
		return nil, 0, err
	}

	return quizzes, count, nil
}
