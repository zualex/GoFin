# Golang finance (GoFin)

## Docker
```
make build
make up
```

Link: http://localhost:8080/
Static frontend: http://localhost:3000/

## Task tracker
https://trello.com/b/J1xAnypP/project-gofin

## Troubleshoot

### Postgres doesn't create database
```
make down
docker volume rm gofin_postgres
make up
docker logs -f gofin-postgres-1
```

### Container app doesn't up
1) Check files end (need LF):
- /golang/start.sh
- /golang/wait-for-it.sh
- /postgres/docker-postgresql-multiple-databases/create-multiple-postgresql-databases.sh