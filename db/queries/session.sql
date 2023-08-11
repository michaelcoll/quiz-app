-- name: CreateOrReplaceSession :exec
INSERT INTO session (uuid, quiz_sha1, user_id)
VALUES (?, ?, ?);

-- name: CreateOrReplaceSessionAnswer :exec
REPLACE INTO session_answer (session_uuid, user_id, question_sha1, answer_sha1, checked)
VALUES (?, ?, ?, ?, ?);

-- name: FindAllSessions :many
SELECT *
FROM session_view
WHERE quiz_active = ?
LIMIT ? OFFSET ?;

-- name: FindAllSessionsForUser :many
SELECT *
FROM session_view
WHERE quiz_active = ?
  AND user_id = ?
LIMIT ? OFFSET ?;

-- name: CountAllSessions :one
SELECT COUNT(*)
FROM session_view
WHERE quiz_active = ?;

-- name: CountAllSessionsForUser :one
SELECT COUNT(*)
FROM session_view
WHERE quiz_active = ?
  AND user_id = ?;

-- name: FindAllSessionsAnswerForSession :many
SELECT srv.quiz_sha1,
       srv.question_sha1,
       srv.answer_sha1,
       srv.session_uuid,
       srv.user_id,
       srv.checked,
       CASE
           WHEN sv.remaining_sec = 0
               THEN srv.result
           END
FROM session_response_view srv
         JOIN session_view sv ON srv.session_uuid = sv.uuid
WHERE session_uuid = ?
  AND srv.user_id = ?
