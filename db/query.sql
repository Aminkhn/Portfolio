/*--------------------------------*/
/**-------- User Query ----------**/
/*________________________________*/
-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;
-- name: ListUser :many
SELECT * FROM user
ORDER BY name;
-- name: CreateUser :execresult
INSERT INTO user (
  username, email, name, family, password, role, image, phone_number, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
-- name: updateUser :one
UPDATE authors
set 
    username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
    image = ?, phone_number = ?, created_at =?
WHERE id = ?
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
/*--------------------------------*/
/**-------- Post Query ----------**/
/*________________________________*/
-- name: GetPost :one
SELECT * FROM post
WHERE id = ? LIMIT 1;
-- name: ListPost :many
SELECT * FROM post
ORDER BY name;
-- name: CreatePost :execresult
INSERT INTO post (
  username, email, name, family, password, role, image, phone_number, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
-- name: updatePost :one
UPDATE authors
set 
    username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
    image = ?, phone_number = ?, created_at =?
WHERE id = ?
RETURNING *;
-- name: DeletePost :exec
DELETE FROM user
WHERE id = ?;
/*--------------------------------*/
/**-------- User Query ----------**/
/*________________________________*/
-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;
-- name: ListUser :many
SELECT * FROM user
ORDER BY name;
-- name: CreateUser :execresult
INSERT INTO user (
  username, email, name, family, password, role, image, phone_number, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
-- name: updateUser :one
UPDATE authors
set 
    username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
    image = ?, phone_number = ?, created_at =?
WHERE id = ?
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
/*--------------------------------*/
/**-------- User Query ----------**/
/*________________________________*/
-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;
-- name: ListUser :many
SELECT * FROM user
ORDER BY name;
-- name: CreateUser :execresult
INSERT INTO user (
  username, email, name, family, password, role, image, phone_number, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
-- name: updateUser :one
UPDATE authors
set 
    username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
    image = ?, phone_number = ?, created_at =?
WHERE id = ?
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
/*--------------------------------*/
/**-------- User Query ----------**/
/*________________________________*/
-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;
-- name: ListUser :many
SELECT * FROM user
ORDER BY name;
-- name: CreateUser :execresult
INSERT INTO user (
  username, email, name, family, password, role, image, phone_number, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
-- name: updateUser :one
UPDATE authors
set 
    username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
    image = ?, phone_number = ?, created_at =?
WHERE id = ?
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;