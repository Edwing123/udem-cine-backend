package postgres

const (
	// User queries.
	selectUser = `
	SELECT id, name, role, password
	FROM "user"
	WHERE id = $1;
	`

	selectAllUsers = `
	SELECT id, name, role, password
	FROM "user";
	`

	insertUser = `
	INSERT INTO "user" (name, role, password)
	VALUES ($1, $2, $3);
	`

	updateUser = `
	UPDATE "user"
	SET name = $2, role = $3
	WHERE id = $1;
	`

	deleteUser = `
	DELETE FROM "user"
	WHERE id = $1;
	`

	selectIdPassword = `
	SELECT id, password
	FROM "user"
	WHERE name = $1;
	`

	// Movie queries.
	selectMovie = `
	SELECT id, title, classification, genre, duration, release_date
	FROM "movie"
	WHERE id = $1;
	`

	selectAllMovies = `
	SELECT id, title, classification, genre, duration, release_date
	FROM "movie";
	`

	insertMovie = `
	INSERT INTO "movie" (title, classification, genre, duration, release_date)
	VALUES($1, $2, $3, $4, $5);
	`

	deleteMovie = `
	DELETE FROM "movie"
	WHERE id = $1;
	`

	// Room queries.
	selectRoom = `
	SELECT r.number, COUNT(1)
	FROM "room" AS r
		INNER JOIN "seat" AS s
		ON r.number = s.room

	WHERE r.number = $1
	GROUP BY r.number;
	`

	selectAllRooms = `
	SELECT r.number, COUNT(1)
	FROM "room" AS r
		INNER JOIN "seat" AS s
		ON r.number = s.room

	GROUP BY r.number;
	`

	selectAllSeats = `
	SELECT number, room FROM "seat"
	WHERE room = $1;
	`

	insertRoom = `
	INSERT INTO "room" (number) values($1);
	`

	insertSeat = `
	INSERT INTO "seat" (number, room) values($1, $2);
	`

	updateRoom = `
	UPDATE "room"
	SET number = $2
	WHERE number = $1;
	`

	deleteRoom = `
	DELETE FROM "room"
	WHERE number = $1;
	`
)
