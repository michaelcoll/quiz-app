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
    sha1    TEXT PRIMARY KEY,
    content TEXT NOT NULL
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