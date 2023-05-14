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

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/stretchr/testify/assert"
)

func Test_readFileContent(t *testing.T) {

	storage := memory.NewStorage()
	fs := memfs.New()

	_, err := git.Clone(storage, fs, &git.CloneOptions{
		URL: "https://github.com/school-by-hiit/quizz-app.git",
	})
	if err != nil {
		assert.Fail(t, "Can't connect", "%v", err)
	}

	actual, err := readFileContent(fs, "marvel-universe.quizz.md")
	if err != nil {
		assert.Fail(t, "Can't read repo file : marvel-universe.quizz.md", "%v", err)
	}

	expected, err := os.ReadFile("../../../../marvel-universe.quizz.md")
	if err != nil {
		assert.Fail(t, "Can't read local file : marvel-universe.quizz.md", "%v", err)
	}

	assert.Equal(t, string(expected), actual)
}

func TestScanGitRepo(t *testing.T) {

	s := New(nil)

	quizzes, err := s.ScanGitRepo("https://github.com/school-by-hiit/quizz-app.git", "")
	if err != nil {
		assert.Fail(t, "Can't scan repo", "%v", err)
	}

	assert.Len(t, quizzes, 1)
	assert.Equal(t, "Marvel Universe", quizzes[0].Name)
	assert.Equal(t, "marvel-universe.quizz.md", quizzes[0].Filename)
	assert.Equal(t, "fccc28a245ee3e92791ec9395d3a3791d17090da", quizzes[0].Sha1)
	assert.Len(t, quizzes[0].Questions, 7)
}
