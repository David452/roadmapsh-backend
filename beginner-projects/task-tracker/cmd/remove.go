package cmd

import (
	"fmt"
	"strconv"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)

func RemoveTask(cCtx *cli.Context) error {

	if cCtx.NArg() == 0 {
		return fmt.Errorf("missing task ID")
	}

	idFromArgs, err := strconv.Atoi(cCtx.Args().First())

	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	err = task.RemoveTask(idFromArgs)

	return err
}