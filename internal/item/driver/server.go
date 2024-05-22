package driver

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/driver/logger"
)

type HTTPServer struct {
	addr string
	srv  *http.Server
}

func NewHTTPServer(addr string, router http.Handler) *HTTPServer {
	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &HTTPServer{addr: addr, srv: &srv}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	logger.FromContext(ctx).Info("[server] start listening", "addr", s.addr)
	return s.srv.ListenAndServe()
}
