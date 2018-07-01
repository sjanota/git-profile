package repo

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func OpenCurrent() (*git.Repository, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	r, err := git.PlainOpen(cwd)
	if err != nil {
		return nil, err
	}

	return r, nil
}
