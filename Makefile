
COMPOSE := $(shell ./scripts/find-compose.sh)

.PHONY:
prerequisites:
	@./scripts/prerequisites.sh

.PHONY:
image: prerequisites
	@echo "Building image..."
	DOCKER_BUILDKIT=1 docker build -f ./build/package/app.Dockerfile -t core .

.PHONY: up
up: prerequisites
	@$(COMPOSE) -f deploy/docker-compose.yaml up

.PHONY: down
down: prerequisites
	@$(COMPOSE) -f deploy/docker-compose.yaml down

.PHONY: test
test:
	@go test ./...

.PHONY: proxy
proxy: prerequisites
	@envoy -c deploy/envoy.yaml

.PHONY: serve
serve: prerequisites
	@go run . serve --listen-address=:8080 --db-driver=postgres --db-dsn=postgres://postgres:postgres@localhost:5432/core?sslmode=disable --log-level=debug