# Pomo CLI

## The Best Pomodoro CLI App Ever!

Pomodoro CLI is a command-line application designed to help you manage your tasks efficiently using the Pomodoro technique. It also includes task management functionalities, allowing you to add, complete, list, remove, and search for tasks.

---

## Installation

To install Pomodoro CLI, follow these steps:

```sh
# Clone the repository
git clone <repository-url>

# Navigate to the project directory
cd pomo-cli

# Install dependencies (if applicable)
go build
```

---

## Usage

Pomodoro CLI provides several commands to help you manage your tasks effectively.

### Add a Task
**Adds a new task or multiple ones.**

```sh
./pomo-cli add <task>
```

#### Examples:
```sh
# Add a single task
./pomo-cli add Pipoca

# Add multiple tasks
./pomo-cli add Pipoca Maluca

# Add a task with multiple words
./pomo-cli add "Pipoca Maluca"
```
> If a task contains spaces, wrap it in double quotes.

---

### Mark a Task as Completed
**Marks a task as completed based on ID or description.**

```sh
./pomo-cli done <task_id | task_description>
```

#### Examples:
```sh
# Mark a task as completed using an ID
./pomo-cli done 1

# Mark a task as completed using a description
./pomo-cli done "Pipoca Maluca"
```

---

### List All Tasks
**Displays all tasks, showing their IDs, descriptions, and completion status.**

```sh
./pomo-cli list
```

> Use flags to filter tasks if necessary.

---

### Remove a Task
**Removes a task or multiple tasks based on ID or description.**

```sh
./pomo-cli remove <task_id | task_description>
```

#### Examples:
```sh
# Remove a task by ID
./pomo-cli remove 1

# Remove tasks by description
./pomo-cli remove Pipoca "Pipoca Maluca"
```

---

### Search for Tasks
**Search for tasks based on ID or description.**

```sh
./pomo-cli search <task_id | task_description>
```

#### Examples:
```sh
# Search for a task by ID
./pomo-cli search 2

# Search for tasks containing the word "Pipoca"
./pomo-cli search Pipoca

# Search for tasks with an exact phrase
./pomo-cli search "Pipoca Maluca"
```

---

### Start a Pomodoro Session
**Starts a Pomodoro session for a specified task or creates a new one.**

```sh
./pomo-cli start <number_of_pomodoros> [-d="task_description"]
```

#### Examples:
```sh
# Start a Pomodoro session with 2 Pomodoros
./pomo-cli start 2

# Start a Pomodoro session for a specific task (or create it if it doesn't exist)
./pomo-cli start 1 -d="Pipoca Maluca"

# Start a Pomodoro session for a specific task using its ID
./pomo-cli start 1 -d="1"
```

---

## License
This project is licensed under the MIT License.

