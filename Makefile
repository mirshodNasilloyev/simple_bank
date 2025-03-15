postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bankdb
dropdb:
	docker exec -it postgres12 dropdb simple_bankdb

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...
sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination=db/mock/store.go github.com/mirshodNasilloyev/simplebank-go/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 sqlc test server mock