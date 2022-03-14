package app

import (
	"fmt"
	"log"
	"lpms/app/routes"
	"lpms/config"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Run(port int) {
	if err := newApp().Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", port))); err != nil {
		log.Fatal("Web server run failed, err is ", err.Error())
	}
}

// func RunJs(port int) {
// 	if err := newJSApp().Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", port))); err != nil {
// 		log.Fatal("Js server run failed, err is ", err.Error())
// 	}
// }

// newApp
func newApp() *iris.Application {
	cfg := config.GetConfig()
	app := iris.New()
	iris.WithOptimizations(app)
	// app.Use(middlewares.Recover())
	// if cfg.DebugModel {
	// 	app.Use(IrisLogger())
	// } else {
	// 	// log recode
	// 	logFile := tools.NewLogFile()
	// 	app.Logger().SetOutput(logFile)
	// }

	app.Use(iris.Compression)
	// app.Use(middlewares.RecordSystemLog())
	// 跨域规则
	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins: cfg.Server.Cors.AllowedOrigins,
		AllowedMethods: []string{
			iris.MethodHead,
			iris.MethodGet,
			iris.MethodPost,
			iris.MethodPut,
			iris.MethodPatch,
			iris.MethodDelete,
			iris.MethodOptions,
		},
		AllowedHeaders:     cfg.Server.Cors.AllowedHeaders,
		ExposedHeaders:     []string{},
		AllowCredentials:   true,
		OptionsPassthrough: false,
	}))
	routes.RegisterRoutes(app)
	return app
}
