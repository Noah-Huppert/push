CREATE TABLE api_tokens (
	ID SERIAL,
	user_id INTEGER REFERENCES users (id) NOT NULL,
	device_id INTEGER REFERENCES devices (id) NOT NULL,
	token_hash TEXT NOT NULL
)
