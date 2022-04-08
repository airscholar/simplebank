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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
