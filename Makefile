include .env
export

.PHONY: *

##
# Сборка
#
# Пример: make build
build:
	docker-compose build

##
# Сборка конкретного образа
##
build-%:
	docker-compose build --no-cache $*

##
# Запустить/пересоздать контейнеров из образов
#
# Пример: make up
up:
#	docker-compose up -d --no-build --force-recreate --remove-orphans
	docker-compose up -d --force-recreate --remove-orphans

##
# Остановить контейнер из образов
#
# Пример: make stop
stop:
	docker-compose stop


# Пример: make down
down:
	docker-compose down

down-volume:
	docker-compose down --volumes

##
# Запустить bash внутри контейнера
#
# Пример: make bash-php
##
bash-migrator:
	docker-compose run migrator bash
bash-go:
	docker-compose exec app bash
bash-%:
	docker-compose exec $* bash

migrate-up:
	docker-compose exec -ti app migrate -path=/migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up
	docker-compose exec -ti app migrate -path=/migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_TEST_NAME}?sslmode=disable up

migrate-down:
	docker-compose exec -ti app migrate -path=/migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable down
	docker-compose exec -ti app migrate -path=/migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_TEST_NAME}?sslmode=disable down