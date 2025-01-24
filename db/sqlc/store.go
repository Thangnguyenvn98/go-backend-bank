package db

import "database/sql"

// Store provides all functions to execute db queries and transactions, composition instead of inheritance
type Store struct {
	db *sql.DB
	*Queries
}

// Newstore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
