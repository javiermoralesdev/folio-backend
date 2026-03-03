-- ============
-- USERS
-- ============
CREATE TABLE IF NOT EXISTS users (
    id       TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- ============
-- BOOKS
-- ============
CREATE TABLE IF NOT EXISTS books (
    id     TEXT PRIMARY KEY,
    title  TEXT NOT NULL,
    author TEXT NOT NULL,
    path   TEXT NOT NULL,
    UNIQUE(title, author)
);

-- ============
-- BOOKMARKS
-- ============
CREATE TABLE IF NOT EXISTS bookmarks (
    id      TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id),
    book_id TEXT NOT NULL REFERENCES books(id),
    page    INTEGER NOT NULL,
    UNIQUE(user_id, book_id)
);