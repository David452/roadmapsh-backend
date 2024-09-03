package cmd

import (
	"fmt"

	"github.com/David452/github-user-activity/internal/githubapi"
)

func List(user string,client *githubapi.Client) error {

	activity, err := client.GetEvents(user)

	if err != nil {
		return err
	}

	fmt.Println(activity)
	
	return nil
}