package main

import "time"

type Task struct {
	ID          int    		`json:"id"`
	Description string 		`json:"description"`
	Status      string 		`json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
}

var tasks []Task

func AddTask(description string) error {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks) - 1].ID + 1
	}

	newTask := Task{
		ID: id,
		Description: description,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}