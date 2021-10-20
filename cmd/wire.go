package main

import (
	"github.com/google/wire"
	"github.com/tafo/rosa/config"
)

func InitContainer() {
	wire.Build(config.Container)
}
