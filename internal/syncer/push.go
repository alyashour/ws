package syncer

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func push(repoPath string, confPath string) error {
	// check that it's initialized
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("Failed to open git repo. Check path or run `ws syncer init`.\nPath: %s\nErr:%w", repoPath, err)
	}

	// get worktree
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// stage
	err = worktree.AddWithOptions(&git.AddOptions{All: true})
	if err != nil {
		return err
	}

	// commit
	fmt.Println("Commiting...")
	_, err = worktree.Commit("Automatic sync commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "WS Syncer",
			Email: "syncer@ws.com",
			When:  time.Now(),
		},
	})
	if err == git.ErrEmptyCommit {
		fmt.Println("Clean working tree... skipping new commit.")
	} else if err != nil {
		return err
	}

	// Get the credentials
	username, token, err := getCredentials(confPath)
	if err != nil {
		return fmt.Errorf("Failed to get credentials. Err: %w", err)
	}

	auth := &http.BasicAuth{
		Username: username,
		Password: token,
	}

	// Push
	repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
		Progress:   os.Stdout,
	})

	return err
}
