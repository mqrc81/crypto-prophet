package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MessageStore struct {
	*sqlx.DB
}

func (store *MessageStore) GetMessages() ([]string, error) {
	var messages []string

	// Execute prepared statement
	if err := store.Select(&messages, `SELECT * FROM messages`); err != nil {
		return []string{}, fmt.Errorf("error getting discord messages: %w", err)
	}

	return messages, nil
}

func (store *MessageStore) DeleteMessages() error {

	if _, err := store.Exec(`DELETE FROM messages`); err != nil {
		return fmt.Errorf("error deleting discord messages: %w", err)
	}

	return nil
}
