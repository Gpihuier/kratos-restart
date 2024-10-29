package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	v1 "kratos-restart/api/user/v1"
	"kratos-restart/app/admin/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewDiscovery, NewUserClient, NewUserRepo)

// Data .
type Data struct {
	uc v1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc v1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{uc: uc}, cleanup, nil
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserClient(c *conf.Service, r registry.Discovery, tp *tracesdk.TracerProvider) v1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.User.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewUserClient(conn)
}
