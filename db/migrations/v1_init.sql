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

CREATE TABLE quizz
(
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE quizz_question
(
    id       TEXT PRIMARY KEY,
    quizz_id TEXT,

    CONSTRAINT quizz_fk FOREIGN KEY (quizz_id) REFERENCES quizz (id)
);

CREATE TABLE quizz_question_answer
(
    id                TEXT PRIMARY KEY,
    quizz_question_id TEXT,
    valid             INTEGER,

    CONSTRAINT quizz_question_fk FOREIGN KEY (quizz_question_id) REFERENCES quizz_question (id)
);
