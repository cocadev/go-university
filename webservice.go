package main

import (
	"os"

	"./script"
	"./server"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "AppLive StarterKit"
	app.Version = "0.0.1"
	app.Author = "DaGe Tian"
	app.Commands = script.Commands()
	app.Action = func(c *cli.Context) {
		println("Running Server...")
		server.Run()
	}
	app.Run(os.Args)
}
