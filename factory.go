package main

import (
	"fmt"
	"github.com/urfave/cli"
)

type Factory interface {
	Create() (Bench, error)
}

func NewFactory(context *cli.Context) (Factory, error) {
	benchName := context.Args().First()
	switch benchName {
	case "dummy":
		return &dummyFactory{}, nil
	default:
		return nil, fmt.Errorf("benchname: \"%s\" is not found", benchName)
	}
}
