package tasks

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/alyashour/ws/internal/config"
)

var TaskNotFoundErr = errors.New("task not found")

// Function responsible for running todo commands
// Parses arguments and performs action
func Run(cfg config.Ws, args []string) {
	// if no verb was given
	if len(args) == 0 {
		usage()
		return
	}

	switch args[0] {
	case "add", "new": // USAGE: add "text"
		if len(args) != 2 {
			fmt.Println("Usage: ws add/new <text>")
			return
		}

		// parse args
		text := args[1]

		// add
		task, err := Add(text, filepath.Join(cfg.GetDefaultTaskFilePath()))

		// print
		if err == nil {
			fmt.Printf("Added task '%+v'\n", task)
		} else {
			fmt.Printf("Err: %s\n", err)
		}
	case "list", "ls":
		if len(args) != 1 {
			fmt.Println("Usage: ws list")
			return
		}

		// get list
		tasks, err := List(cfg.GetDefaultTaskFilePath())
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		// print
		if len(tasks) == 0 {
			fmt.Println("No tasks yet.")
			return
		}

		fmt.Printf("Tasks (%d):\n%s\n", len(tasks), FormatTasks(tasks))
	case "done": // usage: done <id>
		if len(args) != 2 {
			fmt.Println("Usage: ws done <id>")
			return
		}

		// parse args
		id := normalizeID(args[1])

		// check as done
		task, err := Done(cfg.GetDefaultTaskFilePath(), id)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		fmt.Printf("Marked %s as done.\n", task)
	case "edit": // usage: edit <id> <text>
		if len(args) != 3 {
			fmt.Println("Usage: ws edit <id> <text>")
			return
		}

		// parse args
		id := normalizeID(args[1])
		text := args[2]

		task, err := Edit(cfg.GetDefaultTaskFilePath(), id, text)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		fmt.Printf("Edited task: %s", task)
	case "remove", "rm":
		if len(args) != 2 {
			fmt.Println("Usage: ws remove <id>")
			return
		}

		// parse args
		id := normalizeID(args[1])

		task, err := Remove(cfg.GetDefaultTaskFilePath(), id)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return
		}

		fmt.Printf("Removed task %s\n", task)
	default:
		fmt.Println("Uknown verb:", args[0])
		usage()
	}
}

// Load tasks from file, add new, and save to file
// Needs the path of the taskfile (.yaml) file
func Add(text string, taskFilePath string) (Task, error) {
	// first load the tasks from taskfile
	tf, err := load(taskFilePath)
	if err != nil {
		return Task{}, err
	}

	// then generate the next sequential id
	nextId := tf.getNextID()

	// use the id to make a new task
	task := Task{
		Id:        nextId,
		Text:      text,
		CreatedAt: time.Now().Format("03:04 PM 02-01-2006"),
	}

	// add it to the taskfile
	tf.Tasks = append(tf.Tasks, task)

	// save the taskfile
	return task, save(tf, taskFilePath)
}

func List(taskFilePath string) ([]Task, error) {
	// load the tf
	tf, err := load(taskFilePath)
	if err != nil {
		return nil, err
	}

	// return
	return tf.Tasks, nil
}

func Done(taskFilePath string, id ID) (Task, error) {
	// load the tf
	tf, err := load(taskFilePath)
	if err != nil {
		return Task{}, err
	}

	// get the task
	i, _ := tf.find(id)
	if i == -1 {
		// task not found
		return Task{}, TaskNotFoundErr
	}

	// task found

	// modify it
	tf.Tasks[i].Done = true

	// save it
	return tf.Tasks[i], save(tf, taskFilePath)
}

// Only supports changing the text as of now
func Edit(taskFilePath string, id ID, text string) (Task, error) {
	// load the tf
	tf, err := load(taskFilePath)
	if err != nil {
		return Task{}, err
	}

	// find the task to be edited
	i, _ := tf.find(id)
	if i == -1 {
		// not found
		return Task{}, TaskNotFoundErr
	}

	// found

	// modify
	tf.Tasks[i].Text = text

	// save
	return tf.Tasks[i], save(tf, taskFilePath)
}

func Remove(taskFilePath string, id ID) (Task, error) {
	// load the tf
	tf, err := load(taskFilePath)
	if err != nil {
		return Task{}, err
	}

	// get the task
	for i, t := range tf.Tasks {
		if t.Id == id {
			// found, remove from tf
			tf.Tasks = append(tf.Tasks[:i], tf.Tasks[i+1:]...)

			// save tf
			return t, save(tf, taskFilePath)
		}
	}

	// task not found
	return Task{}, TaskNotFoundErr
}

func usage() {
	fmt.Println("Usage: ws todo <verb>")
	fmt.Println(`
Verbs:
- add/new
- list/ls
- done
- edit
- remove/rm`)
}
