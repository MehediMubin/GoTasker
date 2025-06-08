package main

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

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

func ShowTasks(status string) {
	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func UpdateTask(id int, newDescription string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = slices.Delete(tasks, i, i + 1)
			return SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}