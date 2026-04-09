package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var taskFilePath = mustTaskFilePath()

func mustTaskFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error resolving config directory: %v\n", err)
	}

	return filepath.Join(configDir, "task-cli", "tasks.json")
}

func LoadTasks() ([]Task, error) {
	_, err := os.Stat(taskFilePath)
	if err != nil && os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(taskFilePath)
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

	err = os.MkdirAll(filepath.Dir(taskFilePath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(taskFilePath, data, 0644)
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