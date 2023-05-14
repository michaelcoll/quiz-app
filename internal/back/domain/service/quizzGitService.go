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
	"io"
	"regexp"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/school-by-hiit/quizz-app/internal/back/domain/model"
)

func (s *QuizzService) ScanGitRepo(url string, token string) ([]model.Quizz, error) {
	storage := memory.NewStorage()
	fs := memfs.New()

	var option git.CloneOptions
	if len(token) == 0 {
		option = git.CloneOptions{
			URL: url,
		}
	} else {
		option = git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "42", // yes, this can be anything except an empty string
				Password: token,
			},
			URL: url,
		}
	}

	_, err := git.Clone(storage, fs, &option)
	if err != nil {
		return nil, err
	}

	dir, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var quizzes []model.Quizz

	r := regexp.MustCompile(`.*\.quizz\.md`)
	for _, fileInfo := range dir {

		if r.MatchString(fileInfo.Name()) {
			content, err := readFileContent(fs, fileInfo.Name())
			if err != nil {
				return nil, err
			}

			quizz, err := s.Parse(fileInfo.Name(), content)
			if err != nil {
				return nil, err
			}

			quizzes = append(quizzes, quizz)
		}
	}

	return quizzes, nil
}

func readFileContent(fs billy.Filesystem, filename string) (string, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, file)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
