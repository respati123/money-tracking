FROM golang:1.20-alpine


ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 

COPY .. ./


RUN go build -o main cmd/money/main.go
EXPOSE 8080

CMD [ "../main" ]