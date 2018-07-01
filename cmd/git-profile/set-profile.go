package main

import (
	"errors"
	"gopkg.in/urfave/cli.v1"
	"git-profile/pkg/store"
)

type SetProfile struct {
	name string
	args []string
}

func (sp *SetProfile) Run() error {
	s, err := store.NewStoreForCurrentRepo()
	if err != nil {
		panic(err)
	}

	err = s.AttachProfile(sp.name, sp.args)
	if err != nil {
		panic(err)
	}
	return nil
}

func (sp *SetProfile) Parse(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New("missing NAME arg")
	}
	sp.name = c.Args().First()
	sp.args = c.Args().Tail()

	return nil
}
