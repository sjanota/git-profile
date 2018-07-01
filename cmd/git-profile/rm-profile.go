package main

import (
	"gopkg.in/urfave/cli.v1"
	"errors"
	"git-profile/pkg/store"
)

type RmProfile struct {
	name string
}

func (rp *RmProfile) Parse(c *cli.Context) error {
	if c.NArg() != 1 {
		return errors.New("missing NAME arg")
	}
	rp.name = c.Args().First()
	return nil
}

func (rp *RmProfile) Run() error {
	s, err := store.NewStoreForCurrentRepo()
	if err != nil {
		panic(err)
	}

	err = s.DetachProfile(rp.name)
	if err != nil {
		panic(err)
	}
	return nil
}


