package types

type Amount int64

type PAN string

type Currency string

type Card struct {
	Id        int
	AccountID string
	Pan       PAN
	Balance   Amount
}

const (
	UZS Currency = "UZS"
	USD Currency = "USD"
	EUR Currency = "EUR"
	RUB Currency = "RUB"
)
