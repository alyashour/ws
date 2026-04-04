package syncer

import (
	"github.com/go-git/go-git/v6"
)

func pull(path string) error {
	// check that it's initialized
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	// Get the working directory (worktree)
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Pull
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	return err
}
