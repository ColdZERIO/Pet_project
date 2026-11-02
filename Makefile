include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d postgres

env-down:
	docker compose down postgres

env-cleanup:
	@read -p "Clear all environment volume files? RISK OF DATA LOS. [Y/N]: " ans; \
	if [ "$$ans" = "Y" ]; then \
		docker compose down postgres && \
		rm -rf out/pgdata && \
		echo "env file cleared"; \
	else \
		echo "cleanup cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "invalid config seq. Example: make migrate-create seq=init"; \
  		exit 1; \
	fi; \
	docker compose run --rm postgres-migrate \
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
  		echo "invalid config action.Example: make migrate-action action=up"; \
  		exit 1; \
	fi; \
	docker compose run --rm postgres-migrate \
    	-path /migrations \
    	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable \
    	"$(action)"
