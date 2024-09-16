package db

import (
	"context"
	"github.com/bolusarz/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal(err)
	}

	testDB, err = pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
