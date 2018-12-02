package main

import (
	"fmt"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "run bench",
	Action: func(context *cli.Context) error {
		factory, err := NewFactory(context)
		if err != nil {
			return err
		}

		bench, err := factory.Create()
		if err != nil {
			return err
		}

		result, err := bench.Run()
		if err != nil {
			return err
		}

		fmt.Println(result)

		return nil
	},
}
