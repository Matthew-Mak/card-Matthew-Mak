package types

type Amount int64

type PAN string

type Currency string

type Card struct {
	Id        int
	AccountID string
	Pan       PAN
	Balance   Amount
	IsActive  bool
}

const (
	UZS Currency = "UZS"
	USD Currency = "USD"
	EUR Currency = "EUR"
	RUB Currency = "RUB"
)
