CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(32) PRIMARY KEY,
    -- Enough to store a SHA-512 hash in hex
    password VARCHAR(64),
    salt VARCHAR(64)
);