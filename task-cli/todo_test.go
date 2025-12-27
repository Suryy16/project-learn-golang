package main

import (
	"testing"
	"time"
)

func TestTodosAdd(t *testing.T) {
	todos := Todos{}

	// Test adding first todo
	todos.add("First task")
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}

	// Verify first todo properties
	if todos[0].ID != 1 {
		t.Errorf("Expected ID 1, got %d", todos[0].ID)
	}
	if todos[0].Description != "First task" {
		t.Errorf("Expected description 'First task', got '%s'", todos[0].Description)
	}
	if todos[0].Status != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", todos[0].Status)
	}
	if todos[0].UpdatedAt != nil {
		t.Error("Expected UpdatedAt to be nil for new todo")
	}

	// Test adding second todo
	todos.add("Second task")
	if len(todos) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(todos))
	}
	if todos[1].ID != 2 {
		t.Errorf("Expected ID 2, got %d", todos[1].ID)
	}

	// Test adding empty description
	todos.add("")
	if len(todos) != 3 {
		t.Errorf("Expected 3 todos, got %d", len(todos))
	}
	if todos[2].Description != "" {
		t.Errorf("Expected empty description, got '%s'", todos[2].Description)
	}
}

func TestTodosDelete(t *testing.T) {
	// Test deleting from valid index
	todos := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "todo"},
		{ID: 3, Description: "Task 3", Status: "todo"},
	}

	err := todos.delete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(todos) != 2 {
		t.Errorf("Expected 2 todos after deletion, got %d", len(todos))
	}
	if todos[0].Description != "Task 1" {
		t.Errorf("Expected first task to be 'Task 1', got '%s'", todos[0].Description)
	}
	if todos[1].Description != "Task 3" {
		t.Errorf("Expected second task to be 'Task 3', got '%s'", todos[1].Description)
	}

	// Test deleting first element
	todos2 := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "todo"},
	}
	err = todos2.delete(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(todos2) != 1 {
		t.Errorf("Expected 1 todo after deletion, got %d", len(todos2))
	}

	// Test deleting last element
	todos3 := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "todo"},
	}
	err = todos3.delete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(todos3) != 1 {
		t.Errorf("Expected 1 todo after deletion, got %d", len(todos3))
	}

	// Test deleting from invalid index (negative)
	todos4 := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
	}
	err = todos4.delete(-1)
	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}

	// Test deleting from invalid index (out of bounds)
	err = todos4.delete(5)
	if err == nil {
		t.Error("Expected error for out of bounds index, got nil")
	}

	// Test deleting from empty list
	todos5 := Todos{}
	err = todos5.delete(0)
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}
}

func TestTodosValidateIndex(t *testing.T) {
	todos := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "todo"},
		{ID: 3, Description: "Task 3", Status: "todo"},
	}

	// Test valid indices
	validIndices := []int{0, 1, 2}
	for _, idx := range validIndices {
		err := todos.ValidateIndex(idx)
		if err != nil {
			t.Errorf("Expected no error for valid index %d, got %v", idx, err)
		}
	}

	// Test invalid indices
	invalidIndices := []int{-1, -10, 3, 4, 100}
	for _, idx := range invalidIndices {
		err := todos.ValidateIndex(idx)
		if err == nil {
			t.Errorf("Expected error for invalid index %d, got nil", idx)
		}
	}

	// Test with empty list
	emptyTodos := Todos{}
	err := emptyTodos.ValidateIndex(0)
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}
}

func TestTodosUpdate(t *testing.T) {
	// Test updating valid todo
	todos := Todos{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "todo"},
		{ID: 3, Description: "Task 3", Status: "todo"},
	}

	err := todos.update("Updated Task 1", 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[0].Description != "Updated Task 1" {
		t.Errorf("Expected 'Updated Task 1', got '%s'", todos[0].Description)
	}

	// Test updating middle element
	err = todos.update("Updated Task 2", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[1].Description != "Updated Task 2" {
		t.Errorf("Expected 'Updated Task 2', got '%s'", todos[1].Description)
	}

	// Test updating last element
	err = todos.update("Updated Task 3", 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[2].Description != "Updated Task 3" {
		t.Errorf("Expected 'Updated Task 3', got '%s'", todos[2].Description)
	}

	// Test updating with empty string
	err = todos.update("", 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[0].Description != "" {
		t.Errorf("Expected empty description, got '%s'", todos[0].Description)
	}

	// Test updating with invalid index (negative)
	err = todos.update("Invalid", -1)
	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}

	// Test updating with invalid index (out of bounds)
	err = todos.update("Invalid", 10)
	if err == nil {
		t.Error("Expected error for out of bounds index, got nil")
	}

	// Test updating empty list
	emptyTodos := Todos{}
	err = emptyTodos.update("Should fail", 0)
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}
}

func TestTodosStatusChange(t *testing.T) {
	// Test changing status to in-progress
	todos := Todos{
		{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now()},
		{ID: 2, Description: "Task 2", Status: "todo", CreatedAt: time.Now()},
		{ID: 3, Description: "Task 3", Status: "todo", CreatedAt: time.Now()},
	}

	beforeUpdate := time.Now()
	err := todos.StatusChange("mark:in-progress", 0)
	afterUpdate := time.Now()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[0].Status != "in-progress" {
		t.Errorf("Expected status 'in-progress', got '%s'", todos[0].Status)
	}
	if todos[0].UpdatedAt == nil {
		t.Error("Expected UpdatedAt to be set")
	} else {
		if todos[0].UpdatedAt.Before(beforeUpdate) || todos[0].UpdatedAt.After(afterUpdate) {
			t.Error("UpdatedAt is not within expected time range")
		}
	}

	// Test changing status to done
	err = todos.StatusChange("mark:done", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[1].Status != "done" {
		t.Errorf("Expected status 'done', got '%s'", todos[1].Status)
	}
	if todos[1].UpdatedAt == nil {
		t.Error("Expected UpdatedAt to be set")
	}

	// Test changing status to todo
	err = todos.StatusChange("mark:todo", 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[2].Status != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", todos[2].Status)
	}

	// Test with custom status
	err = todos.StatusChange("mark:custom-status", 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[0].Status != "custom-status" {
		t.Errorf("Expected status 'custom-status', got '%s'", todos[0].Status)
	}

	// Test with invalid index (negative)
	err = todos.StatusChange("mark:done", -1)
	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}

	// Test with invalid index (out of bounds)
	err = todos.StatusChange("mark:done", 10)
	if err == nil {
		t.Error("Expected error for out of bounds index, got nil")
	}

	// Test with empty list
	emptyTodos := Todos{}
	err = emptyTodos.StatusChange("mark:done", 0)
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}

	// Test with empty status
	err = todos.StatusChange("mark:", 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if todos[0].Status != "" {
		t.Errorf("Expected empty status, got '%s'", todos[0].Status)
	}
}

func TestTodoStruct(t *testing.T) {
	// Test creating a new Todo
	now := time.Now()
	todo := Todo{
		ID:          1,
		Description: "Test task",
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   nil,
	}

	if todo.ID != 1 {
		t.Errorf("Expected ID 1, got %d", todo.ID)
	}
	if todo.Description != "Test task" {
		t.Errorf("Expected description 'Test task', got '%s'", todo.Description)
	}
	if todo.Status != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", todo.Status)
	}
	if todo.CreatedAt != now {
		t.Error("CreatedAt mismatch")
	}
	if todo.UpdatedAt != nil {
		t.Error("Expected UpdatedAt to be nil")
	}

	// Test with UpdatedAt set
	updatedTime := time.Now()
	todo2 := Todo{
		ID:          2,
		Description: "Updated task",
		Status:      "done",
		CreatedAt:   now,
		UpdatedAt:   &updatedTime,
	}

	if todo2.UpdatedAt == nil {
		t.Error("Expected UpdatedAt to be set")
	} else {
		if *todo2.UpdatedAt != updatedTime {
			t.Error("UpdatedAt mismatch")
		}
	}
}
