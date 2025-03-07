package payment

import (
	"time"

	types "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

type Payment struct {
	CreatingDate time.Time
	Currency     types.Currency
	AccountID    string
	CardNumber   string
	Category     string
	Balance      types.Amount
	ID           int
}
