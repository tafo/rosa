package config

import "github.com/google/wire"

var Container = wire.NewSet(
	NewConnection,
)