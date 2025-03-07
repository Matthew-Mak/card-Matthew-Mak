package account

import (
	"testing"
	"time"
)

func TestAccount_Authorize(t *testing.T) {
	tests := []struct {
		phone    Phone
		accounts []Account
		err      error
		expected *Account
	}{
		{phone: Phone("123456"), accounts: []Account{
			{CreateAt: time.Now(), UniqueID: "123456", Phone: Phone("123456"), ID: 1},
			{CreateAt: time.Now(), UniqueID: "123457", Phone: Phone("123457"), ID: 2},
		}},
		{phone: Phone("123456"), accounts: []Account{
			{CreateAt: time.Now(), UniqueID: "123456", Phone: Phone("123458"), ID: 1},
			{CreateAt: time.Now(), UniqueID: "123457", Phone: Phone("123457"), ID: 2},
		}},
	}
	for _, test := range tests {
		acc := Init()
		result, err := acc.Authorize(test.phone, test.accounts)
		if err != test.err {
			t.Errorf("Authorize(%v, %v): expected %v, got %v", test.phone, test.accounts, test.err, err)
		}
		if result != test.expected {
			t.Errorf("Authorize(%v, %v): expected %v, got %v", test.phone, test.accounts, test.expected, result)
		}
	}
}
