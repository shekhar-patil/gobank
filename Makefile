postgres:
	docker run --name my-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=secret \
  -p 5432:5432 \
  -v pgdata:/var/lib/postgresql/data \
  -d postgres:17.4-alpine

postgres-cli:
	docker exec -it my-postgres psql -U postgres

createdb:
	docker exec -it my-postgres createdb -U postgres --owner=postgres simple_bank

dropdb:
	docker exec -it my-postgres dropdb -U postgres simple_bank

migrate-up:
	migrate -path db/migrations -database postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable up 

migrate-down:
	docker exec -it my-postgres migrate -path /migrations -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" down

PHONY: createdb dropdb