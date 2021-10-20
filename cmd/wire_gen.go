// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tafo/rosa/config"
	"github.com/tafo/rosa/internal/api"
	"github.com/tafo/rosa/internal/auth"
	"github.com/tafo/rosa/internal/auth/helper"
	"github.com/tafo/rosa/internal/auth/middlewares"
	"net/http"
)

// Injectors from wire.go:

func CreateWebServer() *http.Server {
	engine := api.NewRouter()
	db := config.NewConnection()
	accountRepository := auth.NewAccountRepository(db)
	bcryptWrapper := helper.NewBcryptWrapper()
	hmacSecret := helper.NewHMACSecret()
	jwtWrapper := helper.NewJWTWrapper(hmacSecret)
	accountManager := auth.NewAccountManager(accountRepository, bcryptWrapper, jwtWrapper)
	accountHandler := auth.NewAccountHandler(accountManager)
	authMiddleware := middlewares.NewAuthMiddleware(jwtWrapper)
	server := api.NewHttpServer(engine, accountHandler, authMiddleware)
	return server
}
