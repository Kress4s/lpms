package main

import (
	"log"
	"lpms/app"
	"lpms/config"
	"os"
	"os/signal"
	"syscall"
)

// @title 临安区政府投资项目管理后台API
// @version 1.0
// @description 临安区政府投资项目管理后台API

// @contact.name xiayoushuang
// @contact.email york-xia@gmail.com

// @schemes http https
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg := config.GetConfig()
	// go monitor.Start()
	go app.Run(cfg.Server.Port)
	// go app.RunJs(cfg.JsServer.Port)

	// 性能监控
	// go http.ListenAndServe(":9999", nil)

	// init log path
	// createDIR()

	s := waiForSignal()
	log.Fatalf("signal received, server closed, %v", s)
}

func waiForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
