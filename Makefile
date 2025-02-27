postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bankdb
dropdb:
	docker exec -it postgres12 dropdb simple_bankdb

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
sqlc:
	sqlc generate


.PHONY: createdb dropdb postgres migrateup migratedown sqlc test