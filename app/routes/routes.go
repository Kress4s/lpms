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

	implementParty := party.Party("/implement")
	implementApp := mvc.New(implementParty)
	implementApp.Handle(v1.NewImplementGovHandler())
	implementApp.Handle(v1.NewImpleIndustryHandler())
	implementApp.Handle(v1.NewGovProgressHandler())

	inspectParty := party.Party("/inspect")
	inspectApp := mvc.New(inspectParty)
	inspectApp.Handle(v1.NewReserveInspectHandler())
	inspectApp.Handle(v1.NewWindowHandler())
}
