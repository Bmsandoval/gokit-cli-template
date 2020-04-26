package hello_responses

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/cli_binder"
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
)

type GetHelloResponse struct {
	Greeting string `json:"Greeting"`
}

func MakeGetHelloResponseEncoder(appCtx appcontext.Context) cli_binder.ResponseEncoder {
	return func(ctx context.Context, response interface{}) error {
		res := response.(GetHelloResponse)

		if response != nil {
			data, err := json.MarshalIndent(res, "", "\t")
			if err != nil {
				return err
			}

			fmt.Println(string(data))
		}

		return nil
	}
}
