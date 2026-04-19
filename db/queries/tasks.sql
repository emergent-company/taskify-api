-- name: CompleteTask :exec
UPDATE tasks SET completed = TRUE WHERE id = $1;
