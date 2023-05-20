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
	"context"
	"database/sql"
	"errors"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type AuthDBRepository struct {
	domain.AuthRepository

	q *sqlc.Queries
}

func NewAuthRepository(c *sql.DB) *AuthDBRepository {
	return &AuthDBRepository{q: sqlc.New(c)}
}

func (r *AuthDBRepository) FindUserById(ctx context.Context, id string) (*domain.User, error) {
	rows, err := r.q.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	return r.toUser(rows), nil
}

func (r *AuthDBRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := r.q.CreateOrReplaceUser(ctx, sqlc.CreateOrReplaceUserParams{
		ID:        user.Id,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDBRepository) AddRoleToUser(ctx context.Context, userId string, role domain.Role) error {
	err := r.q.AddRoleToUser(ctx, sqlc.AddRoleToUserParams{
		UserID: sql.NullString{
			String: userId,
			Valid:  true,
		},
		RoleID: sql.NullInt64{
			Int64: int64(role),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDBRepository) RemoveAllRoleFromUser(ctx context.Context, userId string) error {
	err := r.q.RemoveAllRoleFromUser(ctx, sql.NullString{
		String: userId,
		Valid:  true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDBRepository) CreateToken(ctx context.Context, token *domain.AccessToken) error {
	err := r.q.CreateOrReplaceToken(ctx, sqlc.CreateOrReplaceTokenParams{
		OpaqueToken: token.OpaqueToken,
		UserID:      token.Sub,
		Expires:     token.Exp,
		Aud:         token.Aud,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDBRepository) FindTokenByTokenStr(ctx context.Context, tokenStr string) (*domain.AccessToken, error) {
	token, err := r.q.FindTokenByTokenStr(ctx, tokenStr)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return r.toAccessToken(token), nil
}

func (r *AuthDBRepository) toUser(entity []sqlc.FindUserByIdRow) *domain.User {

	user := domain.User{
		Roles: []domain.Role{},
	}

	for _, row := range entity {
		user.Id = row.ID
		user.Email = row.Email
		user.Firstname = row.Firstname
		user.Lastname = row.Lastname
		user.Active = intToBool(row.Active)
		role := r.toRole(row.RoleID)
		if role > 0 {
			user.Roles = append(user.Roles, role)
		}
	}

	return &user
}

func (r *AuthDBRepository) toRole(entity sql.NullInt64) domain.Role {
	if !entity.Valid {
		return 0
	}

	switch entity.Int64 {
	case int64(domain.Admin):
		return domain.Admin
	case int64(domain.Teacher):
		return domain.Teacher
	case int64(domain.Student):
		return domain.Student
	}

	return 0
}

func (r *AuthDBRepository) toAccessToken(entity sqlc.FindTokenByTokenStrRow) *domain.AccessToken {
	return &domain.AccessToken{
		Aud:         entity.Aud,
		Sub:         entity.UserID,
		Exp:         entity.Expires,
		Email:       entity.Email,
		Provenance:  domain.Cache,
		OpaqueToken: entity.OpaqueToken,
	}
}
