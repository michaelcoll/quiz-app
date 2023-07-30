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
	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

func (r *QuizDBRepository) toQuiz(entity sqlc.Quiz) *domain.Quiz {
	return &domain.Quiz{
		Sha1:      entity.Sha1,
		Filename:  entity.Filename,
		Name:      entity.Name,
		Version:   entity.Version,
		Duration:  entity.Duration,
		Active:    entity.Active,
		CreatedAt: entity.CreatedAt,
	}
}

func (r *QuizDBRepository) toQuizArray(entities []sqlc.FindAllActiveQuizRow, isAdmin bool) []*domain.Quiz {
	domains := make([]*domain.Quiz, len(entities))

	for i, entity := range entities {
		if isAdmin {
			domains[i] = &domain.Quiz{
				Sha1:      entity.Sha1,
				Filename:  entity.Filename,
				Name:      entity.Name,
				Version:   entity.Version,
				Duration:  entity.Duration,
				Active:    entity.Active,
				CreatedAt: entity.CreatedAt,
			}
		} else {
			domains[i] = &domain.Quiz{
				Sha1:     entity.Sha1,
				Name:     entity.Name,
				Duration: entity.Duration,
			}
		}
	}

	return domains
}

func (r *QuizDBRepository) toSession(entity sqlc.SessionView) *domain.Session {

	d := domain.Session{
		Id:           entity.Uuid,
		QuizSha1:     entity.QuizSha1,
		QuizName:     entity.QuizName,
		QuizActive:   entity.QuizActive,
		UserId:       entity.UserID,
		UserName:     entity.UserName,
		RemainingSec: entity.RemainingSec,
	}

	if entity.RemainingSec == 0 {
		d.Result = &domain.SessionResult{
			GoodAnswer:  entity.Results,
			TotalAnswer: entity.CheckedAnswers,
		}
	}

	return &d
}

func (r *QuizDBRepository) toSessionArray(entities []sqlc.SessionView) []*domain.Session {
	domains := make([]*domain.Session, len(entities))

	for i, entity := range entities {
		domains[i] = r.toSession(entity)
	}

	return domains
}

func (r *QuizDBRepository) toQuizSession(entity sqlc.QuizSessionView, isAdmin bool) *domain.QuizSession {

	d := domain.QuizSession{
		QuizSha1: entity.QuizSha1,
		Name:     entity.QuizName,
		Duration: entity.QuizDuration,
	}

	if isAdmin {
		d.Filename = entity.QuizFilename
		d.Version = entity.QuizVersion
		d.CreatedAt = entity.QuizCreatedAt
	}

	userSession := domain.UserSession{
		SessionId:    entity.SessionUuid,
		UserId:       entity.UserID,
		UserName:     entity.UserName,
		RemainingSec: entity.RemainingSec,
	}
	if entity.RemainingSec == 0 {
		userSession.Result = &domain.SessionResult{
			GoodAnswer:  entity.Results,
			TotalAnswer: entity.CheckedAnswers,
		}
	}
	if userSession.UserId != "" {
		userSessions := make([]*domain.UserSession, 1)
		d.UserSessions = userSessions
		d.UserSessions[0] = &userSession
	}

	return &d
}

func (r *QuizDBRepository) toQuizSessionArray(entities []sqlc.QuizSessionView, isAdmin bool) []*domain.QuizSession {

	m := make(map[string]*domain.QuizSession)
	for _, entity := range entities {
		quizSession := r.toQuizSession(entity, isAdmin)
		if existingSession, found := m[quizSession.QuizSha1]; found {
			existingSession.UserSessions = append(existingSession.UserSessions, quizSession.UserSessions...)
		} else {
			m[quizSession.QuizSha1] = quizSession
		}
	}

	domains := make([]*domain.QuizSession, len(m))
	i := 0
	for _, session := range m {
		domains[i] = session
		i++
	}

	return domains
}
