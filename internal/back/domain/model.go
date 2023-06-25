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
	"time"

	"github.com/google/uuid"
)

type QuizInfos interface {
	GetSha1AndName() (string, string)
	GetQuestions() map[string]QuizQuestion
}

type Quiz struct {
	Sha1 string

	Filename  string
	Name      string
	Version   int
	CreatedAt string
	Active    bool
	Duration  int
	Questions map[string]QuizQuestion
}

func (q *Quiz) GetSha1AndName() (string, string) {
	return q.Sha1, q.Name
}

func (q *Quiz) GetQuestions() map[string]QuizQuestion {
	return q.Questions
}

type QuizQuestion struct {
	Sha1 string

	Content string
	Answers map[string]QuizQuestionAnswer
}

type QuizQuestionAnswer struct {
	Sha1 string

	Content string
	Checked bool
	Valid   bool
}

type SyncStats struct {
	Created int
	Updated int
}

type Role int8

const (
	NoRole  Role = 0
	Admin   Role = 1
	Teacher Role = 2
	Student Role = 3
)

type User struct {
	Id string

	Email     string
	Firstname string
	Lastname  string
	Active    bool
	Role      Role
	Class     *Class
}

type TokenProvenance int8

const (
	Cache = 1
	Parse = 2
)

type IdToken struct {
	Sub string

	Aud         string
	Exp         time.Time
	ExpiresIn   int
	Email       string
	FirstName   string
	LastName    string
	Provenance  TokenProvenance
	JwtStrToken string
}

type SessionResult struct {
	GoodAnswer  int
	TotalAnswer int
}

type Session struct {
	Id uuid.UUID

	QuizSha1     string
	QuizName     string
	QuizActive   bool
	UserId       string
	UserName     string
	RemainingSec int
	Result       *SessionResult
}

type Class struct {
	Id uuid.UUID

	Name string
}

type UserSession struct {
	SessionId uuid.UUID
	UserId    string

	UserName     string
	RemainingSec int
	Result       *SessionResult
}

type QuizSession struct {
	QuizSha1 string

	Name         string
	Duration     int
	Filename     string
	Version      int
	CreatedAt    string
	UserSessions []*UserSession
}

type QuizSessionDetail struct {
	SessionId uuid.UUID

	UserId       string
	RemainingSec int
	Result       *SessionResult
	QuizSha1     string
	Name         string
	Questions    map[string]QuizQuestion
}

func (qd *QuizSessionDetail) GetSha1AndName() (string, string) {
	return qd.QuizSha1, qd.Name
}

func (qd *QuizSessionDetail) GetQuestions() map[string]QuizQuestion {
	return qd.Questions
}
