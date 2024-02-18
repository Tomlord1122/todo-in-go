# Todo list backend in Go 

Look up the Makefile and implement the following command 

```bash
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
```

The following picture is the test coverage. 


![image](https://github.com/Tomlord1122/todo-in-go/assets/79390871/bcc3e20f-f8ed-415a-b8dd-65c1560ac9db)

## Tech Skill
1. Go
2. Gin
3. viper
4. sqlc
