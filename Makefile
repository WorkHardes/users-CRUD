DOCKERFILE_PATH=.

DOCKER_IMAGE_NAME=users-CRUD

help:
	@echo "Commands:";
	@echo "start-server: starts server on localhost and 8000 port"
	@echo "build-docker-image: builds docker image with name $(DOCKER_IMAGE_NAME)";
	@echo "run-docker-imange: runs image $(DOCKER_IMAGE_NAME)"

build:
	go mod download && GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

start-server:
	go run ./cmd/api/main.go

build-docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_PATH)

run-docker-image:
	docker run -it $(DOCKER_IMAGE_NAME)

swag:
	swag init -g internal/app/app.go

lint:
	golangci-lint run
