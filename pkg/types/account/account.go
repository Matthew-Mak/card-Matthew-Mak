package account

import (
	"errors"
	cards "github.com/Matthew-Mak/card-Matthew-Mak/v2/pkg/types/card"
	"time"
)

type Phone string

type Account struct {
	CreateAt time.Time
	UniqueID string
	Phone    Phone
	ID       int
}

var (
	ErrNegativeAmount     = errors.New("error: the amount can't be less than 0")
	ErrMoreThanExpected   = errors.New("error: the amount can't higher than 100 000 000")
	ErrCardNotActive      = errors.New("error: card is not active")
	ErrNotEnoughBalance   = errors.New("error: not enough money on the balance")
	ErrNegativeBalance    = errors.New("error: the balance after transfer is negative")
	ErrTooMuchOnBalance   = errors.New("error: the money on balance can't be more than 50 000 000")
	MaximumBalanceDeposit = 50_000_000
	MaximumAmountWithdraw = 100_000_000
)

func Init() Account {
	return Account{CreateAt: time.Now(), UniqueID: "", Phone: Phone(""), ID: 0}
}

func (a *Account) Withdraw(card *cards.Card, amount cards.Amount) (err error) {
	if amount < 0 {
		return ErrNegativeAmount
	}
	if amount > cards.Amount(MaximumAmountWithdraw) {
		return ErrMoreThanExpected
	}
	if !card.IsActive {
		return ErrCardNotActive
	}
	if amount > cards.Amount(card.Balance) {
		return ErrNotEnoughBalance
	}
	card.Balance -= amount
	return nil
}

func (a *Account) Deposit(card *cards.Card, amount cards.Amount) (err error) {
	if amount < 0 {
		return ErrNegativeAmount
	}
	if !card.IsActive {
		return ErrCardNotActive
	}
	if amount+cards.Amount(card.Balance) < 0 {
		return ErrNegativeBalance
	}
	if amount+cards.Amount(card.Balance) > cards.Amount(MaximumBalanceDeposit) {
		return ErrTooMuchOnBalance
	}
	card.Balance += amount
	return nil
}

var (
	ErrAccountNotFound = errors.New("error: account not found")
)

func (s *Account) Authorize(phone Phone, accounts []Account) (*Account, error) {
	for _, account := range accounts {
		if account.Phone == phone {
			return &account, nil
		}
	}
	return nil, ErrAccountNotFound
}
