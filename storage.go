package main

import (
	"encoding/json"
	"log"
	"os"
)

const taskFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	_, err := os.Stat(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	var loadedTasks []Task
	err = json.Unmarshal(data, &loadedTasks)
	if err != nil {
		return nil, err
	}
	
	return loadedTasks, nil
}

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

func init() {
	var err error
	tasks, err = LoadTasks()
	if err != nil {
		log.Fatalf("Error loading tasks: %v\n", err)
	}
}