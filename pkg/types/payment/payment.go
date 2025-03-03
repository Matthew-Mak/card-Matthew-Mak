package payment

import (
	"time"

	types "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
)

type Payment struct {
	Id           int
	AccountID    string
	CardNumber   string
	Balance      types.Amount
	Currency     types.Currency
	CreatingDate time.Time
	Category     string
}

// func GetPaymentByCategory(payments []Payment, category string) []Payment {
// 	for i := 0; i < len(payments); i++ {
// 		if payments[i].Category != category {
// 			payments = append(payments[:i], payments[i+1:]...)
// 		}
// 	}
// 	return payments
// }
