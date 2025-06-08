package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunCLI() error {
	args := os.Args
	if len(args) < 2 {
		return errors.New("please provide a command")
	}

	command := args[1]
	switch command {
	case "add":
		if len(args) < 3 {
			return errors.New("please provide a task description")
		}
		
		description := strings.Join(args[2:], " ")
		err := AddTask(description)
		if err != nil {
			return err
		}
		fmt.Println("Task added successfully")

	case "show":
		status := ""
		if len(args) > 2 {
			status = args[2]
		}
		ShowTasks(status)
	
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
			return errors.New("usage: done <id>")
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

	default:
		return errors.New("unknown command")
	}

	return nil
}