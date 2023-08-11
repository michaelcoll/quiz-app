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
	sub             = 4242424
	subStr          = "4242424"
	login           = "cordell.walker"
	name            = "Cordell Walker"
	picture         = "https://avatars.githubusercontent.com/u/4242424?v=4"
	sha1Quiz1       = "c152b2d0a2509a82ea5e8a6ae22fea55c7221002"
	sha1Quiz2       = "770ef94955911a984e3d4925d2419c44d3aaca28"
	quizName1       = "Marvel Universe"
	quizName2       = "Video games"
	quizDuration1   = 840
	quizDuration2   = 1200
	quizFilename1   = "marvel-universe.quiz.md"
	quizFilename2   = "video-games.quiz.md"
	quizVersion1    = 1
	quizVersion2    = 2
	quizCreatedAt1  = "2023-06-16T18:26:54+02:00"
	quizCreatedAt2  = "2023-06-16T18:26:54+02:00"
	userId1         = "103275817862301231842"
	remainingSec1   = 0
	checkedAnswers1 = 1
	results1        = 25
	userId2         = "103275817862301234242"
	remainingSec2   = 0
	checkedAnswers2 = 2
	results2        = 25
)

func getDBConnection(t *testing.T, dropBeforeConnect bool) *sql.DB {
	if dropBeforeConnect {
		// Deleting the previous database
		if _, err := os.Stat("data"); err == nil {
			err := os.RemoveAll("data")
			if err != nil {
				assert.Failf(t, "Fail to delete data folder", "%v", err)
			}
		}
	}

	return db.Init("data")
}
