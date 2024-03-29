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

// Code generated by sqlc-addon. DO NOT EDIT.
// versions:
//   sqlc-addon v1.9.4

package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fatih/color"
)

const testVersionTableExists = `
SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='db_version';
`

const initSql = `
CREATE TABLE db_version
(
    version_number INTEGER NOT NULL
);
INSERT INTO db_version (version_number) VALUES (0);
`

const selectVersionStmt = `
SELECT version_number FROM db_version;
`

const updateVersionStmt = `
UPDATE db_version
SET version_number = ?
WHERE 1;
`

const v1Init = `
CREATE TABLE quiz
(
    sha1       TEXT PRIMARY KEY,
    name       TEXT    NOT NULL,
    filename   TEXT    NOT NULL,
    version    INTEGER NOT NULL DEFAULT 1,
    active     INTEGER NOT NULL DEFAULT 1,
    created_at TEXT    NOT NULL,
    duration   INTEGER NOT NULL,

    UNIQUE (filename, version)
);

CREATE TABLE quiz_question
(
    sha1          TEXT PRIMARY KEY,
    position      INTEGER NOT NULL,
    content       TEXT    NOT NULL,
    code          TEXT,
    code_language TEXT
);

CREATE TABLE quiz_question_quiz
(
    quiz_sha1     TEXT NOT NULL,
    question_sha1 TEXT NOT NULL,

    PRIMARY KEY (quiz_sha1, question_sha1),
    FOREIGN KEY (quiz_sha1) REFERENCES quiz (sha1),
    FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1)
);

CREATE TABLE quiz_answer
(
    sha1    TEXT PRIMARY KEY,
    valid   INTEGER NOT NULL,
    content TEXT    NOT NULL
);

CREATE TABLE quiz_question_answer
(
    question_sha1 TEXT NOT NULL,
    answer_sha1   TEXT NOT NULL,

    FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1),
    FOREIGN KEY (answer_sha1) REFERENCES quiz_answer (sha1)
);

CREATE TABLE role
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO role (id, name)
VALUES (1, 'Admin');
INSERT INTO role (id, name)
VALUES (2, 'Teacher');
INSERT INTO role (id, name)
VALUES (3, 'Student');

CREATE TABLE student_class
(
    uuid TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE user
(
    id         TEXT PRIMARY KEY,
    login      TEXT    NOT NULL,
    name       TEXT    NOT NULL,
    picture    TEXT    NOT NULL,
    active     INTEGER NOT NULL DEFAULT 1,
    role_id    INTEGER NOT NULL,
    class_uuid TEXT,

    FOREIGN KEY (role_id) REFERENCES role (id),
    FOREIGN KEY (class_uuid) REFERENCES student_class (uuid) ON DELETE SET NULL
);

CREATE TABLE session
(
    uuid       TEXT PRIMARY KEY,
    quiz_sha1  TEXT      NOT NULL,
    user_id    TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (quiz_sha1, user_id),
    FOREIGN KEY (quiz_sha1) REFERENCES quiz (sha1),
    FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE session_answer
(
    session_uuid  TEXT    NOT NULL,
    question_sha1 TEXT    NOT NULL,
    answer_sha1   TEXT    NOT NULL,
    checked       INTEGER NOT NULL,

    PRIMARY KEY (session_uuid, question_sha1, answer_sha1),
    FOREIGN KEY (session_uuid) REFERENCES session (uuid),
    FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1),
    FOREIGN KEY (answer_sha1) REFERENCES quiz_answer (sha1)
);

CREATE VIEW quiz_answer_count_view
AS
SELECT q.sha1   AS quiz_sha1,
       COUNT(1) AS checked_answers
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question_answer qqa ON qqq.question_sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
WHERE qa.valid = 1
GROUP BY q.sha1;

CREATE VIEW session_response_view
AS
SELECT qqq.quiz_sha1,
       qqq.question_sha1,
       qqa.answer_sha1,
       s.uuid  AS session_uuid,
       s.user_id,
       sa.checked,
       CASE
           WHEN checked IS NOT NULL
               THEN CASE
                        WHEN qa.valid == sa.checked
                            THEN 1
                        ELSE 0
               END
           END AS result
FROM quiz_question_quiz qqq
         JOIN quiz_question_answer qqa ON qqq.question_sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
         LEFT JOIN session s ON qqq.quiz_sha1 = s.quiz_sha1
         LEFT JOIN session_answer sa
                   ON sa.session_uuid = s.uuid
                       AND qa.sha1 = sa.answer_sha1
                       AND sa.question_sha1 = qqq.question_sha1
                       AND sa.answer_sha1 = qqa.answer_sha1;

CREATE VIEW session_view
AS
SELECT s.uuid                                                                                       AS uuid,
       q.sha1                                                                                       AS quiz_sha1,
       q.name                                                                                       AS quiz_name,
       q.active                                                                                     AS quiz_active,
       u.id                                                                                         AS user_id,
       u.name                                                                                       AS user_name,
       u.picture                                                                                    AS user_picture,
       CAST(MAX(q.duration - (STRFTIME('%s', 'now') - STRFTIME('%s', s.created_at)), 0) AS INTEGER) AS remaining_sec,
       checked_answers,
       SUM(srv.result)                                                                              AS results
FROM session s
         JOIN quiz q ON q.sha1 = s.quiz_sha1
         JOIN user u ON u.id = s.user_id
         JOIN quiz_answer_count_view qacv ON s.quiz_sha1 = qacv.quiz_sha1
         JOIN session_response_view srv ON s.uuid = srv.session_uuid
GROUP BY s.uuid, q.sha1, q.name, q.active, u.id, u.name, u.picture;

CREATE TRIGGER verify_remaining_time_create
    BEFORE INSERT
    ON session_answer
BEGIN
    SELECT CASE
               WHEN (SELECT remaining_sec FROM session_view sv WHERE sv.uuid = new.session_uuid) = 0 THEN
                   RAISE(ABORT, 'session is over')
               END;
END;

CREATE TRIGGER verify_remaining_time_update
    BEFORE UPDATE
    ON session_answer
BEGIN
    SELECT CASE
               WHEN (SELECT remaining_sec FROM session_view sv WHERE sv.uuid = new.session_uuid) = 0 THEN
                   RAISE(ABORT, 'session is over')
               END;
END;

CREATE TABLE quiz_class_visibility
(
    class_uuid TEXT NOT NULL,
    quiz_sha1  TEXT NOT NULL,

    PRIMARY KEY (class_uuid, quiz_sha1),
    FOREIGN KEY (class_uuid) REFERENCES student_class (uuid) ON DELETE CASCADE,
    FOREIGN KEY (quiz_sha1) REFERENCES quiz (sha1)
);

CREATE VIEW quiz_session_view
AS
SELECT q.sha1                                                                  AS quiz_sha1,
       q.name                                                                  AS quiz_name,
       q.filename                                                              AS quiz_filename,
       q.version                                                               AS quiz_version,
       q.duration                                                              AS quiz_duration,
       q.created_at                                                            AS quiz_created_at,
       CASE WHEN s.uuid IS NULL THEN '' ELSE s.uuid END                        AS session_uuid,
       CASE WHEN s.user_id IS NULL THEN '' ELSE s.user_id END                  AS user_id,
       CASE WHEN u.name IS NULL THEN '' ELSE u.name END                        AS user_name,
       CASE WHEN u.picture IS NULL THEN '' ELSE u.picture END                  AS user_picture,
       CASE WHEN sc.uuid IS NULL THEN '' ELSE sc.uuid END                      AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END                      AS class_name,
       CASE WHEN sv.remaining_sec IS NULL THEN 0 ELSE sv.remaining_sec END     AS remaining_sec,
       CASE WHEN sv.checked_answers IS NULL THEN 0 ELSE sv.checked_answers END AS checked_answers,
       CASE WHEN sv.results IS NULL THEN 0 ELSE sv.results END                 AS results
FROM quiz q
         LEFT JOIN session s ON q.sha1 = s.quiz_sha1
         LEFT JOIN user u ON s.user_id = u.id
         LEFT JOIN student_class sc ON u.class_uuid = sc.uuid
         LEFT JOIN session_view sv ON q.sha1 = sv.quiz_sha1
WHERE q.active = TRUE;

CREATE VIEW quiz_session_detail_view
AS
SELECT qsv.session_uuid                                          AS session_uuid,
       qsv.user_id                                               AS user_id,
       qsv.remaining_sec                                         AS remaining_sec,
       qsv.quiz_sha1                                             AS quiz_sha1,
       qsv.quiz_name                                             AS quiz_name,
       qsv.quiz_duration                                         AS quiz_duration,
       qsv.checked_answers                                       AS checked_answers,
       qsv.results                                               AS results,
       srv.question_sha1                                         AS question_sha1,
       qq.position                                               AS question_position,
       qq.content                                                AS question_content,
       qq.code                                                   AS question_code,
       qq.code_language                                          AS question_code_language,
       srv.answer_sha1                                           AS answer_sha1,
       qa.content                                                AS answer_content,
       CASE WHEN srv.checked IS NULL THEN 0 ELSE srv.checked END AS answer_checked,
       qa.valid                                                  AS answer_valid
FROM quiz_session_view qsv
         JOIN session_response_view srv ON qsv.session_uuid = srv.session_uuid
         JOIN quiz_question qq ON srv.question_sha1 = qq.sha1
         JOIN quiz_answer qa ON srv.answer_sha1 = qa.sha1
ORDER BY qq.position;

CREATE VIEW user_class_view
AS
SELECT u.id,
       u.login,
       u.name,
       u.picture,
       u.active,
       u.role_id,
       CASE WHEN u.class_uuid IS NULL THEN '' ELSE u.class_uuid END AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END           AS class_name
FROM user u
         LEFT JOIN student_class sc ON u.class_uuid = sc.uuid;

CREATE VIEW quiz_class_view
AS
SELECT q.*,
       CASE WHEN sc.uuid IS NULL THEN '' ELSE sc.uuid END AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END AS class_name
FROM quiz q
         LEFT JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         LEFT JOIN student_class sc ON sc.uuid = qcv.class_uuid
`

const v2FixQuizSessionView = `
DROP VIEW quiz_session_view;

CREATE VIEW quiz_session_view
AS
SELECT q.sha1                                                                  AS quiz_sha1,
       q.name                                                                  AS quiz_name,
       q.filename                                                              AS quiz_filename,
       q.version                                                               AS quiz_version,
       q.duration                                                              AS quiz_duration,
       q.created_at                                                            AS quiz_created_at,
       CASE WHEN s.uuid IS NULL THEN '' ELSE s.uuid END                        AS session_uuid,
       CASE WHEN s.user_id IS NULL THEN '' ELSE s.user_id END                  AS user_id,
       CASE WHEN u.name IS NULL THEN '' ELSE u.name END                        AS user_name,
       CASE WHEN u.picture IS NULL THEN '' ELSE u.picture END                  AS user_picture,
       CASE WHEN sc.uuid IS NULL THEN '' ELSE sc.uuid END                      AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END                      AS class_name,
       CASE WHEN sv.remaining_sec IS NULL THEN 0 ELSE sv.remaining_sec END     AS remaining_sec,
       CASE WHEN sv.checked_answers IS NULL THEN 0 ELSE sv.checked_answers END AS checked_answers,
       CASE WHEN sv.results IS NULL THEN 0 ELSE sv.results END                 AS results
FROM quiz q
         LEFT JOIN session s ON q.sha1 = s.quiz_sha1
         LEFT JOIN user u ON s.user_id = u.id
         LEFT JOIN student_class sc ON u.class_uuid = sc.uuid
         LEFT JOIN session_view sv ON q.sha1 = sv.quiz_sha1 AND s.uuid = sv.uuid
WHERE q.active = TRUE;
`

var migrations = map[int]string{
	1: v1Init,
	2: v2FixQuizSessionView,
}

var migrationVersions = []int{
	1,
	2,
}

type DB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
}

// New creates a new instance of Migrations struct
func New(db DB) *Migrations {
	return &Migrations{db: db}
}

type Migrations struct {
	db DB
}

// Migrate migrates the database using the migration scripts provided
func (m *Migrations) Migrate() {
	initialized, err := m.isInitialized()
	if err != nil {
		log.Fatalf("Can't detect if database is initialized %v", err)
	}
	if initialized {
		version, err := m.getVersion()
		if err != nil {
			log.Fatalf("Can't read database version %v", err)
		}
		if m.needsMigration(version) {
			m.applyMigration(version)
		} else {
			fmt.Printf("%s Database is up to date (v%d)\n",
				color.HiBlueString("i"), version)
		}
	} else {
		m.createDBVersionTable()
		m.applyMigration(0)
	}
}

// needsMigration checks if the database needs to be migrated
func (m *Migrations) needsMigration(currentVersion int) bool {
	return currentVersion < len(migrations)
}

// isInitialized checks if the table db_version is present in the current database
func (m *Migrations) isInitialized() (bool, error) {
	stmt, err := m.db.Prepare(testVersionTableExists)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var tablePresent int
	err = stmt.QueryRow().Scan(&tablePresent)
	if err != nil {
		return false, err
	}

	return tablePresent == 1, nil
}

// getVersion returns the current version of the schema
func (m *Migrations) getVersion() (int, error) {
	stmt, err := m.db.Prepare(selectVersionStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var version int
	err = stmt.QueryRow().Scan(&version)
	if err != nil {
		return 0, err
	}

	return version, nil
}

// applyMigration a migration
func (m *Migrations) createDBVersionTable() {
	_, err := m.db.Exec(initSql)
	if err != nil {
		log.Fatalf("Could not create db_version table %v", err)
	}
}

// applyMigration a migration
func (m *Migrations) applyMigration(fromVersion int) {
	updStmt, err := m.db.Prepare(updateVersionStmt)
	if err != nil {
		log.Fatalf("Could not prepare Stmt : %v", err)
	}
	defer updStmt.Close()

	for _, version := range migrationVersions {
		script := migrations[version]
		if version > fromVersion {

			fmt.Printf("%s Applying database migration v%d\n",
				color.HiBlueString("i"),
				version)

			_, err := m.db.Exec(script)
			if err != nil {
				log.Fatalf("Could not apply migration : %s, %v", script, err)
			}

			_, err = updStmt.Exec(version)
			if err != nil {
				log.Fatalf("Could not update version : %v", err)
			}
		}
	}
}
