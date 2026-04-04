package syncer

import (
	"github.com/go-git/go-git/v6"
)

func repoExists(path string) bool {
	// check that it's initialized
	_, err := git.PlainOpen(path)
	if err != nil {
		return false
	}

	return true
}

func statusShort(path string) (string, error) {
	// 1. Open the repository in the current directory
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	// 2. Get the working tree
	w, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	// 3. Get the status
	status, err := w.Status()
	if err != nil {
		return "", err
	}

	// 4. Print in "short" format
	if status.IsClean() {
		return "Working tree clean", nil
	} else {
		return status.String(), nil
	}
}
