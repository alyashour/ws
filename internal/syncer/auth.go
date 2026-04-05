package syncer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alyashour/ws/internal/constants"
	"github.com/zalando/go-keyring"
	"gopkg.in/yaml.v3"
)

type credConf struct {
	Username string `yaml:"username"`
}

// Returns username, PAT, error
func getCredentials(configPath string) (string, string, error) {
	// join the config path to get the sync.yaml file
	// TODO: move all this config loading into a single struct, load it, then just access it in this and other
	// config-dependant functions
	configPath = filepath.Join(configPath, constants.SyncConfFileName)
	// 1. Read the YAML file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", "", fmt.Errorf("config not found, run 'ws sync init' first")
	}

	var conf credConf
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return "", "", err
	}

	// 2. Fetch the token from Keychain using the username from YAML
	token, err := keyring.Get(constants.KeyRingService, conf.Username)
	if err != nil {
		return "", "", fmt.Errorf("could not find token for %s in keychain", conf.Username)
	}

	return conf.Username, token, nil
}
