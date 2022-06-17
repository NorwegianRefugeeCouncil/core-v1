FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN mkdir /out && go build -o /out/app --tags fts5 && chmod +x /out/app && rm -rf /app
WORKDIR /out
ENTRYPOINT ["/out/app"]