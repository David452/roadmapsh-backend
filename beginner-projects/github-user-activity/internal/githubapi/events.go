package githubapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func (c *Client) GetEvents(user string) ([]Activity, error) {
	
	url := fmt.Sprintf("%s/users/%s/events", c.baseURL, user)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 304 {

		return nil, fmt.Errorf("not modified")
	}

	if resp.StatusCode == 403 {

		return nil, fmt.Errorf("forbidden")
	}

	if resp.StatusCode == 503 {

		return nil, fmt.Errorf("service unavailable")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}


	var activities []Activity
	err = json.NewDecoder(resp.Body).Decode(&activities)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}


	return activities, nil
}