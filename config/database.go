package config

import (
	"fmt"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql" // Driver mandatory hai
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *EnvConfig) {
	// DSN format: libsql://db-name.turso.io?authToken=your-token
	dsn := fmt.Sprintf("%s?authToken=%s", cfg.TURSO_DATABASE_URL, cfg.TURSO_AUTH_TOKEN)

	// GORM connection with Turso (libsql) driver
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "libsql",
		DSN:        dsn,
	}, &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Turso Database: %v", err)
	}

	log.Println("Successfully connected to Turso via GORM!")
	DB = db
}
