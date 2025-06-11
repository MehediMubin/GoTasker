package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"time"
)

type Task struct {
	ID          int    		`json:"id"`
	Description string 		`json:"description"`
	Status      string 		`json:"status"`
	Priority    string    `json:"priority"`
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
		Priority: "low",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask) 
	return SaveTasks(tasks)
}

func ListTasks(status string) error {
	priorityOrder := map[string]int{
		"high": 0,
		"medium": 1,
		"low": 2,
	}

	sort.Slice(tasks, func(i, j int) bool {
		return priorityOrder[tasks[i].Priority] < priorityOrder[tasks[j].Priority]
	})
	found := false
	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("%d - %s [%s] [%s]\n", task.ID, task.Description, task.Status, task.Priority)
			found = true
		}
	}
	if found {
		return nil
	} else {
		return errors.New("no task available with the status: " + status)
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

func MarkStatus(id int, status string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}

func Reset() error {
	for i := range tasks {
		tasks[i].Status = "todo"
	}
	return SaveTasks(tasks)
}

func SetPriority(id int, newPriority string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Priority = newPriority
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}