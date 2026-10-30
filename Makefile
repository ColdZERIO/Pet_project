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