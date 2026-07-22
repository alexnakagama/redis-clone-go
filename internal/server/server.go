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

func (s *Server) Start() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	data := buffer[:n]

	return nil
}

func (s *Server) Close() error {

}
