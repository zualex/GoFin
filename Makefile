##
# Сборка
#
# Пример: make build
build:
	docker-compose build

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