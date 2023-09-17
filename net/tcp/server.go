package tcp

import (
	"fmt"
	"net"

	"github.com/Curtion/yu/net/conn"
)

type Server struct {
	listener net.Listener
	port     int
	connMgr  *conn.ConnManager
	handle   func(*conn.Conn)
}

func NewServer(port int) *Server {
	return &Server{
		port:    port,
		connMgr: conn.NewConnManager(),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}
	s.listener = listener
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn := conn.NewConn(c, c.RemoteAddr().String())
		conn.SetHandle(s.handle)
		s.connMgr.AddConn(conn)
		go conn.Start()
	}
}

func (s *Server) SetHandle(handle func(*conn.Conn)) {
	s.handle = handle
}
