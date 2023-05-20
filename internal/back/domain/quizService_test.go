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
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const (
	Sha1Create = "fccc28a245ee3e92791ec9395d3a3791d17090da"
	Sha1Update = "fccc28a245ee3e92791ec9395d3a3791d17090db"
	Filename   = "marvel-universe.quiz.md"
	Name       = "Marvel Universe"
)

func TestQuizService_saveQuiz_create_quiz(t *testing.T) {

	mockQuizRepository := NewMockQuizRepository(t)

	s := NewQuizService(mockQuizRepository)

	q := &Quiz{
		Sha1:     Sha1Create,
		Filename: Filename,
		Name:     Name,
		Version:  1,
	}

	mockQuizRepository.On("FindLatestVersionByFilename", context.Background(), Filename).Return(nil, nil)
	mockQuizRepository.On("Create", context.Background(), q).Return(nil)

	// Test creation of quiz
	stats, err := s.SaveQuiz(context.Background(), q)
	if err != nil {
		assert.Failf(t, "Fail to save : %w", err.Error())
	}

	assert.Equal(t, 1, stats.Created)
	assert.Equal(t, 0, stats.Updated)

	mockQuizRepository.AssertExpectations(t)
}

func TestQuizService_saveQuiz_update_quiz(t *testing.T) {

	mockQuizRepository := NewMockQuizRepository(t)

	s := NewQuizService(mockQuizRepository)

	lastQuiz := &Quiz{
		Sha1:     Sha1Create,
		Filename: Filename,
		Name:     Name,
		Version:  1,
	}

	quizUpdate := &Quiz{
		Sha1:     Sha1Update,
		Filename: Filename,
		Name:     Name,
		Version:  2,
	}

	mockQuizRepository.On("FindLatestVersionByFilename", context.Background(), Filename).Return(lastQuiz, nil)
	mockQuizRepository.On("Create", context.Background(), quizUpdate).Return(nil)
	mockQuizRepository.On("ActivateOnlyVersion", context.Background(), Filename, 2).Return(nil)

	// Test update of quiz
	stats, err := s.SaveQuiz(context.Background(), &Quiz{
		Sha1:     Sha1Update,
		Filename: Filename,
		Name:     Name,
		Version:  1,
	})
	if err != nil {
		assert.Failf(t, "Fail to save : %w", err.Error())
	}

	assert.Equal(t, 0, stats.Created)
	assert.Equal(t, 1, stats.Updated)

	mockQuizRepository.AssertExpectations(t)
}
