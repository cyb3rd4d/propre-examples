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

func (s *HTTPServer) Run(ctx context.Context) {
	go func() {
		logger.FromContext(ctx).Info("[server] start listening", "addr", s.addr)
		s.srv.ListenAndServe()
	}()

	<-ctx.Done()
	logger.FromContext(ctx).Info("[server] shutdown signal received")
	if ctx.Err() != nil {
		logger.FromContext(ctx).Error("[server] shutdown signal caused by an error", "previous", ctx.Err())
	}

	s.srv.Shutdown(ctx)
}
