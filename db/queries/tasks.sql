-- name: CompleteTask :exec
UPDATE tasks
SET completed = true,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
