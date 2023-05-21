-- name: FindUserById :one
SELECT *
FROM user
WHERE id = ?;

-- name: FindAllUser :many
SELECT *
FROM user;

-- name: CreateOrReplaceUser :exec
REPLACE INTO user (id, email, firstname, lastname, role_id)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUserRole :exec
UPDATE user
SET role_id = ?
WHERE id = ?;

-- name: UpdateUserActive :exec
UPDATE user
SET active = ?
WHERE id = ?;
