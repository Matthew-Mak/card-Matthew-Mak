package account

import "time"

type Account struct {
	Id       int
	UniqueId string
	Phone    string
	CreateAt time.Time
}
