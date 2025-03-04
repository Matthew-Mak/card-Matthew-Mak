package transaction

import (
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
	"time"
)

type transaction struct {
	id            int
	cardId        string
	operationType bool
	amount        cards.Amount
	timestamp     time.Time
}
