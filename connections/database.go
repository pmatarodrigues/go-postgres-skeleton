package connections

import (
	"database/sql"
	"fmt"

	// empty statement because it's not used directly
	// only imported for as concrete sql interface
	_ "github.com/lib/pq"
	"github.com/pmatarodrigues/go-starter/config"
)

func NewDatabase(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	// Get database connection data
	// From config/settings.go
	// postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database)
	postgresInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Database)

	// Connecto to PostgreSQL DB
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database. %v", err)
	}

	fmt.Print("Connected to database \n")
	return db, nil
}
