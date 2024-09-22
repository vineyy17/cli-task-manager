# CLI Task Manager

A simple and efficient command-line interface (CLI) task manager built with Go. This application allows you to manage your tasks directly from the terminal, providing a streamlined workflow for task creation, listing, completion, and deletion.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Features

- Add new tasks
- List all tasks (including option to show completed tasks)
- Mark tasks as complete
- Delete tasks
- Persistent storage using BoltDB
- Human-readable time differences for task creation dates

  ## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/vineyy17/cli-task-manager.git
    ```

2. **Navigate to the project directory**:
    ```bash
    cd cli-task-manager
    ```

2. **Build the application**:
    Make sure you have Go installed on your machine. Then, build the project using the following command:
    ```bash
    go build
    ```

3. **(Optional) Move the binary to a directory in your PATH for easy access:**:
    ```bash
    sudo mv cli-task-manager /usr/local/bin/task
    ```

## Usage

After installation, you can use the `task` command followed by subcommands to manage your tasks.

## Commands

- `task add <task description>`: Add a new task
- `task list`: List all uncompleted tasks
- `task list --all` or `task list -a`: List all tasks, including completed ones
- `task complete <task_id>`: Mark a task as complete
- `task delete <task_id>`: Delete a task

### Examples

```bash
# Add a new task
task add "Go for a walk"

# List all uncompleted tasks
task list

# List all tasks, including completed ones
task list --all

# Mark task 1 as complete
task complete 1

# Delete task 2
task delete 2
```

## Dependencies

- [BoltDB](https://github.com/boltdb/bolt): An embedded key/value database for Go.
- [Cobra](https://github.com/spf13/cobra): A library for creating powerful CLI applications.
- [timediff](https://github.com/mergestat/timediff): A Go library for formatting time differences in a human-readable way.
- [homedir](https://github.com/mitchellh/go-homedir): A Go library for expanding the user's home directory.

## Contributing

Contributions to the CLI Task Manager are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.
