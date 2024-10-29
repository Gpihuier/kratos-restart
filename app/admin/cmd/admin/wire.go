//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-restart/app/admin/internal/biz"
	"kratos-restart/app/admin/internal/conf"
	"kratos-restart/app/admin/internal/data"
	"kratos-restart/app/admin/internal/server"
	"kratos-restart/app/admin/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Service, log.Logger, *trace.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
