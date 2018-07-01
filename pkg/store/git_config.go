package store

import (
	"gopkg.in/src-d/go-git.v4/config"
	fmtConfig "gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"gopkg.in/src-d/go-git.v4"
	"strings"
)

type gitConfigStore struct {
	repo *git.Repository
	config *config.Config
}
var _ ProfileStore = &gitConfigStore{}

func (s *gitConfigStore) SetProfile(name string, args []string) error {
	encArgs := ""
	if args != nil {
		encArgs = strings.Join(args, " ")
	}
	s.config.Raw = s.config.Raw.SetOption("profile", fmtConfig.NoSubsection, name, encArgs)
	return s.repo.Storer.SetConfig(s.config)
}



