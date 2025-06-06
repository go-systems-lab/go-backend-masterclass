DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres15 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateuprds:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

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

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate -path db/migration -database "$(DB_URL)" force $(version)

migratedownall:
	migrate -path db/migration -database "$(DB_URL)" down -all

db_reset: migratedownall migrateup

network:
	docker network create bank-network

proto:
	rm -rf pb/*
	rm -rf doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:8-alpine

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test coverage server mock new_migration migrateup1 migratedown1 migrateforce migratedownall db_reset network proto evans redis