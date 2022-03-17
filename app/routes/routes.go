package routes

import (
	v1 "lpms/app/handlers/v1"
	"lpms/app/handlers/v1/auth"
	"lpms/app/middlewares"
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

	party := app.Party("/api/v1")
	party.Use(middlewares.Auth().Serve)

	reserveParty := party.Party("/reserve")
	reserveApp := mvc.New(reserveParty)
	reserveApp.Handle(v1.NewReserveHandler())

	objectParty := party.Party("/object")
	objectApp := mvc.New(objectParty)
	objectApp.Handle(v1.NewObjectHandler())

	implementGovParty := party.Party("/implement/gov")
	implementGovApp := mvc.New(implementGovParty)
	implementGovApp.Handle(v1.NewImplementGovHandler())
}
