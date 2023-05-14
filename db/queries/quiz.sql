-- name: FindLatestVersionByFilename :one
SELECT max(version)
FROM quiz
WHERE filename = ?;

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

