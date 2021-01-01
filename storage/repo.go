package storage

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

// New creates new Repo structure
func NewRepo(db *sql.DB) *Repo {
	fmt.Print("Data saved to Repo \n")
	return &Repo{
		db: db,
	}
}
