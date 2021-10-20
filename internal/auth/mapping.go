package auth

import (
	"github.com/tafo/rosa/internal/auth/models"
	"time"
)

func (r RegisterRequest) ToEntity() models.Account {
	return models.Account{
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: time.Now().Unix(),
	}
}
