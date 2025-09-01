package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const TaskFile = "tasks.json"

func LoadTasks() (*TaskList, error) {
	taskList := &TaskList{}

	// checking if file exists
	if _, err := os.Stat(TaskFile); os.IsNotExist(err) {
		return taskList, nil
	}

	// reading file into memory
	data, err := os.ReadFile(TaskFile)
	if err != nil {
		return nil, fmt.Errorf("error reading tasks file: %v", err)
	}

	if len(data) == 0 {
		return taskList, nil
	}

	// parsing json data into struct
	err = json.Unmarshal(data, taskList)
	if err != nil {
		return nil, fmt.Errorf("error parsing tasks file: %v", err)
	}

	return taskList, nil
}

func SaveTasks(taskList *TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("error mashalling tasks: %v", err)
	}

	// writing read data to the taskFile
	err = os.WriteFile(TaskFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing tasks file: %v", err)
	}

	return nil
}

// giving tasks an id

func (tl *TaskList) GetNextID() int {
	if len(tl.Tasks) == 0 {
		return 1
	}

	//avoiding duplicates
	maxID := 0
	for _, task := range tl.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func (tl *TaskList) AddTask(title, description string) {
	task := Task{
		ID: tl.GetNextID(),
		Title: title,
		Description: description,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) CompleteTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			now := time.Now()
			tl.Tasks[i].Completed = true
			tl.Tasks[i].CompletedAt = &now
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) DeleteTask(id int) (string, error ){
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]... )
			return fmt.Sprintf("deleted task with id: %d", id), nil
		}
	}

	return "", fmt.Errorf("task with ID %d not found", id)
}