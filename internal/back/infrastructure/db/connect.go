/*
 * Copyright (c) 2022-2023 Michaël COLL.
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

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/school-by-hiit/quizz-app/internal/back/domain/consts"
)

func Connect(readOnly bool, baseLocation string) *sql.DB {
	db, err := sql.Open("sqlite3", getDBUrl(readOnly, baseLocation))
	if err != nil {
		log.Fatalf("Can't open database %v", err)
	}

	return db
}

func getDBUrl(readOnly bool, baseLocation string) string {

	var options string
	if readOnly {
		options = "cache=shared&mode=ro"
	} else {
		options = "cache=shared&mode=rwc&_auto_vacuum=full&_journal_mode=WAL"
	}

	return fmt.Sprintf("file:%s/%s?%s", baseLocation, consts.DatabaseName, options)
}
