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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

func TestQuizDBRepository_toQuizSessionArray_2quiz_1session(t *testing.T) {
	// Given
	r := NewQuizRepository(nil)
	sessions := make([]sqlc.QuizSessionView, 2)
	sessions[0] = sqlc.QuizSessionView{
		QuizSha1:       sha1Quiz1,
		QuizName:       quizName1,
		QuizFilename:   quizFilename1,
		QuizVersion:    quizVersion1,
		QuizDuration:   quizDuration1,
		QuizCreatedAt:  quizCreatedAt1,
		SessionUuid:    uuid.New(),
		UserID:         userId1,
		UserName:       userName,
		RemainingSec:   remainingSec1,
		CheckedAnswers: checkedAnswers1,
		Results:        results1,
	}
	sessions[1] = sqlc.QuizSessionView{
		QuizSha1:      sha1Quiz2,
		QuizName:      quizName2,
		QuizFilename:  quizFilename2,
		QuizVersion:   quizVersion2,
		QuizDuration:  quizDuration2,
		QuizCreatedAt: quizCreatedAt2,
	}

	// When
	actual := r.toQuizSessionArray(sessions, true)

	// Then
	assert.Len(t, actual, 2)

	actualMap := make(map[string]*domain.QuizSession)
	for _, session := range actual {
		assert.NotEmpty(t, session)
		actualMap[session.QuizSha1] = session
	}

	assert.Contains(t, actualMap, sha1Quiz1)
	assert.Contains(t, actualMap, sha1Quiz2)

	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions)
	assert.Len(t, actualMap[sha1Quiz1].UserSessions, 1)
	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions[0])
	assert.Equal(t, userId1, actualMap[sha1Quiz1].UserSessions[0].UserId)
}

func TestQuizDBRepository_toQuizSessionArray_3quiz_2session(t *testing.T) {
	// Given
	r := NewQuizRepository(nil)
	sessions := make([]sqlc.QuizSessionView, 3)
	sessions[0] = sqlc.QuizSessionView{
		QuizSha1:       sha1Quiz1,
		QuizName:       quizName1,
		QuizFilename:   quizFilename1,
		QuizVersion:    quizVersion1,
		QuizDuration:   quizDuration1,
		QuizCreatedAt:  quizCreatedAt1,
		SessionUuid:    uuid.New(),
		UserID:         userId1,
		UserName:       userName,
		RemainingSec:   remainingSec1,
		CheckedAnswers: checkedAnswers1,
		Results:        results1,
	}
	sessions[1] = sqlc.QuizSessionView{
		QuizSha1:       sha1Quiz1,
		QuizName:       quizName1,
		QuizFilename:   quizFilename1,
		QuizVersion:    quizVersion1,
		QuizDuration:   quizDuration1,
		QuizCreatedAt:  quizCreatedAt1,
		SessionUuid:    uuid.New(),
		UserID:         userId2,
		UserName:       userName,
		RemainingSec:   remainingSec2,
		CheckedAnswers: checkedAnswers2,
		Results:        results2,
	}
	sessions[2] = sqlc.QuizSessionView{
		QuizSha1:      sha1Quiz2,
		QuizName:      quizName2,
		QuizFilename:  quizFilename2,
		QuizVersion:   quizVersion2,
		QuizDuration:  quizDuration2,
		QuizCreatedAt: quizCreatedAt2,
	}

	// When
	actual := r.toQuizSessionArray(sessions, true)

	// Then
	assert.Len(t, actual, 2)

	actualMap := make(map[string]*domain.QuizSession)
	for _, session := range actual {
		assert.NotEmpty(t, session)
		actualMap[session.QuizSha1] = session
	}

	assert.Contains(t, actualMap, sha1Quiz1)
	assert.Contains(t, actualMap, sha1Quiz2)

	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions)
	assert.Len(t, actualMap[sha1Quiz1].UserSessions, 2)
	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions[0])
	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions[0].UserId)
	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions[1])
	assert.NotEmpty(t, actualMap[sha1Quiz1].UserSessions[1].UserId)
}
