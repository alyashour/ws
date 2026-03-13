package tasks

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var dataDir = defaultDataDir()

func defaultDataDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ws", "data")
}

func SetDataDir(path string) {
	dataDir = path
}

func taskFilePath() string {
	return filepath.Join(dataDir, "tasks", "tasks.yaml")
}

func save(tf TaskFile) error {
	// make if not already made
	if err := os.MkdirAll(filepath.Dir(taskFilePath()), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(tf)
	if err != nil {
		return err
	}
	return os.WriteFile(taskFilePath(), data, 0644)
}

func load() (TaskFile, error) {
	var tf TaskFile

	// read file
	data, err := os.ReadFile(taskFilePath())

	// 3 cases:
	// 1. file does not exist yet
	if errors.Is(err, fs.ErrNotExist) {
		// return an empty tf
		return tf, nil
	}

	// 2. error occurs
	if err != nil {
		// return error
		return tf, err
	}

	// 3. read was successful
	err = yaml.Unmarshal(data, &tf)
	return tf, nil // return parsed tf
}
