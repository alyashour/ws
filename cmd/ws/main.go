package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alyashour/ws/internal/config"
	"github.com/alyashour/ws/internal/syncer"
	"github.com/alyashour/ws/internal/tasks"
)

var ws = config.Ws{
	ConfPath:            "/tmp/.ws/",
	DefaultTaskFileName: "default.yaml",
}

func main() {
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
