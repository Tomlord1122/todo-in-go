package main

import (
	"database/sql"
	"log"

	"github.com/Tomlord1122/todo-in-go/db/api"
	db "github.com/Tomlord1122/todo-in-go/db/sqlc"
	"github.com/Tomlord1122/todo-in-go/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	query := db.New(conn)
	server := api.NewServer(query)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
