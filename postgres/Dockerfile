FROM postgres:latest

COPY ./schema/init.sql /docker-entrypoint-initdb.d/
COPY ./docker-postgresql-multiple-databases/create-multiple-postgresql-databases.sh /docker-entrypoint-initdb.d/
