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
	"github.com/spf13/viper"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/db"
)

type MaintenanceDBRepository struct {
	domain.MaintenanceRepository

	w *ConnectionWrapper
}

func NewMaintenanceRepository(w *ConnectionWrapper) *MaintenanceDBRepository {
	return &MaintenanceDBRepository{w: w}
}

func (r *MaintenanceDBRepository) Dump() (fullDbLocation string, err error) {
	dbLocation := viper.GetString("db-location")

	err = r.w.Close()
	if err != nil {
		return "", err
	}

	return db.GetDatabaseFullPath(dbLocation), nil
}
