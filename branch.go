package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type RefList struct {
	//	Refs contains all of the specified repository's references
	Refs []plumbing.ReferenceName
}

// FillList fills the RefList's Refs field with all local and remotely
// tracked branches.
//
// The expected string should be a path to a directory where a git
// repository lives. If the directory is not a git repository then an
// error will be returned.
func (r *RefList) FillList(s string) (*RefList, error) {
	repo, err := git.PlainOpen(s)
	if err != nil {
		return nil, err
	}

	refs, err := repo.References()
	if err != nil {
		return nil, err
	}

	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && !ref.Name().IsTag() {
			r.Refs = append(r.Refs, ref.Name())
		}
		return nil
	})
	return r, nil
}
