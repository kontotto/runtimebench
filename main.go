package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "runtimebench"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		listCommand,
		runCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
