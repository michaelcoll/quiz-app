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
	assert.Equal(t, 15*60, actual.Duration)
	assert.Equal(t, 3, len(actual.Questions))
	assert.Equal(t, 4, len(actual.Questions["8e713df4a80094c5708dc4a1a2a1725643aa375f"].Answers))
	assert.Equal(t, "Version Control System", actual.Questions["8e713df4a80094c5708dc4a1a2a1725643aa375f"].Answers["eb3352743a553af25829c32b2492c1a41a739f1e"].Content)
	assert.Equal(t, true, actual.Questions["8e713df4a80094c5708dc4a1a2a1725643aa375f"].Answers["eb3352743a553af25829c32b2492c1a41a739f1e"].Valid)
	assert.Equal(t, false, actual.Questions["0340bede2f41b9b2ce12b867cc5bf0cb1bd4eabd"].Answers["332b7ca50a406b2337e339332f66f3676d885fef"].Valid)
	assert.Equal(t, false, actual.Questions["0340bede2f41b9b2ce12b867cc5bf0cb1bd4eabd"].Answers["22128893c69197141a17149b7de81419aca57e67"].Valid)
	assert.Equal(t, "Question with a `command` ?", actual.Questions["37f33e32ba7d312df5de7c0bcf03ced422a3413b"].Content)
	assert.Equal(t, "command1\ncommand2", actual.Questions["37f33e32ba7d312df5de7c0bcf03ced422a3413b"].Code)
	assert.Equal(t, "shell", actual.Questions["37f33e32ba7d312df5de7c0bcf03ced422a3413b"].CodeLanguage)
}

func Test_extractQuizNameAndDuration(t *testing.T) {
	content := "# Marvel Universe (duration: 14min)"
	name, duration, err := extractQuizNameAndDuration(content)
	if err != nil {
		assert.Failf(t, "Fail to extract quiz name", "%v", err)
	}

	assert.Equal(t, "Marvel Universe", name)
	assert.Equal(t, 14*60, duration)

	content = "# Marvel Universe"
	_, _, err = extractQuizNameAndDuration(content)
	if err == nil {
		assert.Failf(t, "extract quiz name should have failed", "")
	}
}
