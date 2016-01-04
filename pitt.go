package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pitt"
	app.Usage = "go project managment tool"
	app.Run(os.Args)
}