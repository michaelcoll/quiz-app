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

	"github.com/google/uuid"
)

//go:generate mockery --name QuizRepository
type QuizRepository interface {
	FindBySha1(ctx context.Context, sha1 string) (*Quiz, error)
	FindFullBySha1(ctx context.Context, sha1 string) (*Quiz, error)
	FindLatestVersionByFilename(ctx context.Context, filename string) (*Quiz, error)
	FindAllActive(ctx context.Context, limit uint16, offset uint16) ([]*Quiz, error)
	CountAllActive(ctx context.Context) (uint32, error)
	Create(ctx context.Context, quiz *Quiz) error
	ActivateOnlyVersion(ctx context.Context, filename string, version int) error

	FindAllSessions(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16) ([]*Session, error)
	CountAllSessions(ctx context.Context, quizActive bool, userId string) (uint32, error)
	StartSession(ctx context.Context, userId string, quizSha1 string) (uuid.UUID, error)
	AddSessionAnswer(ctx context.Context, sessionUuid uuid.UUID, userId string, questionSha1 string, answerSha1 string, checked bool) error
}

//go:generate mockery --name AuthRepository
type AuthRepository interface {
	CacheToken(token *IdToken) error
	FindTokenByTokenStr(tokenStr string) (*IdToken, error)
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindUserById(ctx context.Context, id string) (*User, error)
	FindAllUser(ctx context.Context) ([]*User, error)
	CreateOrReplaceUser(ctx context.Context, user *User) error
	UpdateUserActive(ctx context.Context, id string, active bool) error
	UpdateUserRole(ctx context.Context, userId string, role Role) error
}

//go:generate mockery --name ClassRepository
type ClassRepository interface {
	FindAll(ctx context.Context, limit uint16, offset uint16) ([]*Class, error)
	CountAll(ctx context.Context) (uint32, error)
}
