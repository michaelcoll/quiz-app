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

package presentation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

const (
	quizSha1         = "42"
	quizFilename     = "42"
	quizName         = "42"
	quizVersion      = 1
	quizCreatedAt    = "sometime"
	quizActive       = true
	quizDuration     = 20
	question1Sha1    = "43"
	question1Content = "Question 1 ?"
	question2Sha1    = "44"
	question2Content = "Question 2 ?"
	answer1Sha1      = "45"
	answer1Content   = "Yes"
	answer2Sha1      = "46"
	answer2Content   = "No"
	answer3Sha1      = "47"
	answer3Content   = "Yes1"
	answer4Sha1      = "48"
	answer4Content   = "No1"
)

func TestQuiz_fromDomain(t *testing.T) {
	questions1 := make(map[string]domain.QuizQuestion)
	answers1 := make(map[string]domain.QuizQuestionAnswer)
	answers1[answer1Sha1] = domain.QuizQuestionAnswer{
		Sha1:    answer1Sha1,
		Content: answer1Content,
		Valid:   false,
	}
	answers1[answer2Sha1] = domain.QuizQuestionAnswer{
		Sha1:    answer2Sha1,
		Content: answer2Content,
		Valid:   false,
	}
	answers2 := make(map[string]domain.QuizQuestionAnswer)
	answers2[answer3Sha1] = domain.QuizQuestionAnswer{
		Sha1:    answer3Sha1,
		Content: answer3Content,
		Valid:   false,
	}
	answers2[answer4Sha1] = domain.QuizQuestionAnswer{
		Sha1:    answer4Sha1,
		Content: answer4Content,
		Valid:   false,
	}

	questions1[question1Sha1] = domain.QuizQuestion{
		Sha1:    question1Sha1,
		Content: question1Content,
		Answers: answers1,
	}
	questions1[question2Sha1] = domain.QuizQuestion{
		Sha1:    question2Sha1,
		Content: question2Content,
		Answers: answers2,
	}

	q := &domain.Quiz{
		Sha1:      quizSha1,
		Filename:  quizFilename,
		Name:      quizName,
		Version:   quizVersion,
		CreatedAt: quizCreatedAt,
		Active:    quizActive,
		Duration:  quizDuration,
		Questions: questions1,
	}

	dto := Quiz{}
	dto.fromDomain(q)

	assert.Equal(t, quizSha1, dto.Sha1)
	assert.Equal(t, quizFilename, dto.Filename)
	assert.Equal(t, quizName, dto.Name)
	assert.Equal(t, quizVersion, dto.Version)
	assert.Equal(t, quizCreatedAt, dto.CreatedAt)
	assert.Equal(t, quizDuration, dto.Duration)

	assert.Len(t, dto.Questions, 2)

	assert.Equal(t, question1Sha1, dto.Questions[0].Sha1)
	assert.Equal(t, question1Content, dto.Questions[0].Content)
	assert.Len(t, dto.Questions[0].Answers, 2)
	assert.Equal(t, answer1Sha1, dto.Questions[0].Answers[0].Sha1)
	assert.Equal(t, answer1Content, dto.Questions[0].Answers[0].Content)
	assert.Equal(t, answer2Sha1, dto.Questions[0].Answers[1].Sha1)
	assert.Equal(t, answer2Content, dto.Questions[0].Answers[1].Content)

	assert.Equal(t, question2Sha1, dto.Questions[1].Sha1)
	assert.Equal(t, question2Content, dto.Questions[1].Content)
	assert.Len(t, dto.Questions[1].Answers, 2)
	assert.Equal(t, answer3Sha1, dto.Questions[1].Answers[0].Sha1)
	assert.Equal(t, answer3Content, dto.Questions[1].Answers[0].Content)
	assert.Equal(t, answer4Sha1, dto.Questions[1].Answers[1].Sha1)
	assert.Equal(t, answer4Content, dto.Questions[1].Answers[1].Content)
}
