package cmd

import (
	"fmt"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)

func ListAllTasks(cCtx *cli.Context) error {
	
	tasks, err := task.GetTasks()

	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(task)
	}

	return nil
}