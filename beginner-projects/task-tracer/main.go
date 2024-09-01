package main

import (
	"log"
	"os"

	"github.com/David452/task-tracer/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "list",
				Description: "List all tasks",
				Action: cmd.ListAllTasks,

			},
			{
				Name: "add",
				Description: "Add a new task",
				Action: cmd.AddTask,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}