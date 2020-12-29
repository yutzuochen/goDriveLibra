package websocket

type Manager interface {
	Read() []byte
	Write(data []byte)
}
