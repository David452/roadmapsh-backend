# Github User Activity
Sample solution for the [github user activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh)

## Prerequisites
- Go should be installed on your system. If you haven't installed it yet, follow the instructions [here](https://golang.org/doc/install)

## How to Run
Clone the repository and run following command:
```bash
git clone https://github.com/David452/roadmapsh-backend.git
cd roadmapsh-backend/beginner-projects/github-user-activity
```

Run the following commands to build and run the project:
```bash
go build -o github-user-activity    # Build the project
./github-user-activity <username>   # # Replace <username> with the GitHub username to fetch activity
```


## Example Usage
Here is a example of how to use the `github-user-activity` tool:

```bash
./github-user-activity jspisak
```

## Example Output

```
Events for jspisak
-  IssueCommentEvent in meta-llama/llama-models
-  PullRequestEvent in meta-llama/llama-models
-  Created branch in meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama-models
-  IssueCommentEvent in meta-llama/llama-toolchain
-  IssueCommentEvent in meta-llama/llama3
-  closed an issue in meta-llama/llama3
-  Pushed 2 commit(s) to meta-llama/llama3
-  PullRequestEvent in meta-llama/llama3
-  PullRequestReviewEvent in meta-llama/llama3
-  Pushed 1 commit(s) to meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama-models
-  Pushed 1 commit(s) to meta-llama/llama3
-  Pushed 1 commit(s) to meta-llama/llama
-  Pushed 1 commit(s) to meta-llama/llama3
-  IssueCommentEvent in meta-llama/llama
-  IssueCommentEvent in meta-llama/llama3
-  PullRequestEvent in meta-llama/PurpleLlama
-  Created branch in meta-llama/PurpleLlama
```

## Notes
- Ensure you are in the correct directory (`roadmapsh-backend/beginner-projects/github-user-activity`) when running the commands