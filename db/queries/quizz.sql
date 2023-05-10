-- name: GetQuizz :one
SELECT filename
FROM quizz
WHERE filename = ?;
