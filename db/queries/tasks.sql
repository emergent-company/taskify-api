-- name: ListTasks :many
SELECT 
    id, 
    title, 
    description, 
    completed, 
    created_at, 
    updated_at
FROM tasks;
