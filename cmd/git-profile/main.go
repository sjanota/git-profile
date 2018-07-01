package main

import (
	"os"
	"errors"
	"git-profile/pkg/repo"
	"git-profile/pkg/store"
)

type SetProfile struct {
	name string
	args []string
}

func main() {
	args, err := parseSetProfile()
	if err != nil {
		panic(err)
	}

	r, err := repo.OpenCurrent()
	if err != nil {
		panic(err)
	}

	s, err := store.NewStore(r)
	if err != nil {
		panic(err)
	}

	err = s.SetProfile(args.name, args.args)
	if err != nil {
		panic(err)
	}
}

func parseSetProfile() (*SetProfile, error){
	args := SetProfile{}
	if len(os.Args) < 2 {
		return nil, errors.New("please provide profile name")
	}
	args.name = os.Args[1]

	if len(os.Args) > 2 {
		args.args = os.Args[2:]
	}

	return &args, nil
}
