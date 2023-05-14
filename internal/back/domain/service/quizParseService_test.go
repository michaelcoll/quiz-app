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

package service

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

	s := New(nil)

	actual, err := s.Parse("quiz.md", string(body))
	if err != nil {
		assert.Failf(t, "Fail to parse", "%v", err)
	}

	assert.Equal(t, "Version Control System", actual.Name)
	assert.Equal(t, 3, len(actual.Questions))
	assert.Equal(t, 3, len(actual.Questions[0].Answers))
	assert.Equal(t, "Version Control System", actual.Questions[0].Answers[0].Content)
	assert.Equal(t, true, actual.Questions[0].Answers[0].Valid)
	assert.Equal(t, true, actual.Questions[2].Answers[1].Valid)
	assert.Equal(t, false, actual.Questions[2].Answers[2].Valid)
	assert.Equal(t, "Question with a `command` ?\n```shell\ncommand \n```", actual.Questions[2].Content)
}
