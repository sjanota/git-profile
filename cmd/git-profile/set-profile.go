package main

import (
	"errors"
	"gopkg.in/urfave/cli.v1"
	"git-profile/pkg/store"
	"git-profile/pkg/repo"
)

type SetProfile struct {
	name string
	args []string
}

func (sp *SetProfile) Run() error {
	r, err := repo.OpenCurrent()
	if err != nil {
		panic(err)
	}

	s, err := store.NewStore(r)
	if err != nil {
		panic(err)
	}

	err = s.SetProfile(sp.name, sp.args)
	if err != nil {
		panic(err)
	}
	return nil
}

func (sp *SetProfile) Parse(c *cli.Context) error {
	if c.NArg() < 2 {
		return errors.New("please provide profile name")
	}
	sp.name = c.Args().First()
	sp.args = c.Args().Tail()

	return nil
}
