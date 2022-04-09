package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbDriver = "postgres"
	dbSource = goDotEnvVariable("DB_URL")
)

var testQueries *Queries
var testDb *sql.DB

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the db", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
