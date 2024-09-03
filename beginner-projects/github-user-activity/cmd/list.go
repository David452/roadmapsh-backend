package cmd

import (
	"fmt"

	"github.com/David452/github-user-activity/internal/githubapi"
)

func List(user string,client *githubapi.Client) error {

	activities, err := client.GetEvents(user)

	if err != nil {
		return err
	}

	if len(activities) == 0 {

		fmt.Println("No events found")
		return nil
	}

	fmt.Printf("Events for %s\n", user)

	for _, activity := range activities {
		var action string
		switch activity.Type {
		case "PushEvent":
			commitCount := len(activity.Payload.Commits)
			action = fmt.Sprintf("Pushed %d commit(s) to %s", commitCount, activity.Repo.Name)
		case "IssuesEvent":
			action = fmt.Sprintf("%s an issue in %s", activity.Payload.Action, activity.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", activity.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", activity.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", activity.Payload.RefType, activity.Repo.Name)
		default:
			action = fmt.Sprintf("%s in %s", activity.Type, activity.Repo.Name)
		}

		fmt.Println("- ", action)
	}
	
	return nil
}