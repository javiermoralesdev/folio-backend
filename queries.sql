-- name: CreateUser :one
INSERT INTO users (id, username, password)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;