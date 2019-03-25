package g

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
	dbType     = "postgres"
	dbHOST     = "localhost"
	dbPORT     = 5432
	dbUSER     = "postgres"
	dbPASSWORD = "123456"
	dbNAME     = "fcms"
)

var (
	PsgCon *sql.DB
	err    error
)

func InitDB() {
	dbSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHOST, dbPORT, dbUSER, dbPASSWORD, dbNAME)
	PsgCon, err = sql.Open(dbType, dbSource)
	if err != nil {
		log.Fatal("[postgres connect open error]", err)
	}
	err=PsgCon.Ping()
	if err != nil {
		log.Fatal("[postgres connect error]", err)
	}
}
