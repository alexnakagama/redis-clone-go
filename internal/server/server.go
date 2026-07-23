package server

import (
	"io"
	"log"
	"net"

	"github.com/alexnakagama/redis-clone-go/internal/commands"
	"github.com/alexnakagama/redis-clone-go/internal/store"
)

type Server struct {
	address  string
	listener net.Listener
	store    *store.Store
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	server := &Server{
		address:  addr,
		listener: listener,
		store:    store.NewStore(),
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

		go func(conn net.Conn) {
			err := handleConnection(conn, s.store)
			if err != nil {
				log.Println(err)
			}
		}(conn)
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func handleConnection(conn net.Conn, s *store.Store) error {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("client disconnected")
				return nil
			}

			return err
		}

		message := string(buffer[:n])

		log.Println("received: ", message)

		response, err := commands.Process(message, s)
		if err != nil {
			return err
		}

		data := []byte(response)

		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
}
