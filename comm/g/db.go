package g

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/humorliang/file-cms/comm/setting"
)

var (
	PsgCon *sql.DB
	err    error
)

func InitDB() {
	db := *setting.DbCfg
	dbSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	PsgCon, err = sql.Open(db.Type, dbSource)
	if err != nil {
		log.Fatal("[postgres connect open error]", err)
	}
	err = PsgCon.Ping()
	if err != nil {
		log.Fatal("[postgres connect error]", err)
	}
}
