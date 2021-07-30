package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func init () {
	DB = initDb()
}

func initDb() *sqlx.DB {
	dbPath := "/tmp/wg-dash.db"

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	return db
}

func MigrateSchema() {
	DB.MustExec(schema)

	// Seeding
	DB.MustExec(`INSERT
INTO users(email,password,is_admin,auth_type)
values("nham", "abc", 1, 0)
ON CONFLICT DO NOTHING
`)


}
