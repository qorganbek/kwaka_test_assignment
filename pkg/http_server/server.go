package http_server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{Handler: handler}

	s := &Server{
		httpServer: httpServer,
		notify:     make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.httpServer.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)

	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
