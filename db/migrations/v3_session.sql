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
    user_id       TEXT    NOT NULL,
    question_sha1 TEXT    NOT NULL,
    answer_sha1   TEXT    NOT NULL,
    checked       INTEGER NOT NULL,

    PRIMARY KEY (session_uuid, question_sha1, answer_sha1),
    FOREIGN KEY (session_uuid) REFERENCES session (uuid),
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1),
    FOREIGN KEY (answer_sha1) REFERENCES quiz_answer (sha1)
);

CREATE VIEW quiz_answer_count_view
AS
SELECT q.sha1          AS quiz_sha1,
       COUNT(qa.valid) AS checked_answers
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question_answer qqa ON qqq.question_sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
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
         LEFT JOIN session_answer sa ON qa.sha1 = sa.answer_sha1 AND sa.question_sha1 = qqq.question_sha1 AND
                                        sa.answer_sha1 = qqa.answer_sha1;

CREATE VIEW session_view
AS
SELECT s.uuid                                                                                       AS uuid,
       q.sha1                                                                                       AS quiz_sha1,
       q.name                                                                                       AS quiz_name,
       q.active                                                                                     AS quiz_active,
       u.id                                                                                         AS user_id,
       CAST(u.firstname || ' ' || u.lastname AS TEXT)                                               AS user_name,
       CAST(MAX(q.duration - (STRFTIME('%s', 'now') - STRFTIME('%s', s.created_at)), 0) AS INTEGER) AS remaining_sec,
       checked_answers,
       SUM(srv.result)                                                                              AS results
FROM session s
         JOIN quiz q ON q.sha1 = s.quiz_sha1
         JOIN user u ON u.id = s.user_id
         JOIN quiz_answer_count_view qacv ON s.quiz_sha1 = qacv.quiz_sha1
         JOIN session_response_view srv ON s.uuid = srv.session_uuid;

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