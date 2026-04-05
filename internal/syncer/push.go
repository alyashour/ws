package syncer

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func push(path string, privateKeyFile string) error {
	// check that it's initialized
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("Failed to open git repo. Check path or run `ws syncer init`.\nPath: %s\nErr:%w", path, err)
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

	// Get auth pubkey
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
	if err != nil {
		return err
	}

	// Push
	repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       publicKeys,
	})

	return err
}
