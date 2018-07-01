package main

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
	"gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"errors"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("please provide profile name"))
	}
	profileName := os.Args[1]
	profileArgsVal := ""

	if len(os.Args) > 2 {
		profileArgs := os.Args[2:]
		profileArgsVal = strings.Join(profileArgs, " ")
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	r, err := git.PlainOpen(cwd)
	if err != nil {
		panic(err)
	}

	c, err := r.Config()
	if err != nil {
		panic(err)
	}

	c.Raw = c.Raw.SetOption("profiles", config.NoSubsection, profileName, profileArgsVal)

	err = r.Storer.SetConfig(c)
	if err != nil {
		panic(err)
	}
}
