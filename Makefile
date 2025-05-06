postgres:
	docker run --name my-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=secret \
  -p 5432:5432 \
  -v pgdata:/var/lib/postgresql/data \
  -d postgres:17.4-alpine3.21

postgres-cli:
	docker exec -it my-postgres psql -U postgres

createdb:
	docker exec -it my-postgres createdb -U postgres --owner=postgres simple_bank

dropdb:
	docker exec -it my-postgres dropdb -U postgres simple_bank

migrate-up:
	migrate -path db/migrations -database postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable up 
	
migrate-down:
	migrate -path db/migrations -database postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable down

sqlc:
	sqlc generate

.PHONY: postgres postgres-cli createdb dropdb migrate-up migrate-down sqlc