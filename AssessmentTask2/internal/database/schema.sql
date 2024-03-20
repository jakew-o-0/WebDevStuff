PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS users_email_idx ON users(email);


CREATE TABLE IF NOT EXISTS tickets (
    ticket_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,

    ticket_uuid TEXT NOT NULL,
    secret TEXT NOT NULL,
    expiry REAL NOT NULL,

    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE INDEX IF NOT EXISTS tickets_uuid_idx ON tickets(ticket_uuid);


CREATE TABLE IF NOT EXISTS sessions (
    token TEXT PRIMARY KEY,
    data BLOB NOT NULL,
    expiry REAL NOT NULL
);
CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions(expiry);


CREATE TABLE IF NOT EXISTS attractions (
    attraction_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS facilities (
    facility_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS learning_content (
    learning_content_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,

    name TEXT NOT NULL,
    description TEXT NOT NULL,
    habitat TEXT NOT NULL,
    diet TEXT NOT NULL,
    conservation_status TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS learning_content_title_idx ON learning_content(name);
