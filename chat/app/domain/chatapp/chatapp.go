package chatapp

import (
	"context"
	"net/http"

	"github.com/Brykl/usdl/foundation/web"
)

type app struct {
}

func newApp() *app {
	return &app{}
}

func (a *app) test(_ context.Context, _ *http.Request) web.Encoder {
	// Web socket implemeted here
	// just perform basic echo
	// Make sure we are handling connetion drops/issues (context)
	// How we will map a connection to user

	return status{
		Status: "ok",
	}

}
