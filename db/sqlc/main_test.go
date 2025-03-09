package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirshodNasilloyev/simplebank-go/db/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config file: %v", err)
	}
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("error to connect database: %v", err)
	}
	defer testDB.Close()
	testQueries = New(testDB)
	fmt.Println("successfully connected to database")
	os.Exit(m.Run())
}
