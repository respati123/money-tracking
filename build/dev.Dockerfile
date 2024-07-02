FROM golang:1.22-alpine

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 

# CMD ["ls" ,"-la"]

COPY . .

RUN echo "ls -la"

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o main cmd/money/main.go" -command="./main"