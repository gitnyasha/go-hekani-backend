package main

import (
	"database/sql"
	"log"

	"github.com/gitnyasha/go-hekani-backend/api"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
	"github.com/gitnyasha/go-hekani-backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.Server)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
