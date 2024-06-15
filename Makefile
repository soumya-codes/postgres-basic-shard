# Define Docker Compose command
DOCKER_COMPOSE := docker-compose

# Docker Compose file
COMPOSE_FILE := ./deployment/docker-compose.yml

# Variables
BINARY_NAME=heartbeat_shard_app
MAIN_GO=main.go
PID_FILE=app.pid

# Targets
build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(MAIN_GO)

run: build
	@echo "Running the application..."
	@./$(BINARY_NAME) & echo $$! > $(PID_FILE)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

#Generate SQL
generate-sql:
	sqlc generate

# Build Docker containers
db-build:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

# Run Docker containers
db-up:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d

# Stop Docker containers
db-down:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

# Restart Docker containers
db-restart: down up

# View logs of Docker containers
db-logs:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f

# Remove Docker containers and associated volumes
db-clean:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down -v

# Sleep for 1 second
sleep:
	@sleep 1

# Create Setup
setup: db-up sleep run

# Bring down the setup
teardown: db-clean
	@echo "Stopping the application..."
	@if [ -f $(PID_FILE) ]; then \
		PID=$$(cat $(PID_FILE)); \
		kill $$PID && rm -f $(PID_FILE); \
		echo "Application stopped."; \
	else \
		echo "No PID file found. Is the application running?"; \
	fi
	@rm -f $(BINARY_NAME)

