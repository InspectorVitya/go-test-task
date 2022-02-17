package httpserver

import (
	"context"
	"github.com/inspectorvitya/go-test-task/internal/app"
	"github.com/inspectorvitya/go-test-task/internal/config"
	"net"
	"net/http"
)

type Server struct {
	HTTPServer *http.Server
	App        *app.Proxy
}

func New(cfg *config.Config, app *app.Proxy) *Server {
	server := &Server{
		HTTPServer: &http.Server{
			Addr: net.JoinHostPort("", cfg.PortHTTP),
		},
		App: app,
	}

	return server
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.reversProxy)
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
