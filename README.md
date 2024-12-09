# gopher-watch

Batteries included server monitoring software written exclusively in Go.

## Postgres Docker Container Setup

For the dev environment:

```
sudo docker run --name gopher-watch-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=gopher_watch -p 5432:5432 -d postgres
```

For the test environment:

```
sudo docker run --name gopher-watch-test-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=gopher_watch -p 5433:5432 -d postgres
```
