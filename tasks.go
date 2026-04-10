package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	defaultTaskStatus   = "todo"
	defaultTaskPriority = "low"
)

var validPriorities = map[string]struct{}{
	"high":   {},
	"medium": {},
	"low":    {},
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var tasks []Task

func isValidPriority(priority string) bool {
	_, ok := validPriorities[priority]
	return ok
}

func normalizePriority(priority string) string {
	return strings.ToLower(strings.TrimSpace(priority))
}

func AddTask(description string, priority string) error {
	priority = normalizePriority(priority)
	if !isValidPriority(priority) {
		return errors.New("invalid priority: use high, medium, or low")
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          id,
		Description: description,
		Status:      defaultTaskStatus,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func ListTasks(status string) error {
	priorityOrder := map[string]int{
		"high":   0,
		"medium": 1,
		"low":    2,
	}

	visibleTasks := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		if (status == "" && task.Status != "done") || task.Status == status {
			visibleTasks = append(visibleTasks, task)
		}
	}

	if len(visibleTasks) == 0 {
		if status == "" {
			return errors.New("no task available")
		}
		return errors.New("no task available with the status: " + status)
	}

	sort.Slice(visibleTasks, func(i, j int) bool {
		return priorityOrder[visibleTasks[i].Priority] < priorityOrder[visibleTasks[j].Priority]
	})

	maxDescriptionWidth := utf8.RuneCountInString("Description")
	for _, task := range visibleTasks {
		if descriptionWidth := utf8.RuneCountInString(task.Description); descriptionWidth > maxDescriptionWidth {
			maxDescriptionWidth = descriptionWidth
		}
	}
	if maxDescriptionWidth > 48 {
		maxDescriptionWidth = 48
	}

	rowWidth := 4 + 1 + maxDescriptionWidth + 1 + 12 + 1 + 8

	fmt.Println()
	fmt.Println("TASK LIST")
	fmt.Println(strings.Repeat("-", rowWidth))
	fmt.Printf("%-4s %-*s %-12s %-8s\n", "ID", maxDescriptionWidth, "Description", "Status", "Priority")
	fmt.Println(strings.Repeat("-", rowWidth))

	for _, task := range visibleTasks {
		descriptionLines := wrapText(task.Description, maxDescriptionWidth)
		for lineIndex, descriptionLine := range descriptionLines {
			if lineIndex == 0 {
				fmt.Printf("%-4d %-*s %-12s %-8s\n", task.ID, maxDescriptionWidth, descriptionLine, task.Status, task.Priority)
				continue
			}

			fmt.Printf("%-4s %-*s %-12s %-8s\n", "", maxDescriptionWidth, descriptionLine, "", "")
		}
	}

	fmt.Println()
	return nil
}

func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}

	runes := []rune(text)
	if len(runes) == 0 {
		return []string{""}
	}

	lines := make([]string, 0, (len(runes)+width-1)/width)
	for start := 0; start < len(runes); start += width {
		end := start + width
		if end > len(runes) {
			end = len(runes)
		}
		lines = append(lines, string(runes[start:end]))
	}

	return lines
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
			tasks = slices.Delete(tasks, i, i+1)
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
		tasks[i].Status = defaultTaskStatus
	}
	return SaveTasks(tasks)
}

func ClearTasks() error {
	tasks = []Task{}
	return SaveTasks(tasks)
}

func SetPriority(id int, newPriority string) error {
	newPriority = normalizePriority(newPriority)
	if !isValidPriority(newPriority) {
		return errors.New("invalid priority: use high, medium, or low")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Priority = newPriority
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
