package main

import (
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
)

type example struct{}

func (s *example) OnConnect(c *connection.Connection) {
	//log.Println(" OnConnect ： ", c.PeerAddr())
}
func (s *example) OnMessage(c *connection.Connection, buffer *ringbuffer.RingBuffer) (out []byte) {
	//log.Println("OnMessage")
	out = buffer.Bytes()
	buffer.RetrieveAll()

	return
}

func (s *example) OnClose() {
	//log.Println("OnClose")
}

func main() {
	handler := new(example)

	s, err := gev.NewServer(handler,
		gev.Network("tcp"),
		gev.Address(":1833"),
		gev.NumLoops(3),
		gev.MaxClient(100000))
	if err != nil {
		panic(err)
	}

	s.Start()
}