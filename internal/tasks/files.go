package tasks

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func save(tf TaskFile, path string) error {
	// make if not already made
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(tf)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Loads a taskfile from memory (yaml) file
func load(path string) (TaskFile, error) {
	var tf TaskFile

	// read file
	data, err := os.ReadFile(path)

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
