package http_server

import "time"

type Option func(*Server)

func WithPort(port string) Option {
	return func(server *Server) {
		server.httpServer.Addr = port
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.httpServer.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.httpServer.WriteTimeout = t
	}
}

func WithShutdownTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.shutdownTimeout = t
	}
}
