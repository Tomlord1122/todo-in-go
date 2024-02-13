// main_test.go 作為測試的入口，會引入所有的測試檔案，並且提供一些共用的函式，例如 createRandomTodo 用來建立隨機的 todo 資料。
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/todo?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource) // return a connection object and an error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn) // create a new Queries object
	// m.Run() runs the tests
	// os.Exit() exits the program with the status code returned by m.Run()
	os.Exit(m.Run()) // run the tests
}
