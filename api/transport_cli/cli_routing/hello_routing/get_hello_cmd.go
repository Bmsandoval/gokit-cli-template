package hello_routing

import (
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/cli_binder"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/cli_routing"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/codecs/requests/hello_requests"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/codecs/responses/hello_responses"
	"github.com/bmsandoval/gokit-esque-cli-template/endpoints/greeting_endpoints"
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
	"github.com/bmsandoval/gokit-esque-cli-template/services"
	"github.com/spf13/cobra"
)

var HelloWorldCmd = &cobra.Command{
	Use:     "helloworld",
	Aliases: []string{"hw"},
	Short:   "prints hello",
	Long:    `this is just a simple example of what you can do`,
}

func init() {
	cli_routing.Bundle(MakeGetHelloCliHandler())
}

func MakeGetHelloCliHandler() cli_routing.Bundlable {
	return func(appCtx appcontext.Context, rootCmd *cobra.Command, services services.Bundle) error {
		ep := greeting_endpoints.MakeGetHelloEndpoint(appCtx, services)
		rq := hello_requests.MakeGetHelloRequestDecoder(appCtx)
		rs := hello_responses.MakeGetHelloResponseEncoder(appCtx)

		cmd, err := cli_binder.NewCliServer(ep, rq, rs)
		if err != nil {
			return err
		}

		HelloWorldCmd.Run = cmd

		rootCmd.AddCommand(HelloWorldCmd)

		return nil
	}
}
