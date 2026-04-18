-- name: InsertTask :one
INSERT INTO tasks (
    id, title, description, status, created_by, org_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;
