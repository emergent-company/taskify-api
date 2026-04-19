-- name: CompleteTask :exec
UPDATE tasks SET completed = true WHERE id = ?;
