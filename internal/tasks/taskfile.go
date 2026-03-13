package tasks

type TaskFile struct {
	Tasks []Task `yaml:"tasks"`
}

func (tf TaskFile) getNextID() ID {
	// map taskfile.Tasks to []ID
	ids := make([]ID, len(tf.Tasks))
	for i, t := range tf.Tasks {
		ids[i] = t.Id
	}

	// use getNextID
	return genNextID(ids)
}

// Returns id to task and copy.
// Returns -1 if not found.
func (tf TaskFile) find(id ID) (int, Task) {
	// get the task
	for i, t := range tf.Tasks {
		if t.Id == id {
			return i, t
		}
	}

	return -1, Task{}
}

func (tf TaskFile) String() string {
	return FormatTasks(tf.Tasks)
}
