package postgres

const (
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
)
