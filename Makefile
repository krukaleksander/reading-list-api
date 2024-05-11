export

COMPOSE = docker-compose -f docker-compose.yml

dev-build:
	$(COMPOSE) build

dev-start:
	$(COMPOSE) up -d

dev-stop:
	$(COMPOSE) down

dev-prune:
	$(COMPOSE) down --rmi all --volumes --remove-orphans
