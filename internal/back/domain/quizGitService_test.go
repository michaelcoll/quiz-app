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

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_readFileContent(t *testing.T) {

	storage := memory.NewStorage()
	fs := memfs.New()

	_, err := git.Clone(storage, fs, &git.CloneOptions{
		URL: "../../../.",
	})
	if err != nil {
		assert.Fail(t, "Can't connect", "%v", err)
	}

	actual, err := readFileContent(fs, "marvel-universe.quiz.md")
	if err != nil {
		assert.Fail(t, "Can't read repo file : marvel-universe.quiz.md", "%v", err)
	}

	expected, err := os.ReadFile("../../../marvel-universe.quiz.md")
	if err != nil {
		assert.Fail(t, "Can't read local file : marvel-universe.quiz.md", "%v", err)
	}

	assert.Equal(t, string(expected), actual)
}

func TestScanGitRepo(t *testing.T) {

	s := NewQuizService(nil)

	viper.Set("repository-url", "../../../.")
	viper.Set("token", "")

	quizzes, err := s.ScanGitRepo()
	if err != nil {
		assert.Fail(t, "Can't scan repo", "%v", err)
	}

	assert.Len(t, quizzes, 2)
	assert.Equal(t, "Marvel Universe", quizzes[0].Name)
	assert.Equal(t, "marvel-universe.quiz.md", quizzes[0].Filename)
	assert.Equal(t, "c152b2d0a2509a82ea5e8a6ae22fea55c7221002", quizzes[0].Sha1)
	assert.Len(t, quizzes[0].Questions, 7)
}
