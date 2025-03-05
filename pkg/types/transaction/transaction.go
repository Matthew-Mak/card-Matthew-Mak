package transaction

import (
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
	"time"
)

type Transaction struct {
	Id            int
	CardId        string
	OperationType string
	Amount        cards.Amount
	Timestamp     time.Time
}
