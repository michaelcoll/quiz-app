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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/school-by-hiit/quiz-app/internal/back/domain/model"
	"github.com/school-by-hiit/quiz-app/internal/back/infrastructure/infra_repository"
)

const (
	Sha1Create = "fccc28a245ee3e92791ec9395d3a3791d17090da"
	Sha1Update = "fccc28a245ee3e92791ec9395d3a3791d17090db"
	Filename   = "marvel-universe.quiz.md"
	Name       = "Marvel Universe"
)

func TestQuizService_saveQuiz(t *testing.T) {

	// Deleting previous database
	if _, err := os.Stat("data"); err == nil {
		err := os.RemoveAll("data")
		if err != nil {
			assert.Failf(t, "Fail to delete data folder", "%v", err)
		}
	}

	r := infra_repository.New()
	defer r.Close()
	s := New(r)

	// Test creation of quiz
	stats, err := s.saveQuiz(context.Background(), model.Quiz{
		Sha1:     Sha1Create,
		Filename: Filename,
		Name:     Name,
		Version:  1,
	}, false)
	if err != nil {
		assert.Failf(t, "Fail to save : %w", err.Error())
	}

	assert.Equal(t, 1, stats.Created)
	assert.Equal(t, 0, stats.Updated)

	quiz, err := r.FindBySha1(context.Background(), Sha1Create)
	if err != nil {
		assert.Failf(t, "Fail to find quiz", "sha1=%s (%v)", Sha1Create, err)
	}

	assert.Equal(t, Sha1Create, quiz.Sha1)
	assert.Equal(t, Filename, quiz.Filename)
	assert.Equal(t, Name, quiz.Name)
	assert.Equal(t, 1, quiz.Version)

	// Test update of quiz
	stats, err = s.saveQuiz(context.Background(), model.Quiz{
		Sha1:     Sha1Update,
		Filename: Filename,
		Name:     Name,
		Version:  1,
	}, false)
	if err != nil {
		assert.Failf(t, "Fail to save : %w", err.Error())
	}

	assert.Equal(t, 0, stats.Created)
	assert.Equal(t, 1, stats.Updated)

	// Testing if quiz is properly updated
	quiz, err = r.FindBySha1(context.Background(), Sha1Update)
	if err != nil {
		assert.Failf(t, "Fail to find quiz", "sha1=%s (%v)", Sha1Update, err)
	}

	assert.Equal(t, Sha1Update, quiz.Sha1)
	assert.Equal(t, Filename, quiz.Filename)
	assert.Equal(t, Name, quiz.Name)
	assert.Equal(t, 2, quiz.Version)

	// Testing if previous version of the quiz is still here
	quiz, err = r.FindBySha1(context.Background(), Sha1Create)
	if err != nil {
		assert.Failf(t, "Fail to find quiz", "sha1=%s (%v)", Sha1Create, err)
	}

	assert.Equal(t, Sha1Create, quiz.Sha1)
	assert.Equal(t, Filename, quiz.Filename)
	assert.Equal(t, Name, quiz.Name)
	assert.Equal(t, 1, quiz.Version)
}
