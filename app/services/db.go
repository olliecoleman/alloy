package services

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// InitDB sets up the database
func InitDB() *sqlx.DB {
	dsn, err := envy.MustGet("DATABASE_URL")

	if err != nil {
		log.Fatalln(err)
	}
	DB = sqlx.MustConnect("postgres", dsn)
	return DB
}
