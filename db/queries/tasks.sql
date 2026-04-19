-- name: CompleteTask :one
UPDATE tasks
SET completed = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
