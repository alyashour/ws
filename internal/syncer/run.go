package syncer

import (
	"fmt"
	"strings"

	"github.com/alyashour/ws/internal/config"
)

func Run(cfg config.Ws, args []string) {
	if len(args) == 0 {
		usage()
		return
	}

	switch strings.ToLower(args[0]) {
	case "init":
		if err := initSync(cfg); err != nil {
			fmt.Println("Err:", err)
		}
	case "clone":
		fmt.Println("Not yet implemented")
	case "pull":
		if err := pull(cfg.GetDataPath(), cfg.ConfPath); err != nil {
			fmt.Println("Failed to pull from remote. Err:", err)
		}
	case "push":
		if err := push(cfg.GetDataPath(), cfg.ConfPath); err != nil {
			fmt.Println("Failed to push to remote. Err:", err)
		}
	case "status":
		status(cfg.GetDataPath())
	default:
		usage()
	}
}

// Does 4 things:
// 1. Print data path
// 2. Checks if the data path dir exists
// 3. Checks if the repository exists at that dir
// 4. Runs `git status`
func status(datapath string) {
	fmt.Println("|| Workspace Data Syncer Status")
	// 1. Print data path
	fmt.Println("Datapath:", datapath)

	// 2. Check if dir exists
	exists, err := dirExists(datapath)
	if !exists {
		fmt.Println("Dir does not exist.")
		return
	}
	if err != nil {
		fmt.Println("Error opening data dir:\n", err)
		return
	}

	// 3. Check if repo exists
	exists = repoExists(datapath)
	if exists {
		fmt.Println("Repo exists")
	} else {
		fmt.Println("Repo does not exist")
		return
	}

	// 4. Run git status
	statusStr, err := statusShort(datapath)
	if err != nil {
		fmt.Println("Failed to get status. Err:", err)
		return
	}

	fmt.Println("\n|| Git Status")
	fmt.Println(statusStr)
}

func usage() {
	fmt.Println("Usage: ws sync <verb>")
	fmt.Println(`
Verbs:
- init		Initializes a new repository
- clone		Clones an existing repository
- pull		Pull from remote
- push		Push to remote
- status    Get status`)
}
