APP_ENV ?= development

ifneq ($(APP_ENV), production)
include env/.env.development
else
include env/.env.production
endif

export

run:
	go run ./cmd/api

psql:
	psql ${GREENLIGHT_DB_DSN}
up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up