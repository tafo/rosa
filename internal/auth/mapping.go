package auth

import (
	"time"
)

func (r RegisterRequest) ToEntity() Account {
	return Account{
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: time.Now().Unix(),
	}
}
