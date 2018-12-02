package main

import (
	"fmt"
	"github.com/urfave/cli"
)

var listCommand = cli.Command{
	Name:  "list",
	Usage: "lists bench",
	Action: func(context *cli.Context) error {
		fmt.Printf("list")
		return nil
	},
}
