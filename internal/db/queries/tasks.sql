-- name: InsertTask :one
INSERT INTO tasks (id, title, description, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
