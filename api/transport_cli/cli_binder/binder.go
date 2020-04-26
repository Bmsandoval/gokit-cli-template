package cli_binder

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/spf13/cobra"
)
type ResponseEncoder func(ctx context.Context, response interface{}) error

type RequestDecoder func(ctx context.Context, cmd *cobra.Command, args []string) (interface{}, error)

type CliServer func()

func NewCliServer(ep endpoint.Endpoint, decoder RequestDecoder, encoder ResponseEncoder) (func(_ *cobra.Command, _ []string), error) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		decodedRequest, err := decoder(ctx, cmd, args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		response, err := ep(ctx, decodedRequest)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := encoder(ctx, response); err != nil {
			fmt.Println(err.Error())
			return
		}
	}, nil
}
