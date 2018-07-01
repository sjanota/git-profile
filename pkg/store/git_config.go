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

const (
	sectionName = "profile"
	argSeparator = " "
)

var _ ProfileStore = &gitConfigStore{}

func (s *gitConfigStore) AttachProfile(name string, args []string) error {
	encArgs := ""
	if args != nil {
		encArgs = strings.Join(args, argSeparator)
	}
	s.config.Raw.SetOption(sectionName, fmtConfig.NoSubsection, name, encArgs)
	return s.store()
}

func (s *gitConfigStore) DetachProfile(name string) error {
	s.config.Raw.Section(sectionName).RemoveOption(name)
	return s.store()
}

func (s *gitConfigStore) store() error {
	return s.repo.Storer.SetConfig(s.config)
}
