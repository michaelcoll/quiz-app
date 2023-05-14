// Code generated by sqlc-addon. DO NOT EDIT.
// versions:
//   sqlc-addon v1.3.0

package db

import (
	"database/sql"
	"log"
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
CREATE TABLE quizz
(
    filename TEXT PRIMARY KEY
);

CREATE TABLE quizz_version
(
    sha1     TEXT PRIMARY KEY,
    filename TEXT,
    version  INTEGER NOT NULL,
    active   INTEGER DEFAULT 1,

    CONSTRAINT filename_fk FOREIGN KEY (filename) REFERENCES quizz (filename),
    CONSTRAINT quizz_version_unique UNIQUE (filename, version)
);

CREATE TABLE quizz_question
(
    sha1    TEXT PRIMARY KEY,
    content TEXT NOT NULL
);

CREATE TABLE quizz_question_version
(
    version_sha1  TEXT,
    question_sha1 TEXT,

    CONSTRAINT pk PRIMARY KEY (version_sha1, question_sha1),
    CONSTRAINT version_fk FOREIGN KEY (version_sha1) REFERENCES quizz_version (sha1),
    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quizz_question (sha1)
);

CREATE TABLE quizz_answer
(
    sha1    TEXT PRIMARY KEY,
    valid   INTEGER,
    content TEXT NOT NULL
);

CREATE TABLE quizz_question_answer
(
    question_sha1 TEXT,
    answer_sha1 TEXT,

    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quizz_question (sha1),
    CONSTRAINT answer_fk FOREIGN KEY (answer_sha1) REFERENCES quizz_answer (sha1)
);
`

var migrations = map[int]string{
	1: v1Init,
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
		m.applyMigration(version)
	} else {
		m.createDBVersionTable()
		m.applyMigration(0)
	}
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

	for version, script := range migrations {
		if version > fromVersion {
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