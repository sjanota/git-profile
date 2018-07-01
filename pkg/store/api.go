package store

import "gopkg.in/src-d/go-git.v4"

type ProfileStore interface {
	SetProfile(name string, args []string) error
}

func NewStore(r *git.Repository) (ProfileStore, error) {
	c, err := r.Config()
	if err != nil {
		return nil, err
	}

	return &gitConfigStore{repo: r, config: c}, nil
}
