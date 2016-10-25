package dbconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDB(connStr string) (ddb *sql.DB, gdb *gorm.DB, err error) {
	ddb, err = sql.Open("mysql", connStr)
	gdb, err = gorm.Open("mysql", connStr)
	return
}
