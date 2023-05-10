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