-- name: GetPosts :many
SELECT title,content 
FROM posts
WHERE post_id > ?
LIMIT 10;

-- name: CreatePost :exec
INSERT INTO posts (
    title,
    content
) VALUES (?,?);


-- name: GetUserByUsername :one
SELECT user_id,username,password
FROM users
WHERE username = ?;

-- name: CreateUser :one
INSERT INTO users (
    username,
    password
) VALUES (?,?)
RETURNING user_id;
