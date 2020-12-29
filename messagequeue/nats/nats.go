package nats

import (
	"example/env"
	mq "example/messagequeue"

	nats "github.com/nats-io/nats.go"
)

type NatsHandler struct {
	connect *nats.Conn
}

func NewNats() mq.ManagerNats {
	hdr := NatsHandler{}
	hdr.connect, _ = nats.Connect(env.Setting.Nats.NatsIP)
	return hdr
}

func (hdr NatsHandler) Subscribe(sbj string, cb func(m *nats.Msg)) error {
	_, err := hdr.connect.Subscribe(sbj, cb)
	return err
}
func (hdr NatsHandler) Publish(sbj string, data []byte) error {
	return hdr.connect.Publish(sbj, data)
}
