# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run
run:
	go run ./cmd/api

## db/psql: connect to the database using psql
.PHONY: psql
psql:
	psql ${GREENLIGHT_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migration/new
db/migration/new:
	@echo 'Creating migration files for ${name}...''
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migration/up: confirm
	@echo "Running up migrations..."
	migrate -path ./migrations -database ${DATABASE_URL} up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: 
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

# ==================================================================================== #
# BUILD
# ==================================================================================== #

current_time := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description := $(shell git describe --always --dirty --tags --long)
linker_flags := -s -X main.buildTime=${current_time} -X main.version=${git_description}
# Build the application and create a Linux binary
.PHONY: build
build:
	@echo 'Building the application...'
	go build -ldflags='$(linker_flags)' -o=./bin/api ./cmd/api
	@echo 'Building for Linux...'
	GOOS=linux GOARCH=amd64 go build -ldflags='$(linker_flags)' -o=./bin/api-linux_amd64 ./cmd/api
	@echo 'Build complete!'

# ==================================================================================== #
# DOCKER
# ==================================================================================== #

## docker/build: Build Docker image
.PHONY: docker/build
docker/build:
	docker build -t lightsaber-api .

## docker/up: Start services using Docker Compose
.PHONY: docker/up
docker/up:
	docker-compose up --build

## docker/down: Stop and remove containers
.PHONY: docker/down
docker/down:
	docker-compose down