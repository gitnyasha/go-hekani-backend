package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gitnyasha/go-hekani-backend/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	testDB, err = sql.Open(config.Driver, config.Source)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
