package cmd

import (
	"fmt"
	"strconv"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)

func UpdateTask(cCtx *cli.Context) error {
	
	if cCtx.NArg() < 2 {
		return fmt.Errorf("missing task ID or description")
	}

	idFromArgs, err := strconv.Atoi(cCtx.Args().Get(0))

	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	err = task.UpdateTask(idFromArgs, cCtx.Args().Get(1))

	return err


}