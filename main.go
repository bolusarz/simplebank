package main

import (
	"database/sql"
	"github.com/bolusarz/simplebank/api"
	db "github.com/bolusarz/simplebank/db/sqlc"
	"github.com/bolusarz/simplebank/util"
	_ "github.com/jackc/pgx/v5"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
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
