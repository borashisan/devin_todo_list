.PHONY: up down build logs clean restart backend-logs frontend-logs db-logs

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
