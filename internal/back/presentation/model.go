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
	"net/http"
	"regexp"
	"sort"

	"github.com/google/uuid"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

type QuizInfos interface {
	setSha1AndName(sha1 string, name string)
	setQuestions(questions []QuizQuestion)
}

type Quiz struct {
	Sha1      string         `json:"sha1"`
	Filename  string         `json:"filename"`
	Name      string         `json:"name"`
	Version   int            `json:"version"`
	CreatedAt string         `json:"createdAt"`
	Duration  int            `json:"duration"`
	Active    bool           `json:"active"`
	Questions []QuizQuestion `json:"questions,omitempty"`
}

func (dto *Quiz) setSha1AndName(sha1 string, name string) {
	dto.Sha1 = sha1
	dto.Name = name
}

func (dto *Quiz) setQuestions(questions []QuizQuestion) {
	dto.Questions = questions
}

type QuizQuestion struct {
	Sha1     string               `json:"sha1"`
	Content  string               `json:"content"`
	Position int                  `json:"position"`
	Answers  []QuizQuestionAnswer `json:"answers,omitempty"`
}

type QuizQuestionAnswer struct {
	Sha1    string `json:"sha1"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
	Valid   bool   `json:"valid,omitempty"`
}

func mapQuizInfos(d domain.QuizInfos, dto QuizInfos) {
	dto.setSha1AndName(d.GetSha1AndName())

	questions := make([]QuizQuestion, len(d.GetQuestions()))

	i := 0
	for _, question := range d.GetQuestions() {

		j := 0
		answers := make([]QuizQuestionAnswer, len(question.Answers))
		for _, a := range question.Answers {
			answers[j] = QuizQuestionAnswer{
				Sha1:    a.Sha1,
				Content: a.Content,
				Checked: a.Checked,
				Valid:   a.Valid,
			}
			j++
		}

		questions[i] = QuizQuestion{
			Sha1:     question.Sha1,
			Content:  question.Content,
			Position: question.Position,
			Answers:  answers,
		}
		i++
	}

	sort.SliceStable(questions, func(i, j int) bool {
		return questions[i].Position < questions[j].Position
	})

	dto.setQuestions(questions)
}

func (dto *Quiz) fromDomain(d *domain.Quiz) *Quiz {
	dto.Filename = d.Filename
	dto.Version = d.Version
	dto.Duration = d.Duration
	dto.CreatedAt = d.CreatedAt
	dto.Active = d.Active

	mapQuizInfos(d, dto)

	return dto
}

func toQuizDtos(domains []*domain.Quiz) []*Quiz {
	dtos := make([]*Quiz, len(domains))

	for i, d := range domains {
		dto := &Quiz{}
		dtos[i] = dto.fromDomain(d)
	}

	return dtos
}

type endPointDef struct {
	regex  *regexp.Regexp
	method string
}

func (e *endPointDef) match(request *http.Request) bool {
	path := request.URL.Path
	method := request.Method

	return e.regex.MatchString(path) && e.method == method
}

type Role string

const (
	NoRole  Role = "NO_ROLE"
	Admin   Role = "ADMIN"
	Teacher Role = "TEACHER"
	Student Role = "STUDENT"
)

func toRoleDto(d domain.Role) Role {
	dto := NoRole
	switch d {
	case domain.Admin:
		dto = Admin
	case domain.Teacher:
		dto = Teacher
	case domain.Student:
		dto = Student
	}
	return dto
}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Active    bool   `json:"active"`
	Role      Role   `json:"role"`
}

func (dto *User) fromDomain(d *domain.User) *User {
	dto.Id = d.Id
	dto.Email = d.Email
	dto.Firstname = d.Firstname
	dto.Lastname = d.Lastname
	dto.Active = d.Active
	dto.Role = toRoleDto(d.Role)

	return dto
}

type SessionResult struct {
	GoodAnswer  int `json:"goodAnswer,omitempty"`
	TotalAnswer int `json:"totalAnswer,omitempty"`
}

type Session struct {
	Id           uuid.UUID      `json:"id"`
	QuizSha1     string         `json:"quizSha1,omitempty"`
	QuizName     string         `json:"quizName,omitempty"`
	QuizActive   bool           `json:"quizActive,omitempty"`
	UserId       string         `json:"userId,omitempty"`
	UserName     string         `json:"userName,omitempty"`
	RemainingSec int            `json:"remainingSec,omitempty"`
	Result       *SessionResult `json:"result,omitempty"`
}

func (dto *Session) fromDomain(d *domain.Session) *Session {
	dto.Id = d.Id
	dto.QuizSha1 = d.QuizSha1
	dto.UserName = d.UserName
	dto.QuizActive = d.QuizActive
	dto.UserId = d.UserId
	dto.UserName = d.UserName
	dto.RemainingSec = d.RemainingSec
	if d.Result != nil {
		dto.Result = &SessionResult{
			GoodAnswer:  d.Result.GoodAnswer,
			TotalAnswer: d.Result.TotalAnswer,
		}
	}

	return dto
}

func toSessionDtos(domains []*domain.Session) []*Session {
	dtos := make([]*Session, len(domains))

	for i, d := range domains {
		dto := &Session{}
		dtos[i] = dto.fromDomain(d)
	}

	return dtos
}

type SessionAnswerRequestBody struct {
	QuestionSha1 string `json:"questionSha1" binding:"required"`
	AnswerSha1   string `json:"answerSha1" binding:"required"`
	Checked      bool   `json:"checked" binding:"required"`
}

type Class struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func toClassDto(domain *domain.Class) *Class {
	return &Class{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func toClassDtos(domains []*domain.Class) []*Class {
	dtos := make([]*Class, len(domains))

	for i, d := range domains {
		dtos[i] = toClassDto(d)
	}

	return dtos
}

type ClassCreateRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type UserSession struct {
	SessionId    *uuid.UUID     `json:"sessionId"`
	UserId       string         `json:"userId"`
	UserName     string         `json:"userName"`
	RemainingSec int            `json:"remainingSec,omitempty"`
	Result       *SessionResult `json:"result,omitempty"`
}

type QuizSession struct {
	QuizSha1     string         `json:"quizSha1"`
	Name         string         `json:"name"`
	Duration     int            `json:"duration"`
	Filename     string         `json:"filename,omitempty"`
	Version      int            `json:"version,omitempty"`
	CreatedAt    string         `json:"createdAt,omitempty"`
	SessionId    *uuid.UUID     `json:"sessionId,omitempty"`
	UserId       string         `json:"userId,omitempty"`
	UserName     string         `json:"userName,omitempty"`
	RemainingSec int            `json:"remainingSec,omitempty"`
	Result       *SessionResult `json:"result,omitempty"`
	UserSessions []*UserSession `json:"userSessions,omitempty"`
}

func toQuizSession(domain *domain.QuizSession, userId string) *QuizSession {
	session := QuizSession{
		QuizSha1:  domain.QuizSha1,
		Name:      domain.Name,
		Duration:  domain.Duration,
		Filename:  domain.Filename,
		Version:   domain.Version,
		CreatedAt: domain.CreatedAt,
	}

	if len(domain.UserSessions) > 0 {
		for _, userSession := range domain.UserSessions {
			if userId == userSession.UserId {
				session.SessionId = &userSession.SessionId
				session.UserId = userSession.UserId
				session.UserName = userSession.UserName
				session.RemainingSec = userSession.RemainingSec

				if userSession.Result != nil {
					session.Result = &SessionResult{
						GoodAnswer:  userSession.Result.GoodAnswer,
						TotalAnswer: userSession.Result.TotalAnswer,
					}
				}
			} else {
				session.UserSessions = append(session.UserSessions, toUserSession(userSession))
			}
		}
	}

	return &session
}

func toUserSession(domain *domain.UserSession) *UserSession {
	return &UserSession{
		SessionId:    &domain.SessionId,
		UserId:       domain.UserId,
		UserName:     domain.UserName,
		RemainingSec: domain.RemainingSec,
		Result: &SessionResult{
			GoodAnswer:  domain.Result.GoodAnswer,
			TotalAnswer: domain.Result.TotalAnswer,
		},
	}
}

func toQuizSessionDtos(domains []*domain.QuizSession, userId string) []*QuizSession {
	dtos := make([]*QuizSession, len(domains))

	for i, d := range domains {
		dtos[i] = toQuizSession(d, userId)
	}

	return dtos
}

type QuizSessionDetail struct {
	SessionId    uuid.UUID      `json:"sessionId"`
	UserId       string         `json:"userId"`
	RemainingSec int            `json:"remainingSec"`
	Result       *SessionResult `json:"result,omitempty"`
	QuizSha1     string         `json:"quizSha1"`
	Name         string         `json:"name"`
	Questions    []QuizQuestion `json:"questions"`
}

func (qd *QuizSessionDetail) setSha1AndName(sha1 string, name string) {
	qd.QuizSha1 = sha1
	qd.Name = name
}

func (qd *QuizSessionDetail) setQuestions(questions []QuizQuestion) {
	qd.Questions = questions
}

func toQuizSessionDetail(d *domain.QuizSessionDetail) *QuizSessionDetail {
	dto := &QuizSessionDetail{
		SessionId:    d.SessionId,
		UserId:       d.UserId,
		RemainingSec: d.RemainingSec,
	}

	if d.Result != nil {
		dto.Result = &SessionResult{
			GoodAnswer:  d.Result.GoodAnswer,
			TotalAnswer: d.Result.TotalAnswer,
		}
	}

	mapQuizInfos(d, dto)

	return dto
}
