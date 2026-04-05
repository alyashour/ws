package syncer

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	workspace "github.com/alyashour/ws/internal/config"
	"github.com/alyashour/ws/internal/constants"
	"github.com/alyashour/ws/internal/io"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
	"github.com/zalando/go-keyring"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
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
func initSync(ws workspace.Ws) error {
	// Check if dir exists
	if exists, err := dirExists(ws.GetDataPath()); err != nil {
		return err
	} else if !exists {
		if io.YNPrompt("Config path dir does not exist. Create? (Y/N)") {
			// Create the directory
			if err := os.MkdirAll(ws.GetDataPath(), 0755); err != nil {
				log.Fatal(err)
			}
		} else {
			return errors.New("could not create dir")
		}
	}

	// Then create the repo
	fmt.Println("Creating repo...")
	repo, err := git.PlainInit(ws.GetDataPath(), false)
	if err != nil {
		return err
	}

	// Initialize the remote
	fmt.Println("Initializing Remote...")
	fmt.Println("Enter GitHub username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Enter Personal Access Token: ")
	byteToken, _ := term.ReadPassword(int(syscall.Stdin))
	token := string(byteToken)
	fmt.Println("\nValidating...")

	// Save to system keychain
	err = keyring.Set(constants.KeyRingService, username, token)
	if err != nil {
		return fmt.Errorf("Failed to save to keychain! %w", err)
	}

	// save the config to .ws
	err = saveCredConfig(ws.ConfPath, username)
	if err != nil {
		return fmt.Errorf("Failed to save config file. %w", err)
	}
	fmt.Println("Auth saved to keychain")

	remURL := io.StringPrompt("Please enter remote URL:")
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remURL},
	})
	fmt.Println("Hint: if you want to make manual changes to this repo it's accessible at", ws.GetDataPath())

	return err
}

func saveCredConfig(confPath string, username string) error {
	conf := credConf{Username: username}

	// Convert struct to YAML bytes
	data, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}

	// Determine save path (e.g., current directory or user home)
	filePath := filepath.Join(confPath, constants.SyncConfFileName)

	// Write file with permissions 0644 (rw-r--r--)
	return os.WriteFile(filePath, data, 0644)
}
