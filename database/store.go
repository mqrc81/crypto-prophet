package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(dsn string) (*Store, error) {

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging postgres database: %w", err)
	}

	defer db.Close()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return &Store{
		&SignalStore{DB: db},
		&MessageStore{DB: db},
	}, nil
}

// Store combines all stores.
type Store struct {
	*SignalStore
	*MessageStore
}
