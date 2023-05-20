-- name: CreateOrReplaceQuiz :exec
REPLACE INTO quiz (sha1, name, filename, version)
VALUES (?, ?, ?, ?);

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

-- name: FindQuizBySha1 :one
SELECT *
FROM quiz
WHERE sha1 = ?;

-- name: FindQuizFullBySha1 :many
SELECT
    q.sha1 as quiz_sha1,
    q.filename as quiz_filename,
    q.name as quiz_name,
    q.version as quiz_version,
    q.created_at as quiz_created_at,
    q.active as quiz_active,
    qq.sha1 as question_sha1,
    qq.content as question_content,
    qa.sha1 as answer_sha1,
    qa.content as answer_content,
    qa.valid as answer_valid
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question qq ON qq.sha1 = qqq.question_sha1
         JOIN quiz_question_answer qqa ON qq.sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
WHERE q.sha1 = ?;

-- name: FindQuizByFilenameAndLatestVersion :one
SELECT *
FROM quiz
WHERE filename = ?
ORDER BY version DESC
LIMIT 1;

-- name: FindAllActiveQuiz :many
SELECT *
FROM quiz
WHERE active = 1
LIMIT ? OFFSET ?;

-- name: CountAllActiveQuiz :one
SELECT count(1)
FROM quiz
WHERE active = 1;
