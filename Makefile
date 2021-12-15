MIGRATIONS_DIR := "./migrations"
PG_DSN := "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable"

.PHONY: migrate-generate
migrate-generate:
	$(GOPATH)/bin/goose -dir $(MIGRATIONS_DIR) create $(name) sql

.PHONY: migrate-up
migrate-up:
	$(GOPATH)/bin/goose -dir $(MIGRATIONS_DIR) postgres "$(PG_DSN)" up
.PHONY: migrate-down
migrate-down:
	$(GOPATH)/bin/goose -dir $(MIGRATIONS_DIR) postgres "$(PG_DSN)" down