package transaction

import (
	"time"

	types "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

type Transaction struct {
	Timestamp     time.Time
	CardID        string
	OperationType string
	Amount        types.Amount
	ID            int
}
