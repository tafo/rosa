package config

import (
	"github.com/google/wire"
	"github.com/tafo/rosa/internal/api"
	"github.com/tafo/rosa/internal/auth"
)

var Container = wire.NewSet(
	NewConnection,
	auth.Provider,
	api.Provider,
)
