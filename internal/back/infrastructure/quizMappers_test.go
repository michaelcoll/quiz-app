/*
 * Copyright (c) 2024-2025 Michaël COLL.
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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

func TestQuizDBRepository_toQuizSessionArray_emptySessions(t *testing.T) {
	r := NewQuizRepository(nil)
	var sessions []sqlc.QuizSessionView

	actual := r.toQuizSessionArray(sessions, userId1, true)

	assert.Len(t, actual, 0)
}

func TestQuizDBRepository_toQuizSessionArray_singleSession(t *testing.T) {
	r := NewQuizRepository(nil)
	sessions := []sqlc.QuizSessionView{
		{
			QuizSha1:       sha1Quiz1,
			QuizName:       quizName1,
			QuizFilename:   quizFilename1,
			QuizVersion:    quizVersion1,
			QuizDuration:   quizDuration1,
			QuizCreatedAt:  quizCreatedAt1,
			SessionUuid:    uuid.New(),
			UserID:         userId1,
			UserName:       name,
			RemainingSec:   remainingSec1,
			CheckedAnswers: checkedAnswers1,
			Results:        results1,
		},
	}

	actual := r.toQuizSessionArray(sessions, userId1, true)

	assert.Len(t, actual, 1)
	assert.Equal(t, sha1Quiz1, actual[0].QuizSha1)
	assert.Len(t, actual[0].UserSessions, 1)
	assert.Equal(t, userId1, actual[0].UserSessions[0].UserId)
}

func TestQuizDBRepository_toQuizSessionArray_multipleSessionsSameQuiz(t *testing.T) {
	r := NewQuizRepository(nil)
	sessions := []sqlc.QuizSessionView{
		{
			QuizSha1:       sha1Quiz1,
			QuizName:       quizName1,
			QuizFilename:   quizFilename1,
			QuizVersion:    quizVersion1,
			QuizDuration:   quizDuration1,
			QuizCreatedAt:  quizCreatedAt1,
			SessionUuid:    uuid.New(),
			UserID:         userId1,
			UserName:       name,
			RemainingSec:   remainingSec1,
			CheckedAnswers: checkedAnswers1,
			Results:        results1,
		},
		{
			QuizSha1:       sha1Quiz1,
			QuizName:       quizName1,
			QuizFilename:   quizFilename1,
			QuizVersion:    quizVersion1,
			QuizDuration:   quizDuration1,
			QuizCreatedAt:  quizCreatedAt1,
			SessionUuid:    uuid.New(),
			UserID:         userId2,
			UserName:       name,
			RemainingSec:   remainingSec2,
			CheckedAnswers: checkedAnswers2,
			Results:        results2,
		},
	}

	actual := r.toQuizSessionArray(sessions, userId1, true)

	assert.Len(t, actual, 1)
	assert.Equal(t, sha1Quiz1, actual[0].QuizSha1)
	assert.Len(t, actual[0].UserSessions, 2)
}

func TestQuizDBRepository_toQuizSessionArray_multipleQuizzes(t *testing.T) {
	r := NewQuizRepository(nil)
	sessions := []sqlc.QuizSessionView{
		{
			QuizSha1:       sha1Quiz1,
			QuizName:       quizName1,
			QuizFilename:   quizFilename1,
			QuizVersion:    quizVersion1,
			QuizDuration:   quizDuration1,
			QuizCreatedAt:  quizCreatedAt1,
			SessionUuid:    uuid.New(),
			UserID:         userId1,
			UserName:       name,
			RemainingSec:   remainingSec1,
			CheckedAnswers: checkedAnswers1,
			Results:        results1,
		},
		{
			QuizSha1:      sha1Quiz2,
			QuizName:      quizName2,
			QuizFilename:  quizFilename2,
			QuizVersion:   quizVersion2,
			QuizDuration:  quizDuration2,
			QuizCreatedAt: quizCreatedAt2,
		},
	}

	actual := r.toQuizSessionArray(sessions, userId1, true)

	assert.Len(t, actual, 2)
}
