CREATE TABLE IF NOT EXISTS users (
    id       TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id     TEXT PRIMARY KEY,
    title  TEXT NOT NULL,
    author TEXT NOT NULL,
    path   TEXT NOT NULL,  -- e.g. /data/books/dune.epub
    UNIQUE(title, author)
);

CREATE TABLE IF NOT EXISTS bookmarks (
    id      TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    book_id TEXT NOT NULL,
    page    INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id),
    UNIQUE(user_id, book_id)
);