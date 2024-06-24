DOCKERCOMPOSE=docker-compose
DOCKER_COMPOSE_FILE_DEVELOPMENT=deployments/docker-compose.development.yaml
DOCKER_COMPOSE_FILE=deployments/docker-compose.yaml

ifeq (,$(wildcard ./app.env))
    $(error app.env file not found)
endif

include app.env

export $(shell sed 's/=.*//' .env)

run:
	@echo DB_HOST is $(DB_HOST)
	@echo DB_PORT is $(DB_PORT)

run-dev:
	cp env/.env.development app.env
	go run cmd/main.go

run-staging:
	cp env/.env.staging app.env
	go run cmd/main.go

up-dev:
	cp env/.env.development app.env
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE_DEVELOPMENT) up --build --force-recreate -d

down-dev:
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE_DEVELOPMENT) down

up-staging:
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build -d

down-staging:
	$(DOCKERCOMPOSE) -f $(DOCKER_COMPOSE_FILE) down 
