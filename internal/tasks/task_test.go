package tasks

import (
	"testing"
)

func TestAdd(t *testing.T) {
	SetDataDir(t.TempDir())

	// Add to TaskFile
	task, err := Add("fix the bug")
	if err != nil {
		t.Fatal(err)
	}
	if task.Text != "fix the bug" {
		t.Errorf("got %q, want %q", task.Text, "fix the bug")
	}

	// Load from TaskFile and ensure it's identical
	tf, err := load()
	idx, loadedTask := tf.find(task.Id)
	if idx == -1 {
		if loadedTask != task {
			t.Errorf("got %+v, want %+v", loadedTask, task)
		}
	}
}

func TestList(t *testing.T) {
	SetDataDir(t.TempDir())

	// list empty
	tasks, err := List()
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected empty list, got %d tasks", len(tasks))
	}

	// add some tasks then list
	Add("fix the bug")
	Add("write tests")

	tasks, err = List()
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestDone(t *testing.T) {
	SetDataDir(t.TempDir())

	task, err := Add("fix the bug")
	if err != nil {
		t.Fatal(err)
	}

	done, err := Done(task.Id)
	if err != nil {
		t.Fatal(err)
	}
	if !done.Done {
		t.Errorf("expected done=true, got false")
	}

	// verify persisted
	tf, _ := load()
	_, loaded := tf.find(task.Id)
	if !loaded.Done {
		t.Errorf("done not persisted to file")
	}
}

func TestRemove(t *testing.T) {
	SetDataDir(t.TempDir())

	task, err := Add("fix the bug")
	if err != nil {
		t.Fatal(err)
	}

	_, err = Remove(task.Id)
	if err != nil {
		t.Fatal(err)
	}

	// verify removed from file
	tf, _ := load()
	idx, _ := tf.find(task.Id)
	if idx != -1 {
		t.Errorf("expected task to be removed, still found at index %d", idx)
	}
}

func TestEdit(t *testing.T) {
	SetDataDir(t.TempDir())

	task, err := Add("fix the bug")
	if err != nil {
		t.Fatal(err)
	}

	updated, err := Edit(task.Id, "fix the other bug")
	if err != nil {
		t.Fatal(err)
	}
	if updated.Text != "fix the other bug" {
		t.Errorf("got %q, want %q", updated.Text, "fix the other bug")
	}

	// verify persisted
	tf, _ := load()
	_, loaded := tf.find(task.Id)
	if loaded.Text != "fix the other bug" {
		t.Errorf("edit not persisted to file")
	}
}
