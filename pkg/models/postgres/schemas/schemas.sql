CREATE TABLE IF NOT EXISTS 'movie' (
    id SERIAL PRIMARY KEY,
    title TEXT,
    clasification TEXT,
    genre TEXT,
    duration TIME,
    release_date TIME
);

CREATE TABLE IF NOT EXISTS 'user' (
    id SERIAL PRIMARY KEY,
    name TEXT,
    role TEXT,
    password CHAR(60)
);