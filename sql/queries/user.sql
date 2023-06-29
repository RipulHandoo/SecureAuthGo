-- name: CreateUser :one
INSERT INTO jwtUSER (Email, Password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM jwtUSER WHERE Email = $1;
