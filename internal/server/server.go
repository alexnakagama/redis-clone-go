package server

import (
	"net"
)

type Server struct {
	Address string
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		return nil, err
	}
}
