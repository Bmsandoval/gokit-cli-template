package hello_requests

import (
	"context"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/cli_binder"
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
	"github.com/spf13/cobra"
)

type GetHelloRequest struct {
}

//Decode the Push notification Req.
func MakeGetHelloRequestDecoder(appCtx appcontext.Context) cli_binder.RequestDecoder {
	return func(ctx context.Context, cmd *cobra.Command, args []string) (interface{}, error) {
		var Req GetHelloRequest

		return Req, nil
	}
}
