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

	"github.com/google/uuid"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type ClassDBRepository struct {
	domain.ClassRepository

	q *sqlc.Queries
}

func NewClassRepository(c *sql.DB) *ClassDBRepository {
	return &ClassDBRepository{q: sqlc.New(c)}
}

func (r *ClassDBRepository) FindAll(ctx context.Context, limit uint16, offset uint16) ([]*domain.Class, error) {
	classes, err := r.q.FindAllClasses(ctx, sqlc.FindAllClassesParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	return r.toClassArray(classes), nil
}

func (r *ClassDBRepository) CountAll(ctx context.Context) (uint32, error) {
	count, err := r.q.CountAllClasses(ctx)
	if err != nil {
		return 0, err
	}

	return uint32(count), nil
}

func (r *ClassDBRepository) toClass(entity sqlc.StudentClass) *domain.Class {
	return &domain.Class{
		Id:   entity.Uuid,
		Name: entity.Name,
	}
}

func (r *ClassDBRepository) toClassArray(entities []sqlc.StudentClass) []*domain.Class {
	domains := make([]*domain.Class, len(entities))

	for i, entity := range entities {
		domains[i] = r.toClass(entity)
	}

	return domains
}

func (r *ClassDBRepository) CreateOrReplace(ctx context.Context, class *domain.Class) error {
	return r.q.CreateOrReplaceClass(ctx, sqlc.CreateOrReplaceClassParams{
		Uuid: class.Id,
		Name: class.Name,
	})
}

func (r *ClassDBRepository) Delete(ctx context.Context, classId uuid.UUID) error {
	return r.q.DeleteClassById(ctx, classId)
}

func (r *ClassDBRepository) ExistsById(ctx context.Context, classId uuid.UUID) bool {
	count, err := r.q.CountClassById(ctx, classId)
	if err != nil {
		return false
	}

	return count == 1
}
