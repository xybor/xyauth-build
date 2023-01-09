DATABASE=./database
SERVER=./server
TEMPLATE=./template

DOCKER=sudo docker
DOCKER_COMPOSE=sudo docker-compose
DOCKERFILE=build/Dockerfile
DOCKER_COMPOSE_FILE=deploy/compose.yaml

IMAGE_NAME=xyauth

.PHONY: run clean docker-gen docker-build docker-start docker-stop docker-clean

database:
	go build -o $(DATABASE) ./cmd/database/*.go

template:
	go build -o $(TEMPLATE) ./cmd/template/*.go

server:
	go build -o $(SERVER) ./cmd/server/*.go

clean:
	rm -f $(DATABASE) $(SERVER) $(TEMPLATE)

run: database server
	$(DATABASE) migrate
	$(SERVER)

docker-gen: template
	$(TEMPLATE) $(DOCKERFILE).template
	$(TEMPLATE) $(DOCKER_COMPOSE_FILE).template

docker-build: clean docker-clean
	$(DOCKER) build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

docker-start:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up

docker-stop:
	-$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

docker-clean: docker-stop
	rm -f $(DOCKERFILE)
	rm -f $(DOCKER_COMPOSE_FILE)
	$(DOCKER) rmi -f $(IMAGE_NAME)
