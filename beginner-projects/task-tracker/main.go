package main

import (
	"log"
	"os"

	"github.com/David452/task-tracer/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	
	app := &cli.App{
		Name: "task-tracer",
		Usage: "A simple task manager",
		Commands: []*cli.Command{
			{
				Name: "list",
				Usage: "List all tasks",
				Action: cmd.ListTasks,
				ArgsUsage: "[todo|in-progress|done]",

			},
			{
				Name: "add",
				Usage: "Add a new task",
				Action: cmd.AddTask,
				ArgsUsage: "<description>",
			},
			{
				Name: "remove",
				Usage: "Remove a task by ID",
				Action: cmd.RemoveTask,
				ArgsUsage: "<ID>",
			},
			{
				Name: "update",
				Usage: "Update a task by ID",
				Action: cmd.UpdateTask,
				ArgsUsage: "<ID> <description>",
			},
			{
				Name: "mark-in-progress",
				Usage: "Mark a task as in progress",
				Action: cmd.MarkInProgress,
				ArgsUsage: "<ID>",
			},
			{
				Name: "mark-done",
				Usage: "Mark a task as done",
				Action: cmd.MarkDone,
				ArgsUsage: "<ID>",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}