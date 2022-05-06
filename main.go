package main

import (
	"database/sql"
	"log"

	"github.com/gitnyasha/go-hekani-backend/api"
	db "github.com/gitnyasha/go-hekani-backend/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://hekanidb:topgear12@localhost:5432/hekani?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
