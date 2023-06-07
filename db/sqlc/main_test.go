package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries //*Queries is from db.go, Queries has DBTX interface type "db"

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource) //db 연결
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn) //New is from db.go

	os.Exit(m.Run())

} 