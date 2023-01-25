package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sametdmr/simplebank/api"
	db "github.com/sametdmr/simplebank/db/sqlc"
	"github.com/sametdmr/simplebank/utility"
)

func main() {
	config, err := utility.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
