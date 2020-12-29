package main

import (
	"example/env"
)

var cfg_Arrango *env.DatabaseConfig
var cfg_Nats *env.NatsConfig
var cfg_redis *env.RedisConfig

// func init() {
// 	env.Setting.Database = &env.DatabaseConfig{AddressIP: "http://10.200.6.37:8529"}
// 	env.Setting.Redis = &env.RedisConfig{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	}
// }

/* NATS

func main() {
	cfg_Nats = &env.NatsConfig{
		"nats://localhost:4222, nats://localhost:4223, nats://localhost:4224", "jim", "password",
	}
	end := make(chan struct{})

	mgr := mynats.NewNatsClusterMgr(cfg_Nats)

	mgr.Publish("foo", []byte("Hello World ,bro !!!"))
	mgr.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		end <- struct{}{}
	})
	mgr.Publish("foo", []byte("Hello World , guys!!!"))
	<-end
}
*/

/* redis


func main() {
	cfg_redis = &env.RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	r := redis.NewRedis(cfg_redis)
	err := r.SetToken("paradise", "13156764")
	if err != nil {
		panic(err)
	}
	token, err := r.GetToken("paradise")
	if err != nil {
		panic(err)
	}
	fmt.Println("the token is: ", token)
}
*/

/* Arrango

func main() {
	cfg_Arrango = &env.DatabaseConfig{"http://10.200.6.37:8529"}
	db, err := arango.NewArango(cfg_Arrango)
	if err != nil {
		panic(err)
	}
	res := db.GetPlayer("jerome")
	fmt.Println("res: ", res)
	for _, v := range res {
		fmt.Println(v)
	}

}
*/
