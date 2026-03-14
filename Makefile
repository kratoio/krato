include .env

DATABASE_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

migrate-up:
	migrate -path migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path migrations -database $(DATABASE_URL) down

migrate-down-one:
	migrate -path migrations -database $(DATABASE_URL) down 1

run-server:
	go run cmd/main.go


air:
	air
