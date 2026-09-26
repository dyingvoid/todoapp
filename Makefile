include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todo-postgres

env-down:
	@docker compose down todo-postgres

env-cleanup:
	@read -p "Clean up volume? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todo-postgres && \
		rm -rf out/pgdata && \
		echo "Volume has been cleaned up"; \
	else \
		echo "Clean up cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq parameter is not specified. Example: make migrate-create seq=init"; \
		exit 1; \
	fi;

	@docker compose run --rm todo-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action is not specified. Example: make migrate-action action=up"; \
		exit 1; \
	fi;
	@docker compose run --rm todo-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

todoapp-run:
	@go run ${PROJECT_ROOT}/cmd/todoapp/