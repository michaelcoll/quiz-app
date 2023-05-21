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

import "time"

type Quiz struct {
	Sha1 string

	Filename  string
	Name      string
	Version   int
	CreatedAt string
	Active    bool
	Questions map[string]QuizQuestion
}

type QuizQuestion struct {
	Sha1 string

	Content string
	Answers map[string]QuizQuestionAnswer
}

type QuizQuestionAnswer struct {
	Sha1 string

	Content string
	Valid   bool
}

type SyncStats struct {
	Created int
	Updated int
}

type Role int8

const (
	Admin   Role = 1
	Teacher Role = 2
	Student Role = 3
)

type User struct {
	Id        string
	Email     string
	Firstname string
	Lastname  string
	Active    bool
	Role      Role
}

type TokenProvenance int8

const (
	Cache = 1
	Api   = 2
)

type AccessToken struct {
	Aud         string
	Sub         string
	Exp         time.Time
	ExpiresIn   int
	Email       string
	Provenance  TokenProvenance
	OpaqueToken string
}
