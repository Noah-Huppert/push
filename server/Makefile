.PHONY: run db

DB_DIR=container-data/postgres
HTTP_PORT=8000
DB_NAME=dev-push-server
DB_USER=dev-push-server

# run server
#
# Arguments
#
#	ARGS    Additional argument to pass to server
run:
	PUSH_SERVER_HTTP_PORT="${HTTP_PORT}" \
	PUSH_SERVER_DB_NAME="${DB_NAME}" \
	PUSH_SERVER_DB_USER="${DB_USER}" \
	go run main.go ${ARGS}

# db runs a local Postgres server with Docker
db:
	mkdir -p "${DB_DIR}"
	docker run \
		-it \
		--rm \
		--net host \
		-e POSTGRES_DB="${DB_NAME}" \
		-e POSTGRES_USER="${DB_USER}" \
		-v "${PWD}/${DB_DIR}:/var/lib/postgresql/data" \
		postgres:alpine
