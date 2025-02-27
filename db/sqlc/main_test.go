package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbSource = "postgres://root:secret@localhost:5432/simple_bankdb?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatalf("error to connect database: %v", err)
	}
	defer testDB.Close()
	testQueries = New(testDB)
	fmt.Println("successfully connected to database")
	os.Exit(m.Run())
}
