CREATE TABLE quizz
(
    sha1       TEXT PRIMARY KEY,
    name       TEXT      NOT NULL,
    filename   TEXT      NOT NULL,
    version    INTEGER   NOT NULL DEFAULT 1,
    active     INTEGER   NOT NULL DEFAULT 1,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT filename_fk FOREIGN KEY (filename) REFERENCES quizz (filename),
    CONSTRAINT quizz_version_unique UNIQUE (filename, version)
);

CREATE TABLE quizz_question
(
    sha1    TEXT PRIMARY KEY,
    content TEXT NOT NULL
);

CREATE TABLE quizz_question_quizz
(
    quizz_sha1    TEXT NOT NULL,
    question_sha1 TEXT NOT NULL,

    CONSTRAINT pk PRIMARY KEY (quizz_sha1, question_sha1),
    CONSTRAINT quizz_fk FOREIGN KEY (quizz_sha1) REFERENCES quizz (sha1),
    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quizz_question (sha1)
);

CREATE TABLE quizz_answer
(
    sha1    TEXT PRIMARY KEY,
    valid   INTEGER NOT NULL,
    content TEXT    NOT NULL
);

CREATE TABLE quizz_question_answer
(
    question_sha1 TEXT NOT NULL,
    answer_sha1   TEXT NOT NULL,

    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quizz_question (sha1),
    CONSTRAINT answer_fk FOREIGN KEY (answer_sha1) REFERENCES quizz_answer (sha1)
);