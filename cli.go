package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type commandHelp struct {
	Usage       string
	Description string
	Examples    []string
}

var commandDocs = map[string]commandHelp{
	"add": {
		Usage:       "add <description> [priority]",
		Description: "Add a new task with an optional priority (high, medium, low).",
		Examples: []string{
			"task-cli add \"Buy groceries\"",
			"task-cli add \"Prepare slides\" high",
		},
	},
	"list": {
		Usage:       "list [status]",
		Description: "List tasks. Without a status, shows non-done tasks.",
		Examples: []string{
			"task-cli list",
			"task-cli list todo",
			"task-cli list in-progress",
			"task-cli list done",
		},
	},
	"update": {
		Usage:       "update <id> <new description>",
		Description: "Update a task description by ID.",
		Examples: []string{
			"task-cli update 2 \"Finish assignment\"",
		},
	},
	"delete": {
		Usage:       "delete <id>",
		Description: "Delete a single task by ID.",
		Examples: []string{
			"task-cli delete 3",
		},
	},
	"mark-in-progress": {
		Usage:       "mark-in-progress <id>",
		Description: "Mark a task as in-progress.",
		Examples: []string{
			"task-cli mark-in-progress 2",
		},
	},
	"mark-done": {
		Usage:       "mark-done <id>",
		Description: "Mark a task as done.",
		Examples: []string{
			"task-cli mark-done 2",
		},
	},
	"reset": {
		Usage:       "reset",
		Description: "Reset all tasks to todo status.",
		Examples: []string{
			"task-cli reset",
		},
	},
	"clear": {
		Usage:       "clear",
		Description: "Delete all tasks.",
		Examples: []string{
			"task-cli clear",
		},
	},
	"priority-high": {
		Usage:       "priority-high <id>",
		Description: "Set task priority to high.",
		Examples: []string{
			"task-cli priority-high 1",
		},
	},
	"priority-mid": {
		Usage:       "priority-mid <id>",
		Description: "Set task priority to medium.",
		Examples: []string{
			"task-cli priority-mid 1",
		},
	},
	"priority-low": {
		Usage:       "priority-low <id>",
		Description: "Set task priority to low.",
		Examples: []string{
			"task-cli priority-low 1",
		},
	},
	"help": {
		Usage:       "help [command]",
		Description: "Show all commands or detailed help for a specific command.",
		Examples: []string{
			"task-cli help",
			"task-cli help add",
		},
	},
}

func printGeneralHelp() {
	fmt.Println("GoTasker - command reference")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task-cli <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")

	keys := make([]string, 0, len(commandDocs))
	for cmd := range commandDocs {
		keys = append(keys, cmd)
	}
	sort.Strings(keys)

	for _, cmd := range keys {
		doc := commandDocs[cmd]
		fmt.Printf("  %-18s %s\n", cmd, doc.Description)
	}

	fmt.Println()
	fmt.Println("Run 'task-cli help <command>' for details.")
}

func printCommandHelp(command string) error {
	doc, ok := commandDocs[command]
	if !ok {
		return errors.New("unknown command: " + command)
	}

	fmt.Printf("Command: %s\n", command)
	fmt.Printf("Usage:   task-cli %s\n", doc.Usage)
	fmt.Printf("About:   %s\n", doc.Description)

	if len(doc.Examples) > 0 {
		fmt.Println("Examples:")
		for _, example := range doc.Examples {
			fmt.Printf("  %s\n", example)
		}
	}

	return nil
}

func parseAddCommandArgs(args []string) (string, string, error) {
	if len(args) < 3 {
		return "", "", errors.New("usage: add <description> [priority]")
	}

	if len(args) > 4 {
		return "", "", errors.New("usage: add <description> [priority]")
	}

	description := args[2]
	priority := defaultTaskPriority
	if len(args) == 4 {
		priority = args[3]
	}

	return description, priority, nil
}

func RunCLI() error {
	args := os.Args
	if len(args) < 2 {
		printGeneralHelp()
		return nil
	}

	command := args[1]
	switch command {
	case "help":
		if len(args) == 2 {
			printGeneralHelp()
			return nil
		}
		if len(args) == 3 {
			return printCommandHelp(args[2])
		}
		return errors.New("usage: help [command]")

	case "add":
		description, priority, err := parseAddCommandArgs(args)
		if err != nil {
			return err
		}

		err = AddTask(description, priority)
		if err != nil {
			return err
		}
		fmt.Println("Task added successfully")

	case "list":
		status := ""
		if len(args) > 2 {
			status = args[2]
		}
		err := ListTasks(status)
		if err != nil {
			return err
		}

	case "update":
		if len(args) < 4 {
			return errors.New("usage: update <id> <new description>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		description := strings.Join(args[3:], " ")
		err = UpdateTask(id, description)
		if err != nil {
			return err
		}
		fmt.Println("Task updated successfully")

	case "delete":
		if len(args) < 3 {
			return errors.New("usage: delete <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		err = DeleteTask(id)
		if err != nil {
			return err
		}
		fmt.Println("Task deleted successfully")

	case "mark-in-progress":
		if len(args) < 3 {
			return errors.New("usage: mark-in-progress <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		err = MarkStatus(id, "in-progress")
		if err != nil {
			return err
		}
		fmt.Println("Status updated successfully")

	case "mark-done":
		if len(args) < 3 {
			return errors.New("usage: mark-done <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task id")
		}

		err = MarkStatus(id, "done")
		if err != nil {
			return err
		}
		fmt.Println("Status updated successfully")

	case "reset":
		if len(args) < 2 {
			return errors.New("usage: reset")
		}
		err := Reset()
		if err != nil {
			return err
		}
		fmt.Println("Task resetted successfully")

	case "clear":
		if len(args) != 2 {
			return errors.New("usage: clear")
		}

		err := ClearTasks()
		if err != nil {
			return err
		}
		fmt.Println("All tasks deleted successfully")

	case "priority-high":
		if len(args) < 3 {
			return errors.New("usage: priority-high <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		err = SetPriority(id, "high")
		if err != nil {
			return err
		}
		fmt.Println("Task priority set successfully")

	case "priority-mid":
		if len(args) < 3 {
			return errors.New("usage: priority-mid <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		err = SetPriority(id, "medium")
		if err != nil {
			return err
		}
		fmt.Println("Task priority set successfully")

	case "priority-low":
		if len(args) < 3 {
			return errors.New("usage: priority-low <id>")
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid task ID")
		}

		err = SetPriority(id, "low")
		if err != nil {
			return err
		}
		fmt.Println("Task priority set successfully")

	default:
		return errors.New("unknown command: run 'task-cli help' to view available commands")
	}

	return nil
}
