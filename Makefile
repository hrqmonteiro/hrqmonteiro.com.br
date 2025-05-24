.PHONY: dev clean css css-watch templ-generate build setup

# Development server with live reload
dev: css
	@echo "Starting development server..."
	air

# Production build
build: clean
	@echo "Building for production..."
	@echo "Generating templates..."
	templ generate
	@echo "Building CSS..."
	npm run css:build
	@echo "Building Go binary..."
	@mkdir -p ./bin
	CGO_ENABLED=0 go build -v -o ./bin/server -ldflags="-s -w" ./cmd/server/main.go
	@echo "Build complete! Binary location: ./bin/server"

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	@find ./templates -name "*_templ.go" -type f -delete
	@rm -f ./static/css/main.css
	@rm -rf ./bin
	@rm -rf ./cmd/server/tmp

# Generate CSS
css:
	@echo "Generating CSS..."
	npm run css

# Watch CSS changes
css-watch:
	@echo "Watching CSS changes..."
	npm run css:watch

# Generate templ files
templ-generate:
	@echo "Generating templ files..."
	templ generate

# Install dependencies
.PHONY: setup

setup:
	@echo "Installing dependencies..."
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	npm install
	templ generate
	go mod tidy

# Docker commands
docker-build:
	docker compose build --no-cache

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f