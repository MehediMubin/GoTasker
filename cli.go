package main

import (
	"errors"
	"fmt"
	"os"
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
	default:
		return errors.New("unknown command")
	}

	return nil
}