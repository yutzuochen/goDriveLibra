package nats

import (
	"example/env"
	mq "example/messagequeue"

	nats "github.com/nats-io/nats.go"
)

const (
	TRACE_LEVEL = 1
)

type handler struct {
	connect *nats.Conn
}

func (hdr *handler) Subscribe(sbj string, cb func(m *nats.Msg)) error {
	_, err := hdr.connect.Subscribe(sbj, cb)
	return err
}

func (hdr *handler) Publish(sbj string, data []byte) error {
	return hdr.connect.Publish(sbj, data)
}

func NewNatsClusterMgr(cfg *env.NatsConfig) mq.ManagerNats {
	connect, _ := nats.Connect(cfg.NatsIP, nats.UserInfo(cfg.UserName, cfg.Password))
	return &handler{
		connect: connect,
	}
}
