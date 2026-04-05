package syncer

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func pull(repoPath string, cfgPath string) error {
	// check that it's initialized
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	// Get the working directory (worktree)
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Load auth
	username, token, err := getCredentials(cfgPath)
	if err != nil {
		return fmt.Errorf("Failed to get credentials. %w", err)
	}
	credentials := &http.BasicAuth{
		Username: username,
		Password: token,
	}

	// Pull
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth:       credentials,
		Progress:   os.Stdout,
	})
	if err == git.NoErrAlreadyUpToDate {
		fmt.Println(err.Error())
		return nil
	}

	return err
}
