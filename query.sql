-- name: GetUser :one
SELECT user_id, username, email, created_at, updated_at FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserWithPassword :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT user_id, username, email, created_at, updated_at FROM users;

-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING user_id, username, email, created_at, updated_at;

-- name: UpdateUser :one
UPDATE users
    set username = $2,
    email = $3,
    password = $4
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
