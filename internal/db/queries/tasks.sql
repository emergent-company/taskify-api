-- name: CompleteTask :exec
UPDATE tasks
SET completed = TRUE,
    updated_at = NOW()
WHERE id = $1;
