CREATE TABLE devices (
	id SERIAL,
	user_id INTEGER REFERENCES users (id) NOT NULL,
	platform PLATFORM NOT NULL,
	name TEXT NOT NULL
)
