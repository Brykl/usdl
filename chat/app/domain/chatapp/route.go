package chatapp

import (
	"net/http"

	"github.com/Brykl/usdl/foundation/web"
)

func Routes(app *web.App) {
	api := newApp()

	app.HandlerFunc(http.MethodGet, "", "/test", api.test)
}
