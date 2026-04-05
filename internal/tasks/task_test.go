package tasks

import (
	"testing"

	"github.com/alyashour/ws/internal/config"
)

func TestAdd(t *testing.T) {
	var ws = config.Ws{
		ConfPath: t.TempDir(),
	}

	// Add to TaskFile
	task, err := Add("fix the bug", ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}
	if task.Text != "fix the bug" {
		t.Errorf("got %q, want %q", task.Text, "fix the bug")
	}

	// Load from TaskFile and ensure it's identical
	tf, err := load(ws.GetDefaultTaskFilePath())
	idx, loadedTask := tf.find(task.Id)
	if idx == -1 {
		if loadedTask != task {
			t.Errorf("got %+v, want %+v", loadedTask, task)
		}
	}
}

func TestList(t *testing.T) {
	var ws = config.Ws{
		ConfPath: t.TempDir(),
	}

	// list empty
	tasks, err := List(ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected empty list, got %d tasks", len(tasks))
	}

	// add some tasks then list
	Add("fix the bug", ws.GetDefaultTaskFilePath())
	Add("write tests", ws.GetDefaultTaskFilePath())

	tasks, err = List(ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestDone(t *testing.T) {
	var ws = config.Ws{
		ConfPath: t.TempDir(),
	}

	task, err := Add("fix the bug", ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}

	done, err := Done(ws.GetDefaultTaskFilePath(), task.Id)
	if err != nil {
		t.Fatal(err)
	}
	if !done.Done {
		t.Errorf("expected done=true, got false")
	}

	// verify persisted
	tf, _ := load(ws.GetDefaultTaskFilePath())
	_, loaded := tf.find(task.Id)
	if !loaded.Done {
		t.Errorf("done not persisted to file")
	}
}

func TestRemove(t *testing.T) {
	var ws = config.Ws{
		ConfPath: t.TempDir(),
	}

	task, err := Add("fix the bug", ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}

	_, err = Remove(ws.GetDefaultTaskFilePath(), task.Id)
	if err != nil {
		t.Fatal(err)
	}

	// verify removed from file
	tf, _ := load(ws.GetDefaultTaskFilePath())
	idx, _ := tf.find(task.Id)
	if idx != -1 {
		t.Errorf("expected task to be removed, still found at index %d", idx)
	}
}

func TestEdit(t *testing.T) {
	var ws = config.Ws{
		ConfPath: t.TempDir(),
	}

	task, err := Add("fix the bug", ws.GetDefaultTaskFilePath())
	if err != nil {
		t.Fatal(err)
	}

	updated, err := Edit(ws.GetDefaultTaskFilePath(), task.Id, "fix the other bug")
	if err != nil {
		t.Fatal(err)
	}
	if updated.Text != "fix the other bug" {
		t.Errorf("got %q, want %q", updated.Text, "fix the other bug")
	}

	// verify persisted
	tf, _ := load(ws.GetDefaultTaskFilePath())
	_, loaded := tf.find(task.Id)
	if loaded.Text != "fix the other bug" {
		t.Errorf("edit not persisted to file")
	}
}
