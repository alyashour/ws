package syncer

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alyashour/ws/internal/io"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
)

func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		// path exists, is it a dir?
		return info.IsDir(), nil
	}

	// check if it exists but is not dir
	if errors.Is(err, os.ErrNotExist) {
		return false, nil // not exists - no error
	}

	// some other error occured
	return false, err
}

/*
Initializes the syncing git repo at path.
Prints an error on failure.
Returns whether the operation succeeded.
Notes:
- Prompts to create dir if does not exist
- Fails if git repo already exists there
*/
func initRepo(path string) error {
	// Check if dir exists
	if exists, err := dirExists(path); err != nil {
		return err
	} else if !exists {
		fmt.Print("")
		if io.YNPrompt("Config path dir does not exist. Create? (Y/N)") {
			// Create the directory
			if err := os.MkdirAll(path, 0755); err != nil {
				log.Fatal(err)
			}
		} else {
			return errors.New("could not create dir")
		}
	}

	// Then create the repo
	fmt.Println("Creating repo...")
	repo, err := git.PlainInit(path, false)
	if err != nil {
		return err
	}

	// Initialize the remote
	fmt.Println("Initializing Remote...")
	fmt.Println("Create a new remote repository now (GitHub, GitLab, self-hosted, etc.).")
	remURL := io.StringPrompt("Please enter remote URL:")
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remURL},
	})
	fmt.Println("Hint: if you want to make manual changes to this repo it's accessible at", path)

	return err
}
