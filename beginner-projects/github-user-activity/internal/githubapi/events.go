package githubapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func (c *Client) GetEvents(user string) (*Activity, error) {
	
	url := fmt.Sprintf("%s/users/%s/events", c.baseURL, user)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &Activity{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &Activity{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 304 {

		return &Activity{}, fmt.Errorf("not modified")
	}

	if resp.StatusCode == 403 {

		return &Activity{}, fmt.Errorf("forbidden")
	}

	if resp.StatusCode == 503 {

		return &Activity{}, fmt.Errorf("service unavailable")
	}

	if resp.StatusCode != 200 {
		return &Activity{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}


	var activity Activity
	err = json.NewDecoder(resp.Body).Decode(&activity)
	if err != nil {
		return &Activity{}, fmt.Errorf("error decoding JSON response: %w", err)
	}


	return &activity, nil
}