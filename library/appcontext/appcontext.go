package appcontext

import (
	"context"
	"github.com/bmsandoval/gokit-esque-cli-template/configs"
	"github.com/bmsandoval/gokit-esque-cli-template/services/grpc_service"
)

type Context struct {
	Config    configs.Configuration
	GoContext context.Context
	Grpc      grpc_service.Connection
}