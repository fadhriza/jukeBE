.PHONY: run-local run-staging run-prod migrate-up migrate-down

# Default to local
ENV ?= local

include .env.$(ENV)
export $(shell sed 's/=.*//' .env.$(ENV))

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

run:
	APP_ENV=$(ENV) go run cmd/api/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

# Examples of usage:
# make run ENV=local
# make run ENV=staging
# make migrate-up ENV=local
