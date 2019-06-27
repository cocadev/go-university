package server

import (
	"../api"
	"../config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run echo framework...
func Run() {
	printLogo()

	ch := make(chan struct{})
	startServer(ch)
	<-ch
}

func startServer(exitCh chan struct{}) {
	e := echo.New()

	// Global middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Start Route API
	api.RouteAPI(e)

	if err := e.Start(":" + config.Port); err != nil {
		e.Logger.Fatal(err)
	}
	exitCh <- struct{}{}
}

func printLogo() {
	println("*********************************************")
	println("*-------------------------------------------*")
	println("*       AppLive StarterKit WebService       *")
	println("*-------------------------------------------*")
	println("* Author: DaGe Tian")
	println("* Version: 1.0.0")
	println("* Host:", config.HostURL)
	println("* Port:", config.Port)
	println("* DB: MongoDB")
	println("* DBHost:", config.MongoHost)
	println("*********************************************")
}
