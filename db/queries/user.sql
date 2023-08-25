-- name: FindActiveUserById :one
SELECT *
FROM user_class_view
WHERE id = ?
  AND active = 1;

-- name: FindUserById :one
SELECT *
FROM user_class_view
WHERE id = ?;

-- name: FindAllUser :many
SELECT *
FROM user_class_view;

-- name: CreateOrReplaceUser :exec
REPLACE INTO user (id, login, name, picture, role_id)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUserRole :exec
UPDATE user
SET role_id = ?
WHERE id = ?;

-- name: UpdateUserActive :exec
UPDATE user
SET active = ?
WHERE id = ?;

-- name: UpdateUserClass :exec
UPDATE user
SET class_uuid = ?
WHERE id = ?;

-- name: ClearClassUsers :exec
UPDATE user
SET class_uuid = NULL
WHERE class_uuid = ?;
