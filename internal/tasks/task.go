package tasks

import (
	"fmt"
	"strings"
)

type Task struct {
	Id        ID     `yaml:"id"`
	Text      string `yaml:"text"`
	Done      bool   `yaml:"done"`
	CreatedAt string `yaml:"created_at"`
	Due       string `yaml:"due,omitempty"`
}

func (t Task) String() string {
	return fmt.Sprintf("{ %s: %s }", FormatID(t.Id, "t"), t.Text)
}

func FormatTasks(tasks []Task) string {
	var sb strings.Builder
	for _, t := range tasks {
		check := " "
		if t.Done {
			check = "x"
		}
		fmt.Fprintf(&sb, "%s. [%s] %s\t(%s)\n", FormatID(t.Id, "t"), check, t.Text, t.CreatedAt)
	}
	return sb.String()
}
