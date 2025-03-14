package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mirshodNasilloyev/simplebank-go/api"
	db "github.com/mirshodNasilloyev/simplebank-go/db/sqlc"
	"github.com/mirshodNasilloyev/simplebank-go/db/utils"
	"log"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatalf("error to connect database: %v", err)
	}
	defer conn.Close()
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("error to start server: %v", err)
	}
}
