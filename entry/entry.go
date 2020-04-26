package entry

import (
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/cli_routing"
	"github.com/bmsandoval/gokit-esque-cli-template/configs"
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
	"github.com/bmsandoval/gokit-esque-cli-template/services"
	"github.com/bmsandoval/gokit-esque-cli-template/services/grpc_service"
)

func Entry() {
	// Get Configs
	config, err := configs.Configure()
	if err != nil {
		panic(err) }

	grpcConnection, err := grpc_service.Start(*config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := grpc_service.Stop(); err != nil {
			panic(err)
		}
	}()

	// Build Context
	ctx := appcontext.Context{
		Config: *config,
		Grpc: *grpcConnection,
		// Redis
	}

	// Bundle Services
	serviceBundle, err := services.NewBundle(ctx)
	if err != nil {
		panic(err) }

	serviceBundle.GrpcSvc = *grpcConnection

	cli_routing.BundleAll(ctx, cli_routing.RootCmd, *serviceBundle)

	cli_routing.Execute()
}
