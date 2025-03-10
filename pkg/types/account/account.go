package account

import "time"

type Account struct {
	Id       int
	UniqueId string
	Phone    string
	Balance  int
	CreateAt time.Time
}
