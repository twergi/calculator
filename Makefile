include ./deployments/.env

MIGRATION_FOLDER=$(CURDIR)/migrations
POSTGRES_SETUP=user=$(PGUSER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=$(PGPORT) sslmode=disable
COMPOSE_PATH=$(CURDIR)/deployments/docker-compose.yaml

.PHONY: run-dev
run-dev:
	go run $(CURDIR)/cmd/app/main.go

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-up-one
migration-up-one:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up-by-one

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: migration-reset
migration-reset:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" reset

.PHONY: environment-up-db
environment-up-db:
	docker compose -f $(COMPOSE_PATH) up calculator-db -d

.PHONY: environment-down
environment-down:
	docker compose -f $(COMPOSE_PATH) down




# Tests
.PHONY: test-environment-up
test-environment-up:
	docker compose -f ./tests/docker-compose.yaml up -d

.PHONY: test-environment-down
test-environment-down:
	docker compose -f ./tests/docker-compose.yaml down

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=postgres password=postgres dbname=calculator host=localhost port=5432 sslmode=disable" up

.PHONY: test-migration-reset
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=postgres password=postgres dbname=calculator host=localhost port=5432 sslmode=disable" reset





# Proto
.PHONY: generate
generate: generate-calc go-tidy

.PHONY: go-tidy
go-tidy:
	go mod tidy

.PHONY: generate-calc
generate-calc: generate-calc-grpc

.PHONY: generate-calc-grpc
generate-calc-grpc:
	protoc -I internal/proto \
		--go_out internal/proto/gen/go \
		--go_opt paths=source_relative \
		--go-grpc_out internal/proto/gen/go \
		--go-grpc_opt paths=source_relative \
		internal/proto/service/service.proto