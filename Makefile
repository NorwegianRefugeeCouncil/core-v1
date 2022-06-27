
.PHONY:
image:
	@echo "Building image..."
	DOCKER_BUILDKIT=1 docker build -f ./build/package/app.Dockerfile -t core .

.PHONY: up
up:
	@docker-compose -f deploy/docker-compose.yaml up

.PHONY: down
down:
	@docker-compose -f deploy/docker-compose.yaml down

.PHONY: proxy
proxy:
	@envoy -c deploy/envoy.yaml