package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/go-systems-lab/go-backend-masterclass/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// conn, err := pgx.Connect(context.Background(), dbSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	testQueries = New(testDB)
	os.Exit(m.Run())
}
