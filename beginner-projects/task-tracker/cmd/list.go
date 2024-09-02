package cmd

import (
	"github.com/David452/task-tracer/internal/task"

	"github.com/urfave/cli/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/andanhm/go-prettytime"
)

func ListTasks(cCtx *cli.Context) error {
	
	tasks, err := task.GetTasks()

	if err != nil {
		return err
	}

	prettyTable := table.NewWriter()
	prettyTable.SetOutputMirror(cCtx.App.Writer)
	prettyTable.AppendHeader(table.Row{"ID", "Description", "Status", "Created", "Updated"})


	for _, t := range tasks {
		switch cCtx.Args().First() {
		case "todo":
			if t.Status == task.StatusTodo {
				prettyTable.AppendRow(table.Row{t.ID, t.Description, t.Status, prettytime.Format(t.CreatedAt), prettytime.Format(t.UpdatedAt)})
			}
		case "in-progress":
			if t.Status == task.StatusInProgress {
				prettyTable.AppendRow(table.Row{t.ID, t.Description, t.Status, prettytime.Format(t.CreatedAt), prettytime.Format(t.UpdatedAt)})
			}
		case "done":
			if t.Status == task.StatusDone {
				prettyTable.AppendRow(table.Row{t.ID, t.Description, t.Status, prettytime.Format(t.CreatedAt), prettytime.Format(t.UpdatedAt)})
			}
		default:
			prettyTable.AppendRow(table.Row{t.ID, t.Description, t.Status, prettytime.Format(t.CreatedAt), prettytime.Format(t.UpdatedAt)})
		}
	}

	prettyTable.Render()

	return nil
}