CREATE TYPE "user_role" AS ENUM('admin', 'taquillero');

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    role "user_role" NOT NULL,
    password CHAR(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS "movie" (
    id SERIAL PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    classification TEXT NOT NULL,
    genre TEXT NOT NULL,
    duration INTERVAL HOUR TO MINUTE NOT NULL,
    release_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS "room" (
    number SMALLINT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "seat" (
    number SMALLINT,
    room SMALLINT NOT NULL REFERENCES "room"(number) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT seat_primary_key PRIMARY KEY (number, room)
);