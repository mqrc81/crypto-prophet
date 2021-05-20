package bot

import (
	"time"
)

type Signal struct {
	ID               string    `db:"id"`
	CryptoBase       string    `db:"crypto_currency"` // MTC
	CryptoQuote      string    `db:"crypto_pair"`     // USDT
	Amount           float64   `db:"amount"`
	Increments       []float64 `db:"increments"`
	CurrentIncrement int       `db:"current_increment"` // Stop - Initial - T1 - T2 - T3 - T4 - T5 - T6
	OrderProfit      string    `db:"order_profit"`
	OrderLoss        string    `db:"order_loss"`
	DateStart        time.Time `db:"date_start"`
	DateEnd          time.Time `db:"date_end"`
	Profit           float64   `db:"profit"`
	Closed           bool      `db:"closed"`
}
