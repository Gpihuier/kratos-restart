package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/gorilla/handlers"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	v1 "kratos-restart/api/admin/v1"
	"kratos-restart/app/admin/internal/conf"
	"kratos-restart/app/admin/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, sa *service.AdminService, tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	v1.RegisterAdminHTTPServer(srv, sa)
	return srv
}
