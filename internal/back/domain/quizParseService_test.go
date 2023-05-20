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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	body, err := os.ReadFile("quiz.md")
	if err != nil {
		assert.Failf(t, "Fail to read quiz.md file", "%v", err)
	}

	s := NewQuizService(nil)

	actual, err := s.Parse("quiz.md", string(body))
	if err != nil {
		assert.Failf(t, "Fail to parse", "%v", err)
	}

	assert.Equal(t, "Version Control System", actual.Name)
	assert.Equal(t, 3, len(actual.Questions))
	assert.Equal(t, 3, len(actual.Questions["cb4d6d8a2f29f188503607fe26acee8e3786e63f"].Answers))
	assert.Equal(t, "Version Control System", actual.Questions["cb4d6d8a2f29f188503607fe26acee8e3786e63f"].Answers["eb3352743a553af25829c32b2492c1a41a739f1e"].Content)
	assert.Equal(t, true, actual.Questions["cb4d6d8a2f29f188503607fe26acee8e3786e63f"].Answers["eb3352743a553af25829c32b2492c1a41a739f1e"].Valid)
	assert.Equal(t, false, actual.Questions["0340bede2f41b9b2ce12b867cc5bf0cb1bd4eabd"].Answers["332b7ca50a406b2337e339332f66f3676d885fef"].Valid)
	assert.Equal(t, false, actual.Questions["0340bede2f41b9b2ce12b867cc5bf0cb1bd4eabd"].Answers["22128893c69197141a17149b7de81419aca57e67"].Valid)
	assert.Equal(t, "Question with a `command` ?\n```shell\ncommand \n```", actual.Questions["0340bede2f41b9b2ce12b867cc5bf0cb1bd4eabd"].Content)
}
