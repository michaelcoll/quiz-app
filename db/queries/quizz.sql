-- name: FindLatestVersionByFilename :one
SELECT max(version)
FROM quizz
WHERE filename = ?;

-- name: CreateOrReplaceQuizz :exec
REPLACE INTO quizz (sha1, name, filename, version)
VALUES (?, ?, ?, ?);

-- name: CreateOrReplaceQuestion :exec
REPLACE INTO quizz_question (sha1, content)
VALUES (?, ?);

-- name: CreateOrReplaceAnswer :exec
REPLACE INTO quizz_answer (sha1, content, valid)
VALUES (?, ?, ?);

-- name: LinkQuestion :exec
REPLACE INTO quizz_question_quizz (quizz_sha1, question_sha1)
VALUES (?, ?);

-- name: LinkAnswer :exec
REPLACE INTO quizz_question_answer (question_sha1, answer_sha1)
VALUES (?, ?);

