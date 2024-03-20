-- name: CreateUser :exec
INSERT INTO users (
    first_name,
    last_name,
    email,
    password
)
VALUES (?,?,?,?);

-- name: GetUserID :one
SELECT user_id
FROM users
WHERE email = ?;


-- name: GetUserByEmail :one
SELECT *
FROM users 
WHERE email = ?;


-- name: PageAttractions :many
SELECT *
FROM attractions
WHERE attraction_id > ?
LIMIT 10;


-- name: PageFacilities :many
SELECT *
FROM facilities 
WHERE facility_id > ?
LIMIT 10;


-- name: GetLearingContent :many
SELECT *
FROM learning_content
WHERE name LIKE ?;


-- name: CreateAttraction :exec
INSERT INTO attractions (
    title,
    content
) VALUES (?,?);

-- name: CreateFacility :exec
INSERT INTO facilities (
    title,
    content
) VALUES (?,?);

-- name: CreateFactFile :exec
INSERT INTO learning_content (
    name,
    description,
    habitat,
    diet,
    conservation_status
) VALUES (?,?,?,?,?);

-- name: PageLearngingContent :many
SELECT * 
FROM learning_content
WHERE learning_content_id > ?
LIMIT 10;
