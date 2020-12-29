//要在這裡跑主流程
package nats_test

import (
	"example/env"
	mynats "example/messagequeue/nats"
	"fmt"
	"testing"

	nats "github.com/nats-io/nats.go"
)

var (
	cfg *env.NatsConfig
)

// func Test_nats(t *testing.T) {
// 	fmt.Println("funct:Test_main")
// 	end := make(chan struct{})
// 	mgr := mynats.NewNats()
// 	mgr.Publish("foo", []byte("Hello World !!!"))
// 	mgr.Subscribe("foo", func(m *nats.Msg) {
// 		fmt.Printf("Received a message: %s\n", string(m.Data))
// 		end <- struct{}{}
// 	})
// 	mgr.Publish("foo", []byte("Hello World !!!"))
// 	<-end

// }

func Test_cluster(t *testing.T) {
	cfg = &env.NatsConfig{
		"nats://localhost:4222, nats://localhost:4223, nats://localhost:4224", "jim", "password",
	}
	end := make(chan struct{})

	mgr := mynats.NewNatsClusterMgr(cfg)

	mgr.Publish("foo", []byte("Hello World ,bro !!!"))
	mgr.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		end <- struct{}{}
	})
	mgr.Publish("foo", []byte("Hello World , guys!!!"))
	<-end
}
