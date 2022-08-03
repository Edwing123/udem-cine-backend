CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    role TEXT CHECK (role IN ('admin', 'taquillero')),
    password CHAR(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS "movie" (
    id SERIAL PRIMARY KEY,
    title TEXT,
    clasification TEXT,
    genre TEXT,
    duration TIME,
    release_date TIME
);