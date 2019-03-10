CREATE TABLE pushes (
	ID SERIAL,
	user_id INTEGER REFERENCES users (id) NOT NULL,
	from_device_id INTEGER REFERENCES devices (id) NOT NULL,
	data TEXT NOT NULL
)
