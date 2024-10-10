
# Spin up a postgres container
docker run --name postgres-demo-crud -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres


# migrate postgres: create
migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable \
    -source file://./cmd/migrate/migration/  \
    create -seq -dir ./cmd/migrate/migration/ -ext sql some_changes