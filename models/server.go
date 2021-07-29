package models

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) Run() {
	for cmd := range s.commands {
		switch cmd.id {
		case

		}
	}
}


func (s *server) NewClient(c net.Conn) {

	log.Fatalf("New Client has cinnected %s:", c.RemoteAddr().String())

	cl := &client{
		conn:     c,
		nick:     "Anonymous",
		commands: s.commands,
	}

	cl.ReadInput()


}
