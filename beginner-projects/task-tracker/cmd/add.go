package cmd

import (
	"fmt"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)

func AddTask(cCtx *cli.Context) error {

	if cCtx.NArg() == 0 {
		return fmt.Errorf("missing task description")
	}

	err := task.AddTask(cCtx.Args().First())

	if err != nil {
		return err
	}

	return nil
}