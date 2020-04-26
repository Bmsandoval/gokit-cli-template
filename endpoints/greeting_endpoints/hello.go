package greeting_endpoints

import (
	"context"
	"github.com/bmsandoval/gokit-esque-cli-template/api/transport_cli/codecs/responses/hello_responses"
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
	"github.com/bmsandoval/gokit-esque-cli-template/services"
	"github.com/go-kit/kit/endpoint"
)

//Validation
//Find Campaing date
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return hello_responses.GetHelloResponse{
			Greeting: "hi there",
		}, nil
	}
}