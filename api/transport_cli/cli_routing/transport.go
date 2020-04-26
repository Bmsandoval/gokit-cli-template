package cli_routing

import (
	"github.com/bmsandoval/gokit-esque-cli-template/library/appcontext"
	"github.com/bmsandoval/gokit-esque-cli-template/services"
	"github.com/spf13/cobra"
)

var cfgFile string


type ServerContext struct {
	AppCtx appcontext.Context
	Bundle services.Bundle
}

type Bundlable func(appCtx appcontext.Context, rootCmd *cobra.Command, services services.Bundle) error

var Bundlables []Bundlable

func Bundle(bundlable Bundlable) {
	Bundlables = append(Bundlables, bundlable)
}

func BundleAll(appCtx appcontext.Context, rootCmd *cobra.Command, services services.Bundle) {
	for _, bundlable := range Bundlables {
		bundlable(appCtx, rootCmd, services)
	}
}
