package main

import (
	"fmt"
	"os"
	"log"

	"github.com/David452/github-user-activity/internal/githubapi"
	"github.com/David452/github-user-activity/cmd"

	"github.com/urfave/cli/v2"
)

func main() {

	client := githubapi.NewClient("https://api.github.com")
	
	app := &cli.App{
		Name: "github-user-activity",
		Usage: "Get user activity",
		Action: func (cCtx *cli.Context) error {
			if cCtx.NArg() == 0 {
				return fmt.Errorf("missing user name")
			}

			return cmd.List(cCtx.Args().First(), client)
		},
		ArgsUsage: "[user]",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}