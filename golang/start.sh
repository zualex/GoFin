#!/bin/bash

./wait-for-it.sh -s -t 0 ${DB_HOST}:${DB_PORT} -- \
    migrate -path /migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable up \
    && migrate -path /migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_TEST_NAME?sslmode=disable up

air