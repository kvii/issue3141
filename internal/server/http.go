package server

import (
	nt "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/kvii/issue3141/api/helloworld/v1"
	"github.com/kvii/issue3141/internal/conf"
	"github.com/kvii/issue3141/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.MyServiceService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterMyServiceHTTPServer(srv, greeter)

	// Add middleware for logging all the http request, include 404 routes.
	h := srv.Handler
	helper := log.NewHelper(log.With(logger, "module", "http"))
	srv.Handler = nt.HandlerFunc(func(w nt.ResponseWriter, r *nt.Request) {
		helper.Infow("msg", "request", "method", r.Method, "url", r.URL.String()) // no status code
		h.ServeHTTP(w, r)
	})

	return srv
}
