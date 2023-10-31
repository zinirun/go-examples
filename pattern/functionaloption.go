package main

import (
	"fmt"
)

// Server 구조체
type Server struct {
	Host string
	Port int
}

// Option 함수 타입
type Option func(*Server)

// NewServer 함수는 새로운 Server를 생성하고 옵션을 설정
func NewServer(options ...Option) *Server {
	server := &Server{
		Host: "localhost", // 기본 값
		Port: 8080,        // 기본 값
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func FunctionalOption() {
	s1 := NewServer()
	fmt.Printf("Server without options: Host=%s, Port=%d\n", s1.Host, s1.Port)

	s2 := NewServer(WithHost("127.0.0.1"))
	fmt.Printf("Server with host option: Host=%s, Port=%d\n", s2.Host, s2.Port)

	s3 := NewServer(WithPort(9000))
	fmt.Printf("Server with port option: Host=%s, Port=%d\n", s3.Host, s3.Port)

	s4 := NewServer(WithHost("example.com"), WithPort(80))
	fmt.Printf("Server with both options: Host=%s, Port=%d\n", s4.Host, s4.Port)
}
