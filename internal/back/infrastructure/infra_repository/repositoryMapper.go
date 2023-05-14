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

package infra_repository

import (
	"github.com/school-by-hiit/quiz-app/internal/back/domain/model"
	"github.com/school-by-hiit/quiz-app/internal/back/infrastructure/sqlc"
)

func toDomain(entity sqlc.Quiz) model.Quiz {
	return model.Quiz{
		Sha1:      entity.Sha1,
		Filename:  entity.Filename,
		Name:      entity.Name,
		Version:   int(entity.Version),
		CreatedAt: entity.CreatedAt,
	}
}
