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
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type UserDBRepository struct {
	domain.UserRepository

	w  *ConnectionWrapper
	uc *cache.Cache
}

func NewUserRepository(w *ConnectionWrapper) *UserDBRepository {
	userCache := cache.New(30*time.Minute, 10*time.Minute)
	return &UserDBRepository{w: w, uc: userCache}
}

func (r *UserDBRepository) FindActiveUserById(ctx context.Context, id string) (*domain.User, error) {

	if user, found := r.uc.Get(id); found {
		return user.(*domain.User), nil
	}

	entity, err := r.w.queries().FindActiveUserById(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user := r.toUser(entity)

	r.uc.Set(id, user, cache.DefaultExpiration)

	return user, nil
}

func (r *UserDBRepository) FindUserById(ctx context.Context, id string) (*domain.User, error) {

	entity, err := r.w.queries().FindUserById(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user := r.toUser(entity)

	return user, nil
}

func (r *UserDBRepository) CreateOrReplaceUser(ctx context.Context, user *domain.User) error {
	err := r.w.queries().CreateOrReplaceUser(ctx, sqlc.CreateOrReplaceUserParams{
		ID:      user.Id,
		Login:   user.Login,
		Name:    user.Name,
		Picture: user.Picture,
		RoleID:  int8(user.Role),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserDBRepository) UpdateUserRole(ctx context.Context, userId string, role domain.Role) error {
	err := r.w.queries().UpdateUserRole(ctx, sqlc.UpdateUserRoleParams{
		ID:     userId,
		RoleID: int8(role),
	})
	if err != nil {
		return err
	}

	r.uc.Delete(userId)

	return nil
}

func (r *UserDBRepository) FindAllUser(ctx context.Context) ([]*domain.User, error) {
	entities, err := r.w.queries().FindAllUser(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(entities))

	for i, entity := range entities {
		users[i] = r.toUser(entity)
	}

	return users, nil
}

func (r *UserDBRepository) UpdateUserActive(ctx context.Context, id string, active bool) error {
	return r.w.queries().UpdateUserActive(ctx, sqlc.UpdateUserActiveParams{
		Active: active,
		ID:     id,
	})
}

func (r *UserDBRepository) UpdateUserInfo(ctx context.Context, user *domain.User) error {
	r.uc.Delete(user.Id)

	return r.w.queries().UpdateUserInfo(ctx, sqlc.UpdateUserInfoParams{
		Login:   user.Login,
		Name:    user.Name,
		Picture: user.Picture,
		ID:      user.Id,
	})
}

func (r *UserDBRepository) AssignUserToClass(ctx context.Context, userId string, classId uuid.UUID) error {
	return r.w.queries().AssignUserToClass(ctx, sqlc.AssignUserToClassParams{
		ClassUuid: classId,
		ID:        userId,
	})
}

func (r *UserDBRepository) toUser(entity sqlc.UserClassView) *domain.User {
	d := &domain.User{
		Id:      entity.ID,
		Login:   entity.Login,
		Name:    entity.Name,
		Picture: entity.Picture,
		Active:  entity.Active,
		Role:    r.toRole(entity.RoleID),
	}

	if entity.ClassName != "" {
		d.Class = r.toClass(entity.ClassUuid, entity.ClassName)
	}

	return d
}

func (r *UserDBRepository) toRole(entity int8) domain.Role {
	switch entity {
	case int8(domain.Admin):
		return domain.Admin
	case int8(domain.Teacher):
		return domain.Teacher
	case int8(domain.Student):
		return domain.Student
	}

	return 0
}

func (r *UserDBRepository) toClass(classId uuid.UUID, className string) *domain.Class {
	return &domain.Class{
		Id:   classId,
		Name: className,
	}
}
