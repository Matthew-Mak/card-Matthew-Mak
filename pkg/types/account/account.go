package account

import "time"

type Account struct {
	CreateAt time.Time
	UniqueID string
	Phone    string
	ID       int
}
