-- name: GetUser :one
SELECT * FROM user
WHERE ID = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM user
ORDER BY name;

-- name: CreateUser :execresult
INSERT INTO user (
  Firstname, LastName, Biography
) VALUES (
  ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM user
WHERE ID = ?;

-- update
UPDATE user 
SET Firstname = ?, LastName = ?, Biography = ?
WHERE ID = ?