package routes

import (
	"lpms/app/handlers/v1/auth"
	"lpms/config"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"

	_ "lpms/docs"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRoutes(app *iris.Application) {
	cfg := config.GetConfig()
	if cfg.DebugModel {
		app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler))
	}
	authApp := app.Party("/auth")
	mvc.New(authApp).Handle(auth.NewLoginHandler())
}
