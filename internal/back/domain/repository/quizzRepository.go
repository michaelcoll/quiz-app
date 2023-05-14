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

package repository

import (
	"context"

	"github.com/school-by-hiit/quizz-app/internal/back/domain/model"
)

type QuizzRepository interface {
	Connect()
	Close()

	FindLatestVersionByFilename(ctx context.Context, filename string) (model.Quizz, error)
	FindAllActive(ctx context.Context) ([]model.Quizz, error)
	Create(ctx context.Context, quizz model.Quizz) error
	Update(ctx context.Context, quizz model.Quizz) error
	Deactivate(ctx context.Context, sha1 string) error
}
