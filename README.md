# Go CLI To-Do List

## Overview

This is a simple command-line to-do list application written in Go (Golang). It allows users to manage their tasks directly from the terminal. The application supports basic functionalities such as adding a new task, listing all tasks, marking a task as completed, and deleting a task.

## Features

- **Add a Task**: Users can add new tasks to their to-do list.
- **List Tasks**: Display all the tasks along with their completion status.
- **Mark Task as Done**: Users can mark tasks as completed.
- **Delete a Task**: Allows users to delete tasks from the list.

## Installation

To use this application, you must have Go installed on your machine.

1. Clone the repository:
```bash
   git clone [repository-url]
```
2. Navigate to the project directory:
```bash
cd go-cli-todo
```

## Usage
To run the application, use the following command in the project directory:

```bash
go run cli.go

```

Once the application starts, it will display a menu with options to add, list, mark, or delete tasks. Follow the on-screen prompts to manage your to-do list.

## Structure

- **cli.go**: The main file containing the source code of the application.

- **Task struct**: A structure to define a task with a Name and Done status.

## Contributing
Contributions to this project are welcome. Please ensure to update tests as appropriate.


