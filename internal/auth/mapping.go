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

func (request LoginRequest) ToEntity() Account {
	return Account{
		Email: request.Email,
		Password: request.Password,
	}
}
