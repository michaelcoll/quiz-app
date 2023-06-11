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

type UserService struct {
	r UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return UserService{r: userRepository}
}

func (s *UserService) FindUserById(ctx context.Context, id string) (*User, error) {
	user, err := s.r.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Errorf(NotFound, "user with id '%s' not found", id)
	}

	return user, nil
}

func (s *UserService) FindAllUser(ctx context.Context) ([]*User, error) {
	return s.r.FindAllUser(ctx)
}

func (s *UserService) DeactivateUser(ctx context.Context, id string) error {
	return s.r.UpdateUserActive(ctx, id, false)
}

func (s *UserService) ActivateUser(ctx context.Context, id string) error {
	return s.r.UpdateUserActive(ctx, id, true)
}

func (s *UserService) AssignUserToClass(ctx context.Context, userId string, classId uuid.UUID) error {
	return s.r.AssignUserToClass(ctx, userId, classId)
}
