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

	"github.com/school-by-hiit/quiz-app/internal/back/domain/model"
)

type QuizRepository interface {
	Connect()
	Close()

	FindLatestVersionByFilename(ctx context.Context, filename string) (model.Quiz, error)
	FindAllActive(ctx context.Context) ([]model.Quiz, error)
	Create(ctx context.Context, quiz model.Quiz) error
	Update(ctx context.Context, quiz model.Quiz) error
	Deactivate(ctx context.Context, sha1 string) error
}
