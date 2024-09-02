package cmd

import (
	"fmt"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)

func ListTasks(cCtx *cli.Context) error {
	
	tasks, err := task.GetTasks()

	if err != nil {
		return err
	}

	for _, t := range tasks {
		switch cCtx.Args().First() {
		case "todo":
			if t.Status == task.StatusTodo {
				fmt.Println(t)
			}
		case "in-progress":
			if t.Status == task.StatusInProgress {
				fmt.Println(t)
			}
		case "done":
			if t.Status == task.StatusDone {
				fmt.Println(t)
			}
		default:
			fmt.Println(t)
		}
	}

	return nil
}