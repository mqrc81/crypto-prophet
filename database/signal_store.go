package database

import (
	"fmt"

	"github.com/mqrc81/crypto-prophet/bot"

	"github.com/jmoiron/sqlx"
)

type SignalStore struct {
	*sqlx.DB
}

func (store *SignalStore) GetSignals(closed bool) ([]bot.Signal, error) {
	var signals []bot.Signal

	if err := store.Select(&signals, `SELECT * FROM signals WHERE closed = $1`, closed); err != nil {
		return []bot.Signal{}, fmt.Errorf("error getting trade signals: %w", err)
	}

	return signals, nil
}

func (store *SignalStore) CreateSignal(s bot.Signal) error {

	if _, err := store.Exec(`INSERT INTO signals (crypto_currency, crypto_pair, amount, increments, current_increment, 
                     				order_profit, order_loss, date_start, profit) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		s.CryptoCurrency, s.CryptoPair, s.Amount, s.Increments, s.CurrentIncrement, s.OrderProfit, s.OrderLoss,
		s.DateStart, s.Profit); err != nil {
		return fmt.Errorf("error updating trade signal: %w", err)
	}

	return nil
}

func (store *SignalStore) UpdateSignal(s bot.Signal) error {

	if _, err := store.Exec(`UPDATE signals 
								SET crypto_currency = $1, 
									crypto_pair = $2, 
									amount = $3, 
									increments = $4, 
									current_increment = $5, 
									order_profit = $6,
								    order_loss = $7,
									date_start = $8, 
									date_end = $9,
									profit = $10,
								    closed = $11
								WHERE id = $12`,
		s.CryptoCurrency, s.CryptoPair, s.Amount, s.Increments, s.CurrentIncrement, s.OrderProfit, s.OrderLoss,
		s.DateStart, s.DateEnd, s.Closed, s.ID); err != nil {
		return fmt.Errorf("error updating trade signal: %w", err)
	}

	return nil
}
