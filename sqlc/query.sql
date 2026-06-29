-- name: GetUser :one
SELECT *
FROM "user"
WHERE ID = $1 LIMIT 1;

-- name: ListUser :many
SELECT *
FROM "user"
ORDER BY Firstname;

-- name: CreateUser :exec
INSERT INTO "user" (Firstname, LastName, Biography)
VALUES ($1, $2, $3);

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE ID = $1;

-- name: UpdateUser :exec
UPDATE "user"
SET Firstname = $1,
    LastName  = $2,
    Biography = $3
WHERE ID = $4;

-- name: SearchByName :one
SELECT *
FROM "user"
WHERE Firstname = $1;