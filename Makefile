include .env
export

run: 
	@go run cmd/app/main.go

postgres-up:
	@docker compose up postgres -d

postgres-down:
	@docker compose down postgres

env-cleanup:
	@docker compose down -v

migrate-up:
	@migrate -path=./migrations -database ${DB_CONN} up

migrate-down:
	@migrate -path=./migrations -database ${DB_CONN} down

migrate-force:
	@migrate -path=./migrations -database ${DB_CONN} force 1