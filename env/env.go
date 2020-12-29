package env

var Setting Config

type Config struct {
	Database *DatabaseConfig
	Redis    *RedisConfig
	Nats     *NatsConfig
}
type DatabaseConfig struct {
	AddressIP string
	//Port    int
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type NatsConfig struct {
	NatsIP   string
	UserName string
	Password string
}
