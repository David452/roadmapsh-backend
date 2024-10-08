package githubapi

type Activity struct {
	Type      string `json:"type"`
	Repo      struct {
		Name string `json:"name"`
	}	`json:"repo"`

	CreatedAt string `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}