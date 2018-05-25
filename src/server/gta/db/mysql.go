package db

import	(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
func Mysql() *sql.DB{
	if db==nil {
		db,_=sql.Open("mysql","gta:gta@/gta")
	}
	return db
}
