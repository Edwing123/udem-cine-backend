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

	selectUserIdPassword = `
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

	updateMovie = `
	UPDATE "movie"
	SET title = $2, classification = $3, genre = $4, duration = $5, release_date = $6
	WHERE id = $1;
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

	// Schedule queries.
	selectSchedule = `
	SELECT id, name, time
	FROM "schedule"
	WHERE id = $1;
	`

	selectAllSchedules = `
	SELECT id, name, time
	FROM "schedule";
	`

	insertSchedule = `
	INSERT INTO "schedule" (name, time)
	VALUES ($1, $2);
	`

	updateSchedule = `
	UPDATE "schedule"
	SET name = $2, time = $3
	WHERE id = $1;
	`

	deleteSchedule = `
	DELETE FROM "schedule"
	WHERE id = $1;
	`

	// Function queries.
	selectFunction = `
	SELECT id, price, created_at, movie_id, room, schedule_id
	FROM "function"
	WHERE id = $1; 
	`

	selectFunctionDetails = `
	SELECT f.id, f.price, f.created_at, m.title, room, s.name

	FROM "function" AS f
	INNER JOIN "movie" AS m
	ON f.movie_id = m.id

	INNER JOIN "schedule" as s
	ON f.schedule_id = s.id
	`

	selectAllFunctions = `
	SELECT id, price, created_at, movie_id, room, schedule_id
	FROM "function";
	`

	insertFunction = `
	INSERT INTO "function" (price, movie_id, room, schedule_id)
	VALUES($1 , $2, $3, $4);
	`

	updateFunction = `
	UPDATE "function"
	SET price = $2,
		movie_id = $3,
		room = $4,
		schedule_id = $5
	WHERE id = $1;
	`

	deleteFunction = `
	UPDATE "function"
	SET room = NULL
	WHERE id = $1;
	`
)
