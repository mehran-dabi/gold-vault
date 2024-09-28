# Makefile

# Define the Docker Compose file
DOCKER_COMPOSE_FILE=docker-compose.yaml

# Paths to binaries
MIGRATE = migrate

# Variables
DB_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable
MIGRATIONS_DIR = ./internal/infrastructure/db/migrations

# Run Docker Compose to bring up the containers
up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

# Run Docker Compose to bring down the containers
down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

# Run swag to generate Swagger documentation
swag:
	swag init -g ./cmd/app/main.go -o ./docs

# Command to create a new migration file (both up and down)
.PHONY: create-migration
create-migration:
	@if [ -z "$(NAME)" ]; then echo "Error: Please specify a migration NAME."; exit 1; fi
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

# Command to force migration version (reset or set version)
.PHONY: set-migration-version
set-migration-version:
	@if [ -z "$(VERSION)" ]; then echo "Error: Please specify a VERSION."; exit 1; fi
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(VERSION)

# Command to run migrations
.PHONY: migrate-up
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Command to rollback migrations
.PHONY: migrate-down
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Example environment variables
export POSTGRES_USER = pgsql
export POSTGRES_PASSWORD = 123456
export POSTGRES_DB = user-service