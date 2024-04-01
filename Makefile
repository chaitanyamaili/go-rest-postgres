#  variables block
COMPOSE_FILE="./docker-compose.yaml"

# docker-compose commands
run:
	docker-compose -f ${COMPOSE_FILE} up -d

ls:
	docker-compose -f ${COMPOSE_FILE} ps

stop:
	docker-compose -f ${COMPOSE_FILE} down

remove: stop
	docker prune -f

remove-all: remove
	docker volume prune -f

# docker commands
logs:
	docker logs ${SERVICE}

# application sprcific commands
run-app-local:
	go run main.go

run-app:
	docker-compose -f ${COMPOSE_FILE} up -d app_server

run-db:
	docker-compose -f ${COMPOSE_FILE} up -d db_server

migrate:
	docker-compose -f ${COMPOSE_FILE} up -d db_flyway

db-connect: run-db
	docker-compose -f ${COMPOSE_FILE} exec db_server psql -h localhost -U postgres -d albums -W
