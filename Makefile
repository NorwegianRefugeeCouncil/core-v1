
COMPOSE := $(shell ./scripts/find-compose.sh)

.PHONY:
prerequisites:
	@./scripts/prerequisites.sh

.PHONY:
image: prerequisites
	@echo "Building image..."
	DOCKER_BUILDKIT=1 docker build -f ./build/package/app.Dockerfile -t core .

.PHONY: up
## Launches the docker-compose dependency stack
up: prerequisites
	@$(COMPOSE) -f deploy/docker-compose.yaml up

.PHONY: down
## Stops the docker-compose dependency stack
down: prerequisites
	@$(COMPOSE) -f deploy/docker-compose.yaml down

.PHONY: .dev-image
## Builds the development image
.dev-image:
	@DOCKER_BUILDKIT=1 docker build -f ./build/package/app.Dockerfile -t core-dev --target=dev --build-arg uid=$(shell id -u) --build-arg gid=$(shell id -g) .

.PHONY: coverage
## Runs the tests and generates a coverage report
coverage: .dev-image
	@mkdir -p ./reports
	docker run --rm -it -v $(shell pwd)/reports:/app/reports core-dev coverage

.PHONY: test
## Runs the tests
test: .dev-image
	docker run --rm -it -v $(shell pwd)/reports:/app/reports core-dev test

.PHONY: proxy
## Runs the envoy proxy
proxy: prerequisites
	@envoy -c deploy/envoy.yaml -l debug

.PHONY: serve
# Starts the server
serve: prerequisites
	@go run . serve \
		--listen-address=:8080 \
		--db-driver=postgres \
		--db-dsn=postgres://postgres:postgres@localhost:5432/core?sslmode=disable \
		--log-level=debug \
		--jwt-global-admin-group="NRC_Core_GlobalAdmin" \
		--auth-header-name="Authorization" \
		--auth-header-format="bearer-token" \
		--oidc-issuer="https://localhost:10000" \
		--oauth-client-id="foo" \
		--logout-url="https://localhost:10000/oauth2/sign_out?rd=https%3A%2F%2Flocalhost:10000%2Fsession%2Fend"

.PHONY: bootstrap
bootstrap:
	@cd web/theme && yarn && yarn build


.PHONY: generate
# Generates source code
generate: prerequisites
	@go generate ./...
	@go fix ./...

.DEFAULT_GOAL := show-help
# See <https://gist.github.com/klmr/575726c7e05d8780505a> for explanation.
.PHONY: show-help
show-help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)";echo;sed -ne"/^## /{h;s/.*//;:d" -e"H;n;s/^## //;td" -e"s/:.*//;G;s/\\n## /---/;s/\\n/ /g;p;}" ${MAKEFILE_LIST}|LC_ALL='C' sort -f|awk -F --- -v n=$$(tput cols) -v i=19 -v a="$$(tput setaf 6)" -v z="$$(tput sgr0)" '{printf"%s%*s%s ",a,-i,$$1,z;m=split($$2,w," ");l=n-i;for(j=1;j<=m;j++){l-=length(w[j])+1;if(l<= 0){l=n-i-length(w[j])-1;printf"\n%*s ",-i," ";}printf"%s ",w[j];}printf"\n";}'|more $(shell test $(shell uname) == Darwin && echo '-Xr')

