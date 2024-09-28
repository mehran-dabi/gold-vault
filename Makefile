# Variables
COMPOSE_FILE = docker-compose.yml

# Default command to start the containers
up:
	@echo "Starting Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) up -d

# Bring up the services with fresh build
up-build:
	@echo "Building and starting Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) up -d --build

# Stop the containers
stop:
	@echo "Stopping Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) stop

# Stop and remove containers
down:
	@echo "Stopping and removing Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) down

# Rebuild containers without cache
rebuild:
	@echo "Rebuilding Docker Compose services without cache..."
	docker compose -f $(COMPOSE_FILE) build --no-cache

# Show logs for all services
logs:
	@echo "Showing logs for all Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) logs -f

# Show logs for a specific service
logs-service:
	@echo "Showing logs for service $(service)..."
	docker compose -f $(COMPOSE_FILE) logs -f $(service)

# Execute a shell inside a running service container
shell:
	@echo "Opening shell in service $(service)..."
	docker compose -f $(COMPOSE_FILE) exec $(service) sh

# Remove all stopped containers, volumes, and networks
clean:
	@echo "Cleaning up Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) down -v --rmi all --remove-orphans

# Show the status of all services
status:
	@echo "Showing status of Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) ps

# Restart all containers
restart:
	@echo "Restarting Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) down
	docker compose -f $(COMPOSE_FILE) up -d

# Build images for all services
build:
	@echo "Building Docker Compose services..."
	docker compose -f $(COMPOSE_FILE) build

# Generate Swagger docs for all microservices
swagger:
	@echo "Generating Swagger docs for user-service..."
	swag init -g ./user-service/cmd/app/main.go -o ./user-service/docs
	@echo "Generating Swagger docs for asset-service..."
	swag init -g ./asset-service/cmd/app/main.go -o ./asset-service/docs
	@echo "Generating Swagger docs for wallet-service..."
	swag init -g ./wallet-service/cmd/app/main.go -o ./wallet-service/docs
	@echo "Generating Swagger docs for trading-service..."
	swag init -g ./trading-service/cmd/app/main.go -o ./trading-service/docs
