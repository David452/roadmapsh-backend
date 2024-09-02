package cmd

import (
	"fmt"
	"strconv"

	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
)



func MarkDone(cCtx *cli.Context) error {

	if cCtx.NArg() == 0 {
		return fmt.Errorf("missing task ID")
	}

	idFromArgs, err := strconv.Atoi(cCtx.Args().First())

	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	fmt.Fprintf(cCtx.App.Writer, "Marked task done: %s\n", cCtx.Args().First())

	return task.MarkTask(idFromArgs, task.StatusDone)
}


func MarkInProgress(cCtx *cli.Context) error {

	if cCtx.NArg() == 0 {
		return fmt.Errorf("missing task ID")
	}

	idFromArgs, err := strconv.Atoi(cCtx.Args().First())

	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	fmt.Fprintf(cCtx.App.Writer, "Marked task in progress: %s\n", cCtx.Args().First())

	return task.MarkTask(idFromArgs, task.StatusInProgress)

}