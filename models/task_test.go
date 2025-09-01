package models

import (
	"os"
	"testing"
)

func setupAndTeardown(t *testing.T) *TaskList {
	t.Cleanup(func() {
		os.Remove(TaskFile)
	})

	return &TaskList{}
}

func TestAddTask(t *testing.T) {
	tl := &TaskList{}

	tl.AddTask("Test title", "test description")

	if len(tl.Tasks) != 1 {
		t.Errorf("expected task list lenght to be 1, got %d", len(tl.Tasks))
	}

	addedTask := tl.Tasks[0]
	if addedTask.Title != "Test title" {
		t.Errorf("task title is '%s'", addedTask.Title)
	}
	if addedTask.Completed != false {
		t.Errorf("shouldnt be completed")
	}
	if addedTask.ID != 1 {
		t.Errorf("id is %d", addedTask.ID)
	}
}

func TestGetNextID(t *testing.T) {
	// 1. Test on an empty list
	tl := &TaskList{}
	if id := tl.GetNextID(); id != 1 {
		t.Errorf("Expected next ID to be 1 for empty list, but got %d", id)
	}

	// 2. Test on a list with items
	tl.Tasks = []Task{{ID: 1}, {ID: 5}, {ID: 3}}
	if id := tl.GetNextID(); id != 6 {
		t.Errorf("Expected next ID to be 6, but got %d", id)
	}
}

func TestCompleteTask(t *testing.T) {
	tl := &TaskList{}
	tl.AddTask("Incomplete Task", "Some description")

	taskID := tl.Tasks[0].ID

	// 1. Test completing an existing task
	err := tl.CompleteTask(taskID)
	if err != nil {
		t.Errorf("Did not expect an error when completing a task, but got %v", err)
	}
	if !tl.Tasks[0].Completed {
		t.Error("Expected task to be marked as completed")
	}
	if tl.Tasks[0].CompletedAt == nil {
		t.Error("Expected CompletedAt to be set")
	}

	// 2. Test completing a non-existent task
	err = tl.CompleteTask(999)
	if err == nil {
		t.Error("Expected an error when completing a non-existent task, but got nil")
	}
}

func TestDeleteTask(t *testing.T) {
	tl := &TaskList{}
	tl.AddTask("Task to delete", "desc 1")
	tl.AddTask("Another task", "desc 2")

	taskIDToDelete := tl.Tasks[0].ID

	// 1. Test deleting an existing task
	msg, err := tl.DeleteTask(taskIDToDelete)
	if err != nil {
		t.Errorf("Did not expect an error when deleting a task, but got %v", err)
	}
	if len(tl.Tasks) != 1 {
		t.Errorf("Expected task list length to be 1 after deletion, but got %d", len(tl.Tasks))
	}
	if tl.Tasks[0].ID == taskIDToDelete {
		t.Error("The deleted task is still in the list")
	}
    if msg == "" {
        t.Error("Expected a success message, but got an empty string")
    }

	// 2. Test deleting a non-existent task
	_, err = tl.DeleteTask(999)
	if err == nil {
		t.Error("Expected an error when deleting a non-existent task, but got nil")
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	// Use the helper to handle file cleanup
	tl := setupAndTeardown(t)

	// 1. Create some tasks and save them
	tl.AddTask("First Task", "Description 1")
	tl.AddTask("Second Task", "Description 2")
	tl.CompleteTask(1) // Mark the first task as complete for a more robust test

	err := SaveTasks(tl)
	if err != nil {
		t.Fatalf("Failed to save tasks: %v", err)
	}

	// 2. Load the tasks back into a new list
	loadedTl, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	// 3. Verify the loaded data
	if len(loadedTl.Tasks) != 2 {
		t.Fatalf("Expected to load 2 tasks, but got %d", len(loadedTl.Tasks))
	}

	// Check details of the loaded tasks
	if loadedTl.Tasks[0].Title != "First Task" || !loadedTl.Tasks[0].Completed {
		t.Error("First loaded task data does not match original")
	}
    if loadedTl.Tasks[0].CompletedAt == nil {
        t.Error("First loaded task's CompletedAt field should not be nil")
    }
	if loadedTl.Tasks[1].Title != "Second Task" || loadedTl.Tasks[1].Completed {
		t.Error("Second loaded task data does not match original")
	}
}

func TestLoadNonExistentFile(t *testing.T) {
    // Ensure no tasks.json exists before this test
	os.Remove(TaskFile) 

    tl, err := LoadTasks()
    if err != nil {
        t.Fatalf("Loading a non-existent file should not produce an error, but got %v", err)
    }
    if len(tl.Tasks) != 0 {
        t.Errorf("Expected task list to be empty, but got %d tasks", len(tl.Tasks))
    }
}