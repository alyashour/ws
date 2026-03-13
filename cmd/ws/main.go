package main

import (
	"fmt"
	"os"

	"github.com/alyashour/ws/internal/tasks"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "todo", "tasks":
		tasks.Run(os.Args[2:])
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: ws <cmd> [options]")
	fmt.Println("Commands:\n- todo/tasks")
}
