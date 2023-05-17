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
	"time"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

type Quiz struct {
	Sha1      string         `json:"sha1"`
	Filename  string         `json:"filename"`
	Name      string         `json:"name"`
	Version   int            `json:"version"`
	CreatedAt time.Time      `json:"createdAt"`
	Active    bool           `json:"active"`
	Questions []QuizQuestion `json:"questions,omitempty"`
}

type QuizQuestion struct {
	Sha1    string               `json:"sha1"`
	Content string               `json:"content"`
	Answers []QuizQuestionAnswer `json:"answers,omitempty"`
}

type QuizQuestionAnswer struct {
	Sha1    string `json:"sha1"`
	Content string `json:"content"`
}

func fromDomain(domain domain.Quiz) Quiz {
	quiz := Quiz{
		Sha1:      domain.Sha1,
		Filename:  domain.Filename,
		Name:      domain.Name,
		Version:   domain.Version,
		CreatedAt: domain.CreatedAt,
		Active:    domain.Active,
		Questions: make([]QuizQuestion, len(domain.Questions)),
	}

	i := 0
	for _, q := range domain.Questions {

		j := 0
		answers := make([]QuizQuestionAnswer, len(q.Answers))
		for _, a := range q.Answers {
			answers[j] = QuizQuestionAnswer{
				Sha1:    a.Sha1,
				Content: a.Content,
			}
			j++
		}

		quiz.Questions[i] = QuizQuestion{
			Sha1:    q.Sha1,
			Content: q.Content,
			Answers: answers,
		}
		i++
	}

	return quiz
}

func fromDomains(domains []domain.Quiz) []Quiz {
	dtos := make([]Quiz, len(domains))

	for i, d := range domains {
		dtos[i] = fromDomain(d)
	}

	return dtos
}
