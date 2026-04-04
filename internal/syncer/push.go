package syncer

import (
	"fmt"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func push(path string, privateKeyFile string) error {
	// check that it's initialized
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("Failed to open git repo. Check path or run `ws syncer init`.\nPath: %s\nErr:%w", path, err)
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
