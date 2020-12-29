package redis

type Manager interface {
	SetToken(playerID string, token string) error
	GetToken(playerID string) (string, error)
}
