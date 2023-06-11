-- name: CreateOrReplaceClass :exec
REPLACE INTO student_class (uuid, name)
VALUES (?, ?);

-- name: FindAllClasses :many
SELECT *
FROM student_class
LIMIT ? OFFSET ?;

-- name: DeleteClassById :exec
DELETE
FROM student_class
WHERE uuid = ?;

-- name: AssignUserToClass :exec
UPDATE user
SET class_uuid = ?
WHERE id = ?;

-- name: CreateQuizClassVisibility :exec
REPLACE INTO quiz_class_visibility (class_uuid, quiz_sha1)
VALUES (?, ?);

-- name: DeleteQuizClassVisibility :exec
DELETE
FROM quiz_class_visibility
WHERE class_uuid = ?
  AND quiz_sha1 = ?;
