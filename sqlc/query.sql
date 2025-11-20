-- name: GetUser :one
SELECT * FROM user
WHERE ID = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM user
ORDER BY Firstname;

-- name: CreateUser :execresult
INSERT INTO user (
  Firstname, LastName, Biography
) VALUES (
  ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM user
WHERE ID = ?;

-- name: UpdateUser :exec
UPDATE user 
SET Firstname = ?, LastName = ?, Biography = ?
WHERE ID = ?;

-- name: SearchByName :one 
SELECT * FROM user WHERE Firstname = ?;