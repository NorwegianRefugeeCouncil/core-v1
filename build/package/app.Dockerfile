FROM golang:1.19 as base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp/.cache
RUN apt-get install -yq --no-install-recommends curl && \
    curl -fsSL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get install -yq --no-install-recommends nodejs && \
    npm install -g yarn
ARG uid=1000
ARG gid=1000
RUN groupadd -g ${GROUP_ID} core || echo "Group already exists"
RUN useradd -l -m -d /home/app -u ${USER_ID} -g ${GROUP_ID} core || echo "User already exists"
RUN mkdir /out && chown ${uid}:${gid} /out
USER ${uid}:${gid}
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download && \
    go install github.com/jstemmer/go-junit-report/v2@latest && \
    go install github.com/axw/gocov/gocov@latest && \
    go install github.com/AlekSi/gocov-xml@latest
COPY --chown=${uid}:${gid} . .

FROM base as dev
ENTRYPOINT ["/app/scripts/docker-dev-entrypoint.sh"]

FROM base as builder
RUN go build -ldflags '-w -s' -o /out/core --tags fts5 && chmod +x /out/core
WORKDIR /out
ENTRYPOINT ["/out/app"]

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /out/core /app/core
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/core"]