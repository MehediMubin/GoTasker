package main

import (
	"encoding/json"
	"os"
)

const taskFile = "tasks.json"

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return err
	}
	
	return nil
}