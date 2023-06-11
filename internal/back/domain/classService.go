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

type ClassService struct {
	r ClassRepository
}

func NewClassService(classRepository ClassRepository) ClassService {
	return ClassService{r: classRepository}
}

func (s *ClassService) FindAllClasses(ctx context.Context, limit uint16, offset uint16) ([]*Class, uint32, error) {
	classes, err := s.r.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.r.CountAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	return classes, total, nil
}

func (s *ClassService) Create(ctx context.Context, name string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return s.r.CreateOrReplace(ctx, &Class{
		Id:   id,
		Name: name,
	})
}

func (s *ClassService) Update(ctx context.Context, class *Class) error {

	if !s.r.ExistsById(ctx, class.Id) {
		return Errorf(NotFound, "class with id (%s) not found", class.Id)
	}

	return s.r.CreateOrReplace(ctx, class)
}

func (s *ClassService) Delete(ctx context.Context, id uuid.UUID) error {

	if !s.r.ExistsById(ctx, id) {
		return Errorf(NotFound, "class with id (%s) not found", id)
	}

	return s.r.Delete(ctx, id)
}

func (s *ClassService) CreateQuizClassVisibility(ctx context.Context, quizSha1 string, classId uuid.UUID) error {
	return s.r.CreateQuizClassVisibility(ctx, quizSha1, classId)
}

func (s *ClassService) DeleteQuizClassVisibility(ctx context.Context, quizSha1 string, classId uuid.UUID) error {
	return s.r.DeleteQuizClassVisibility(ctx, quizSha1, classId)
}
