# Task CLI

A simple command-line task management application written in Go. This tool allows you to manage your to-do list directly from the terminal with persistent JSON storage.

## Features

- ✅ Add new tasks
- 📝 Update task descriptions
- 🔄 Change task status
- 🗑️ Delete tasks
- 📋 List all tasks with formatted table output
- 💾 Persistent JSON storage
- ⏰ Automatic timestamp tracking (created and updated)

## Installation

1. Clone this repository
2. Navigate to the task-cli directory:
```bash
cd task-cli
```

3. Build the application:
```bash
go build
```

## Usage

### Add a Task
```bash
./task-cli -add "Buy groceries"
```

### List All Tasks
```bash
./task-cli -list
```

### Update a Task Description
Update task by ID with new description (format: `id:new_description`):
```bash
./task-cli -update "0:Buy groceries and milk"
```

### Change Task Status
Change task status by ID (format: `id:status`):
```bash
./task-cli -status "0:done"
./task-cli -status "1:in-progress"
```

### Delete a Task
Delete task by ID:
```bash
./task-cli -delete 0
```

## Data Structure

Tasks are stored with the following properties:
- **ID**: Unique identifier (auto-incremented)
- **Description**: Task description
- **Status**: Current status (todo, in-progress, done, etc.)
- **CreatedAt**: Timestamp when task was created
- **UpdatedAt**: Timestamp when task was last modified

## Storage

Tasks are persisted to `first-todos.json` in JSON format. The file is automatically created and updated with each operation.

## Dependencies

- [aquasecurity/table](https://github.com/aquasecurity/table) - For formatted table output

Install dependencies:
```bash
go mod download
```

## Project Structure

```
task-cli/
├── main.go          # Application entry point
├── todo.go          # Todo struct and operations (add, delete, update, print)
├── command.go       # Command-line flag handling and execution
├── storage.go       # Generic JSON storage implementation
├── *_test.go        # Unit tests
├── go.mod           # Go module file
└── README.md        # This file
```

## Development

### Running Tests
```bash
go test -v
```

### Code Overview

- **Todo**: Represents a single task with ID, description, status, and timestamps
- **Todos**: Collection of Todo items with methods for CRUD operations
- **Storage**: Generic storage implementation for saving/loading data to JSON
- **Command**: Command-line interface for user interactions

## Example Workflow

```bash
# Add some tasks
./task-cli -add "Learn Go"
./task-cli -add "Build CLI app"
./task-cli -add "Write tests"

# List all tasks
./task-cli -list

# Update a task status
./task-cli -status "0:done"

# Update task description
./task-cli -update "1:Build awesome CLI app"

# Delete a task
./task-cli -delete 2

# View final list
./task-cli -list
```

## License

This is a learning project for practicing Go programming.
