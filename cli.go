package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Name string
	Done bool
}

var tasks = []Task{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			markTaskDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Exiting application.")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task name: ")
	scanner.Scan()
	taskName := scanner.Text()
	tasks = append(tasks, Task{Name: taskName, Done: false})
	fmt.Println("Added task:", taskName)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Name, status)
	}
}

func markTaskDone(scanner *bufio.Scanner) {
	listTasks()
	fmt.Print("Enter task number to mark as done: ")
	scanner.Scan()
	var taskNum int
	_, err := fmt.Sscanf(scanner.Text(), "%d", &taskNum)
	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[taskNum-1].Done = true
	fmt.Println("Marked task as done:", tasks[taskNum-1].Name)
}

func deleteTask(scanner *bufio.Scanner) {
	listTasks()
	fmt.Print("Enter task number to delete: ")
	scanner.Scan()
	var taskNum int
	_, err := fmt.Sscanf(scanner.Text(), "%d", &taskNum)
	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
}
