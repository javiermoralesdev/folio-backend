-- ============
-- USERS
-- ============

-- name: CreateUser :one
INSERT INTO users (id, username, password)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

-- ============
-- BOOKS
-- ============

-- name: CreateBook :one
INSERT INTO books (id, title, author, path)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBookByID :one
SELECT * FROM books
WHERE id = ?;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = ?;

-- ============
-- BOOKMARKS
-- ============

-- name: UpsertBookmark :exec
INSERT INTO bookmarks (id, user_id, book_id, page)
VALUES (?, ?, ?, ?)
ON CONFLICT(user_id, book_id) DO UPDATE SET page = excluded.page;

-- name: GetBookmark :one
SELECT * FROM bookmarks
WHERE user_id = ? AND book_id = ?;

-- name: GetUserBookmarks :many
SELECT * FROM bookmarks
WHERE user_id = ?;

-- name: DeleteBookmark :exec
DELETE FROM bookmarks
WHERE user_id = ? AND book_id = ?;