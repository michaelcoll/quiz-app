-- name: FindBySha1 :one
SELECT *
FROM quiz
WHERE sha1 = ?;

-- name: FindLatestVersionByFilename :one
SELECT *
FROM quiz
WHERE filename = ?
ORDER BY version DESC
LIMIT 1;

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
AND version <> ?