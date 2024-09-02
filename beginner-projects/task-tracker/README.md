# Task Tracker
Sample solution for the [task tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh)

## Prerequisites
- Go should be installed on your system. If you haven't installed it yet, follow the instructions [here](https://golang.org/doc/install)

## How to Run
Clone the repository and run following command:
```bash
git clone https://github.com/David452/roadmapsh-backend.git
cd roadmapsh-backend/beginner-projects/task-tracker
```

Run the following commands to build and run the project:
```bash
go build -o task-tracker    # Build the project
./task-tracker --help       # To see the list of available commands
```

## Commands Overview

Below are the available comands for the `task-tracker` tool:

- `list`: List all tasks.
- `list todo`: List tasks with "todo" status.
- `list in-progress`: List tasks that are in progress.
- `list done`: List tasks that are marked as done.
- `add <task description>`: Add a new task with the specified description.
- `remove <task ID>`: Remove a task by its ID.
- `update <task ID> <task description>`: Update the description of a task by its ID.
- `mark-in-progress <task ID>`: Mark a task as in progress by its ID.
- `mark-done <task ID>`: Mark a task as done by its ID.

## Example Usage
Here are some examples of how to use the `task-tracker` tool:

### Listing Tasks
- **List all tasks**:
    ```bash 
    ./task-tracker list
    ```
- **List tasks with "To Do" status**:
    ```bash 
    ./task-tracker list todo
    ```
- **List tasks that are in progress**:
    ```bash 
    ./task-tracker list in-progress
    ```
- **List tasks that are marked as done**:
    ```bash 
    ./task-tracker list done
    ```
### Managing Tasks
- **Add a task**:
    ```bash
    ./task-tracker add "Do another project"
    ```
    This command adds a new task with the description "Do another project."
- **Remove a task**:
    ```bash
    ./task-tracker remove 1
    ```
    This command removes the task with ID 1.
- **Update a task**:
    ```bash
    ./task-tracker update 1 "Hello World"
    ```
    This command updates the description of the task with ID 1.
### Changing Task Status
- **Mark a task as in progress**:
    ```bash
    ./task-tracker mark-in-progress 1
    ```
    This command marks the task with ID 1 as in progress.
- **Mark a task as done**:
    ```bash
    ./task-tracker mark-done 1
    ```
    This command marks the task with ID 1 as done.

## Notes
- Ensure you are in the correct directory (`roadmapsh-backend/beginner-projects/task-tracker`) when running the commands
- Each task is identified by a unique ID that is generated when the task is added.