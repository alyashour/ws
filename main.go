package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alyashour/ws/internal/config"
	"github.com/alyashour/ws/internal/syncer"
	"github.com/alyashour/ws/internal/tasks"
)

func main() {
	// init ws
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to initialize ws. UserHomeDir could not be determined. Please submit a bug report with your OS.")
		return
	}

	var ws = config.Ws{
		ConfPath:            filepath.Join(homedir, ".ws"),
		DefaultTaskFileName: "default.yaml",
	}

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "todo", "tasks":
		tasks.Run(ws, os.Args[2:])
	case "sync", "syncer":
		syncer.Run(ws, os.Args[2:])
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: ws <cmd> [options]")
	fmt.Println(`
Commands:
- todo/tasks    Task management
- sync          Syncing`)
}
