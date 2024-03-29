postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todo

dropdb:
	docker exec -it postgres12 dropdb todo

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
sqlc:
	sqlc generate
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown test sqlc server


