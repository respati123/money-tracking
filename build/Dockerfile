# FROM golang:1.20-alpine


# ENV PROJECT_DIR=/app \
#     GO111MODULE=on \
#     CGO_ENABLED=0

# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download 

# COPY .. ./


# RUN go build -o main cmd/money/main.go
# EXPOSE 8080

# CMD [ "../main" ]

FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .. .

RUN go build -o main ./cmd/money/main.go

COPY ../.env .

RUN echo "Running ls -la in /app" && ls -la

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN echo ls -la


EXPOSE 8080

ENTRYPOINT [ "/app/main" ]