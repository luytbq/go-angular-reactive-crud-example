
# Spin up a postgres container
docker run --name postgres-demo-crud -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
