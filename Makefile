DOCKERCOMPOSE=docker compose
DOCKER_COMPOSE_FILE_DEVELOPMENT=deployments/docker-compose.development.yaml
DOCKER_COMPOSE_FILE=deployments/docker-compose.yaml

ifeq (,$(wildcard ./.env))
    $(error .env file not found)
endif

include .env

export $(shell sed 's/=.*//' .env)

swagger-generate:
	swag init -g cmd/money/main.go

run:
	@echo DB_HOST is $(DB_HOST)
	@echo DB_PORT is $(DB_PORT)

run-local:
	cp env/.env.local .env
	go run cmd/money/main.go

run-dev:
	cp env/.env.development .env
	go run cmd/money/main.go

run-staging:
	cp env/.env.staging .env
	go run cmd/money/main.go

up-dev:
	cp env/.env.development .env
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE_DEVELOPMENT) up --build --force-recreate -d

down-dev:
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE_DEVELOPMENT) down

up-staging:
	cp env/.env.development .env
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build -d

down-staging:
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE) down 

migrations-up: 
	migrate -database postgresql://$(DB_USER):$(DB_PASS)@localhost:5434/$(DB_NAME)?sslmode=disable -path db/migrations up

migrations-down: 
	migrate -database postgresql://$(DB_USER):$(DB_PASS)@localhost:5434/$(DB_NAME)?sslmode=disable -path db/migrations down

migrations-force: 
	migrate -database postgresql://$(DB_USER):$(DB_PASS)@localhost:5434/$(DB_NAME)?sslmode=disable -path db/migrations force
