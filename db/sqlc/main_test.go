// main_test.go 作為測試的入口，會引入所有的測試檔案，並且提供一些共用的函式，例如 createRandomTodo 用來建立隨機的 todo 資料。
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Tomlord1122/todo-in-go/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
