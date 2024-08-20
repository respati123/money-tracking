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

COPY .. ./

RUN go build -o main ./cmd/money/main.go

COPY . .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/env .

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]