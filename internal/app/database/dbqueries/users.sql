-- name: GetUser :one
SELECT * FROM users 
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users 
ORDER BY first_name;

-- name: CreateUserWithEmailAndPassword :execresult
INSERT INTO users (
  first_name, last_name, email, password, registration_method_id
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;