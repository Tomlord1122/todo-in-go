-- name: CreateTodo :one
INSERT INTO todos (
    "owner",
    "title",
    "category",
    "description",
    "completed"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos LIMIT $1 OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos SET
    "title" = $2,
    "category" = $3,
    "description" = $4,
    "completed" = $5
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
