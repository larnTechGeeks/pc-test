include .env

up:
	docker compose down -v
	docker compose build
	docker compose up -d --remove-orphans

down:
	docker compose down -v --remove-orphans

api:
	gow run cmd/api/*.go

migration:
	goose -dir internal/db/migrations create $(name) sql

rollback:
	goose -dir 'internal/db/migrations' postgres ${DATABASE_URL} down

migrate:
	goose -dir 'internal/db/migrations' postgres ${DATABASE_URL} up
