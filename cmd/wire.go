//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tafo/rosa/config"
	"net/http"
)

func CreateWebServer() *http.Server {
	wire.Build(config.Container)
	return &http.Server{}
}
