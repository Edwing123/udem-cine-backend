CREATE TYPE "user_role" AS ENUM('admin', 'taquillero');
CREATE TYPE "movie_classification" AS ENUM('G', 'PG', 'PG-13', 'R', 'NC-17');

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
    duration SMALLINT,
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

CREATE TABLE IF NOT EXISTS "schedule" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    time TIME NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "function" (
    id SERIAL PRIMARY KEY,
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    movie_id INT NOT NULL REFERENCES "movie"(id),
    room SMALLINT NULL REFERENCES "room"(number),
    schedule_id INT NOT NULL REFERENCES "schedule"(id),
    CONSTRAINT function_room_schedule_unique UNIQUE (room, schedule_id)
);