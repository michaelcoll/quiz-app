-- name: FindUserById :many
SELECT u.*, ur.role_id
FROM user u LEFT JOIN user_role ur ON u.id = ur.user_id
WHERE id = ?;

-- name: CreateOrReplaceUser :exec
REPLACE INTO user (id, email, firstname, lastname)
VALUES (?, ?, ?, ?);

-- name: FindTokenByTokenStr :one
SELECT t.*, u.email
FROM token t JOIN user u ON u.id = t.user_id
WHERE opaque_token = ?;

-- name: CreateOrReplaceToken :exec
REPLACE INTO token (opaque_token, user_id, expires, aud)
VALUES (?, ?, ?, ?);

-- name: AddRoleToUser :exec
REPLACE INTO user_role (user_id, role_id)
VALUES (?, ?);

-- name: RemoveAllRoleFromUser :exec
DELETE FROM user_role
WHERE user_id = ?

