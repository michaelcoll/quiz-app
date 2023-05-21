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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/db"
)

const (
	aud           = "aud"
	sub           = "103275817862301231842"
	expStr        = "1684494062"
	expInStr      = "3591"
	expIn         = 3591
	exp           = 1684494062
	firstName     = "Cordell"
	lastName      = "Walker"
	email         = "cordell.walker@texas-ranger.com"
	emailVerified = "true"
)

func getDBConnection(t *testing.T, dropBeforeConnect bool) *sql.DB {
	if dropBeforeConnect {
		// Deleting previous database
		if _, err := os.Stat("data"); err == nil {
			err := os.RemoveAll("data")
			if err != nil {
				assert.Failf(t, "Fail to delete data folder", "%v", err)
			}
		}
	}

	return db.Init("data")
}
