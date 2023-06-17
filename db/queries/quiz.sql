-- name: CreateOrReplaceQuiz :exec
REPLACE INTO quiz (sha1, name, filename, version, duration, created_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateOrReplaceQuestion :exec
REPLACE INTO quiz_question (sha1, content)
VALUES (?, ?);

-- name: CreateOrReplaceAnswer :exec
REPLACE INTO quiz_answer (sha1, content, valid)
VALUES (?, ?, ?);

-- name: LinkQuestion :exec
REPLACE INTO quiz_question_quiz (quiz_sha1, question_sha1)
VALUES (?, ?);

-- name: LinkAnswer :exec
REPLACE INTO quiz_question_answer (question_sha1, answer_sha1)
VALUES (?, ?);

-- name: ActivateOnlyVersion :exec
UPDATE quiz
SET active = 0
WHERE filename = ?
  AND version <> ?;

-- name: FindQuizFullBySha1 :many
SELECT q.sha1       AS quiz_sha1,
       q.filename   AS quiz_filename,
       q.name       AS quiz_name,
       q.version    AS quiz_version,
       q.created_at AS quiz_created_at,
       q.duration   AS quiz_duration,
       q.active     AS quiz_active,
       qq.sha1      AS question_sha1,
       qq.content   AS question_content,
       qa.sha1      AS answer_sha1,
       qa.content   AS answer_content,
       qa.valid     AS answer_valid
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question qq ON qq.sha1 = qqq.question_sha1
         JOIN quiz_question_answer qqa ON qq.sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.sha1 = ?
    AND u.id = ''
   OR u.id = ?;

-- name: FindQuizByFilenameAndLatestVersion :one
SELECT *
FROM quiz
WHERE filename = ?
ORDER BY version DESC
LIMIT 1;

-- name: FindAllActiveQuiz :many
SELECT *
FROM quiz q
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.active = 1
    AND u.id = ''
   OR u.id = ?
LIMIT ? OFFSET ?;

-- name: CountAllActiveQuiz :one
SELECT COUNT(1)
FROM quiz q
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.active = 1
    AND u.id = ''
   OR u.id = ?;

-- name: FindAllQuizSessions :many
SELECT *
FROM quiz_session_view
LIMIT ? OFFSET ?;

-- name: FindAllQuizSessionsForUser :many
SELECT qsv.*
FROM quiz_session_view qsv
         JOIN quiz_class_visibility qcv ON qsv.quiz_sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid AND qsv.user_id = u.id
WHERE user_id = ?

LIMIT ? OFFSET ?;
