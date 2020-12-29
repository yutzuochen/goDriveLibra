package database

import (
	"example/dao"
)

type Manager interface {
	// Connect(config *env.DatabaseConfig) error
	GetPlayer(username string) []*dao.Player
	GetPlayerBalance(playerId string) (int, error, int64)
	// NewArango() (*arango.Handler, error)
}
