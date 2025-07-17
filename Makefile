DEV_COMPOSE_ARGS=--env-file .env.local -f Docker-compose.yml
DEV_COMPOSE_ENV=docker compose $(DEV_COMPOSE_ARGS)
DEV_COMPOSE=docker compose $(DEV_COMPOSE_ARGS)


dev-local-build:
	$(DEV_COMPOSE) build
dev-local-up: dev-local-build
	$(DEV_COMPOSE) --env-file .env.local up -d