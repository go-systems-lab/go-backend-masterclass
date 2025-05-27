postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

coverage:
	go test -v -cover -coverprofile=coverage.out ./...
	gcov2lcov -infile=coverage.out -outfile=lcov.info

server:
	go run main.go

mock:
	mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mockdb -aux_files=github.com/go-systems-lab/go-backend-masterclass/db/sqlc=db/sqlc/querier.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test coverage server mock