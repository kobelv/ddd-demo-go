package http

import (
	"fmt"
	"net/http"
	"time"
)

type Conf struct {
	Addr         string
	ReadTimeout  time.Duration // 单位 ms
	WriteTimeout time.Duration // 单位 ms
	IdleTimeout  time.Duration
}

type Server struct {
	*http.Server
}

func NewServer(c *Conf, handler http.Handler) (*Server, error) {
	s := &Server{
		Server: &http.Server{},
	}
	s.Addr = c.Addr
	s.Handler = handler
	s.ReadTimeout = c.ReadTimeout
	s.WriteTimeout = c.WriteTimeout
	s.IdleTimeout = c.IdleTimeout
	return s, nil
}

func (s *Server) Start() error {
	fmt.Println(s.Addr)
	return s.ListenAndServe()
}
