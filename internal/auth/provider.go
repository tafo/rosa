package auth

import (
	"github.com/google/wire"
	"github.com/tafo/rosa/internal/auth/helper"
	"github.com/tafo/rosa/internal/auth/middlewares"
	"github.com/tafo/rosa/internal/auth/models"
)

var Provider = wire.NewSet(
	NewAccountRepository,
	wire.Bind(new(models.Repository), new(AccountRepository)),
	helper.NewHMACSecret,
	helper.NewJWTWrapper,
	wire.Bind(new(models.JWTWrapper), new(helper.JWTWrapper)),
	NewAccountManager,
	wire.Bind(new(Manager), new(AccountManager)),
	NewAccountHandler,
	helper.NewBcryptWrapper,
	wire.Bind(new(models.BcryptWrapper), new(helper.BcryptWrapper)),
	middlewares.NewAuthMiddleware,
	middlewares.NewAdminMiddleware,

)
