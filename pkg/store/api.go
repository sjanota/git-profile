package store

import (
	"gopkg.in/src-d/go-git.v4"
	"git-profile/pkg/repo"
)

type ProfileStore interface {
	AttachProfile(name string, args []string) error
	DetachProfile(name string) error
}

func NewStore(r *git.Repository) (ProfileStore, error) {
	c, err := r.Config()
	if err != nil {
		return nil, err
	}

	return &gitConfigStore{repo: r, config: c}, nil
}

func NewStoreForCurrentRepo() (ProfileStore, error) {
	r, err := repo.OpenCurrent()
	if err != nil {
		panic(err)
	}

	return NewStore(r)
}
