-- name: ListTasks :many
SELECT id, title, description, created_at, updated_at
FROM tasks
ORDER BY created_at DESC;
