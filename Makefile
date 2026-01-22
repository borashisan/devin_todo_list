.PHONY: up down build logs clean restart backend-logs frontend-logs db-logs migrate migrate-down migrate-status

up:
	docker compose up -d

up-build:
	docker compose up -d --build

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

backend-logs:
	docker compose logs -f backend

frontend-logs:
	docker compose logs -f frontend

db-logs:
	docker compose logs -f db

clean:
	docker compose down -v --rmi local

restart:
	docker compose restart

ps:
	docker compose ps

migrate:
	docker compose run --rm migrate

migrate-down:
	docker compose run --rm migrate -dir /app/migrations mysql "${DB_USER}:${DB_PASSWORD}@tcp(db:3306)/${DB_NAME}?parseTime=true" down

migrate-status:
	docker compose run --rm migrate -dir /app/migrations mysql "${DB_USER}:${DB_PASSWORD}@tcp(db:3306)/${DB_NAME}?parseTime=true" status
