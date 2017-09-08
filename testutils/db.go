package testutils

import (
	"fmt"
	"log"

	"github.com/olliecoleman/alloy/app/services"
	"github.com/gobuffalo/envy"
	"github.com/pressly/goose"
)

func SetupDB() {
	td := "postgres://postgres:@localhost/alloy_test?sslmode=disable"
	envy.Set("DATABASE_URL", td)
	migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	if err != nil {
		log.Fatal("MIGRATIONS_DIR variable not set")
	}

	services.InitDB()
	if err := goose.Run("up", services.DB.DB, migrationsDir); err != nil {
		log.Fatal("ERR", err)
	}
}

func DropDB() {
	migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	if err != nil {
		log.Fatal("MIGRATIONS_DIR variable not set")
	}
	if err := goose.Run("down", services.DB.DB, migrationsDir); err != nil {
		log.Fatal("ERR", err)
	}
}

func ResetTable(tablename string) {
	log.Printf("DELETE from %s", tablename)
	services.DB.Exec(fmt.Sprintf(`DELETE FROM %s`, tablename))
}
