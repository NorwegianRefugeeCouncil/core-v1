FROM golang:1.19 as base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp/.cache
RUN apt-get install -yq --no-install-recommends curl && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -yq --no-install-recommends nodejs && \
    npm install -g yarn
ARG uid=1000
ARG gid=1000
RUN groupadd -g ${gid} core || echo "Group already exists"
RUN useradd -l -m -d /home/app -u ${uid} -g ${gid} core || echo "User already exists"
RUN mkdir /out && chown ${uid}:${gid} /out
WORKDIR /app
ADD go.mod ./
ADD go.sum ./
USER ${uid}:${gid}
RUN go mod download -x -json && \
    go install github.com/jstemmer/go-junit-report/v2@latest && \
    go install github.com/axw/gocov/gocov@latest && \
    go install github.com/AlekSi/gocov-xml@latest
ADD --chown=${uid}:${gid} web web
WORKDIR web/theme
RUN yarn --ignore-scripts && yarn build && yarn rtlcss
WORKDIR /app
ADD --chown=${uid}:${gid} scripts scripts
ADD --chown=${uid}:${gid} cmd cmd
ADD --chown=${uid}:${gid} internal internal
ADD --chown=${uid}:${gid} pkg pkg
ADD --chown=${uid}:${gid} tools tools
ADD --chown=${uid}:${gid} main.go main.go
RUN go run . template
RUN go build -ldflags '-w -s' -o /out/core --tags fts5 && chmod +x /out/core

FROM base as dev
ENTRYPOINT ["/app/scripts/docker-dev-entrypoint.sh"]

FROM base as builder
WORKDIR /out
ENTRYPOINT ["/out/app"]

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /out/core /app/core
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/core"]
