package messagequeue

import "github.com/nats-io/nats.go"

type ManagerNats interface {
	//Subscribe(topic string, callback func(message []byte) error) error
	//Publish(topic string, data interface{}) error
	Subscribe(topic string, callback func(m *nats.Msg)) error
	Publish(topic string, data []byte) error
}
