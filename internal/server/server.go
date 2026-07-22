package server

import (
	"log"
	"net"

	"github.com/alexnakagama/redis-clone-go/internal/commands"
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
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		log.Println("client connected")

		err = handleConnection(conn)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	message := string(buffer[:n])

	log.Println("received: ", message)

	response, err := commands.Process(message)
	if err != nil {
		return err
	}

	data := []byte(response)

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
