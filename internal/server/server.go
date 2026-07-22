package server

import (
	"net"
)

type Server struct {
	address  string
	listener net.Listener
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	server := &Server{
		address:  addr,
		listener: listener,
	}

	return server, nil
}
