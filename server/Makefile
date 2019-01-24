.PHONY: run db

DB_DIR=container-data/postgres

# run server
run:
	go run main.go

# db runs a local Postgres server with Docker
db:
	mkdir -p "${DB_DIR}"
	docker run \
		-it \
		--rm \
		--net host \
		-e POSTGRES_USER=dev-push-server \
		-e POSTGRES_DB=dev-push-server \
		-v "${PWD}/${DB_DIR}:/var/lib/postgresql/data" \
		postgres:alpine
