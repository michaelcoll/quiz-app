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
	"database/sql"

	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/db"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type ConnectionWrapper struct {
	dbLocation string
	c          *sql.DB
	isClosed   bool
}

func NewConnectionWrapper(dbLocation string) *ConnectionWrapper {
	return &ConnectionWrapper{dbLocation: dbLocation, c: db.Init(dbLocation)}
}

func (w *ConnectionWrapper) queries() *sqlc.Queries {
	if w.isClosed {
		w.c = db.Init(w.dbLocation)
		w.isClosed = false
	}

	return sqlc.New(w.c)
}

func (w *ConnectionWrapper) Close() error {
	defer func() {
		w.isClosed = true
	}()
	return w.c.Close()
}
