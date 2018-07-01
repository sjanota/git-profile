package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

type Cmd interface {
	Parse(c *cli.Context) error
	Run() error
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command {
		{
			Name: "set",
			Action: Action(&SetProfile{}),
			SkipFlagParsing: true,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Action(cmd Cmd) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		err := cmd.Parse(ctx)
		if err != nil {
			return err
		}

		return cmd.Run()
	}
}


