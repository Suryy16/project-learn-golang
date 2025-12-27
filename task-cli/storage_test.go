package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestNewStorage(t *testing.T) {
	// Test creating new storage with valid filename
	storage := NewStorage[Todos]("test.json")
	if storage == nil {
		t.Error("Expected non-nil storage")
	}
	if storage.FileName != "test.json" {
		t.Errorf("Expected filename 'test.json', got '%s'", storage.FileName)
	}

	// Test with different types
	storageString := NewStorage[string]("string.json")
	if storageString.FileName != "string.json" {
		t.Errorf("Expected filename 'string.json', got '%s'", storageString.FileName)
	}

	// Test with empty filename
	storageEmpty := NewStorage[Todos]("")
	if storageEmpty.FileName != "" {
		t.Errorf("Expected empty filename, got '%s'", storageEmpty.FileName)
	}
}

func TestStorageSave(t *testing.T) {
	testFile := "test_save.json"
	defer os.Remove(testFile) // Cleanup after test

	storage := NewStorage[Todos](testFile)

	// Test saving valid todos
	todos := Todos{
		{
			ID:          1,
			Description: "Task 1",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
		{
			ID:          2,
			Description: "Task 2",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
	}

	err := storage.Save(todos)
	if err != nil {
		t.Errorf("Expected no error saving todos, got %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}

	// Verify file content
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
	}

	var loadedTodos Todos
	err = json.Unmarshal(data, &loadedTodos)
	if err != nil {
		t.Errorf("Failed to unmarshal saved data: %v", err)
	}

	if len(loadedTodos) != len(todos) {
		t.Errorf("Expected %d todos, got %d", len(todos), len(loadedTodos))
	}

	// Test saving empty list
	emptyTodos := Todos{}
	err = storage.Save(emptyTodos)
	if err != nil {
		t.Errorf("Expected no error saving empty todos, got %v", err)
	}

	data, _ = os.ReadFile(testFile)
	var loadedEmpty Todos
	json.Unmarshal(data, &loadedEmpty)
	if len(loadedEmpty) != 0 {
		t.Errorf("Expected 0 todos, got %d", len(loadedEmpty))
	}
}

func TestStorageSaveInvalidPath(t *testing.T) {
	// Test saving to invalid path (directory that doesn't exist and can't be created)
	storage := NewStorage[Todos]("/invalid/path/that/does/not/exist/test.json")
	todos := Todos{
		{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now()},
	}

	err := storage.Save(todos)
	if err == nil {
		t.Error("Expected error saving to invalid path, got nil")
	}
}

func TestStorageLoad(t *testing.T) {
	testFile := "test_load.json"
	defer os.Remove(testFile) // Cleanup after test

	// Create test data
	originalTodos := Todos{
		{
			ID:          1,
			Description: "Task 1",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
		{
			ID:          2,
			Description: "Task 2",
			Status:      "in-progress",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
		{
			ID:          3,
			Description: "Task 3",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
	}

	// Save test data
	data, _ := json.MarshalIndent(originalTodos, "", "   ")
	os.WriteFile(testFile, data, 0644)

	// Test loading
	storage := NewStorage[Todos](testFile)
	loadedTodos := Todos{}
	err := storage.Load(&loadedTodos)

	if err != nil {
		t.Errorf("Expected no error loading todos, got %v", err)
	}

	if len(loadedTodos) != len(originalTodos) {
		t.Errorf("Expected %d todos, got %d", len(originalTodos), len(loadedTodos))
	}

	for i := range originalTodos {
		if loadedTodos[i].ID != originalTodos[i].ID {
			t.Errorf("Todo %d: Expected ID %d, got %d", i, originalTodos[i].ID, loadedTodos[i].ID)
		}
		if loadedTodos[i].Description != originalTodos[i].Description {
			t.Errorf("Todo %d: Expected description '%s', got '%s'", i, originalTodos[i].Description, loadedTodos[i].Description)
		}
		if loadedTodos[i].Status != originalTodos[i].Status {
			t.Errorf("Todo %d: Expected status '%s', got '%s'", i, originalTodos[i].Status, loadedTodos[i].Status)
		}
	}
}

func TestStorageLoadNonExistentFile(t *testing.T) {
	// Test loading from non-existent file
	storage := NewStorage[Todos]("non_existent_file.json")
	todos := Todos{}
	err := storage.Load(&todos)

	if err == nil {
		t.Error("Expected error loading non-existent file, got nil")
	}
}

func TestStorageLoadInvalidJSON(t *testing.T) {
	testFile := "test_invalid.json"
	defer os.Remove(testFile)

	// Create invalid JSON file
	invalidJSON := []byte("{ invalid json content }")
	os.WriteFile(testFile, invalidJSON, 0644)

	// Test loading invalid JSON
	storage := NewStorage[Todos](testFile)
	todos := Todos{}
	err := storage.Load(&todos)

	if err == nil {
		t.Error("Expected error loading invalid JSON, got nil")
	}
}

func TestStorageLoadEmptyFile(t *testing.T) {
	testFile := "test_empty.json"
	defer os.Remove(testFile)

	// Create empty file
	os.WriteFile(testFile, []byte(""), 0644)

	// Test loading empty file
	storage := NewStorage[Todos](testFile)
	todos := Todos{}
	err := storage.Load(&todos)

	if err == nil {
		t.Error("Expected error loading empty file, got nil")
	}
}

func TestStorageSaveAndLoad(t *testing.T) {
	testFile := "test_save_load.json"
	defer os.Remove(testFile)

	storage := NewStorage[Todos](testFile)

	// Create original todos with various scenarios
	now := time.Now()
	updatedTime := now.Add(1 * time.Hour)
	originalTodos := Todos{
		{
			ID:          1,
			Description: "First task",
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   nil,
		},
		{
			ID:          2,
			Description: "Second task with special chars: @#$%",
			Status:      "in-progress",
			CreatedAt:   now,
			UpdatedAt:   &updatedTime,
		},
		{
			ID:          3,
			Description: "",
			Status:      "done",
			CreatedAt:   now,
			UpdatedAt:   nil,
		},
	}

	// Save
	err := storage.Save(originalTodos)
	if err != nil {
		t.Errorf("Failed to save: %v", err)
	}

	// Load
	loadedTodos := Todos{}
	err = storage.Load(&loadedTodos)
	if err != nil {
		t.Errorf("Failed to load: %v", err)
	}

	// Verify
	if len(loadedTodos) != len(originalTodos) {
		t.Errorf("Expected %d todos, got %d", len(originalTodos), len(loadedTodos))
	}

	for i := range originalTodos {
		if loadedTodos[i].ID != originalTodos[i].ID {
			t.Errorf("Mismatch at index %d: ID", i)
		}
		if loadedTodos[i].Description != originalTodos[i].Description {
			t.Errorf("Mismatch at index %d: Description", i)
		}
		if loadedTodos[i].Status != originalTodos[i].Status {
			t.Errorf("Mismatch at index %d: Status", i)
		}
		if (originalTodos[i].UpdatedAt == nil) != (loadedTodos[i].UpdatedAt == nil) {
			t.Errorf("Mismatch at index %d: UpdatedAt nil status", i)
		}
	}
}

func TestStorageWithDifferentTypes(t *testing.T) {
	// Test with string type
	testFileString := "test_string.json"
	defer os.Remove(testFileString)

	storageString := NewStorage[string](testFileString)
	testString := "Hello, World!"

	err := storageString.Save(testString)
	if err != nil {
		t.Errorf("Failed to save string: %v", err)
	}

	var loadedString string
	err = storageString.Load(&loadedString)
	if err != nil {
		t.Errorf("Failed to load string: %v", err)
	}

	if loadedString != testString {
		t.Errorf("Expected '%s', got '%s'", testString, loadedString)
	}

	// Test with int slice
	testFileInts := "test_ints.json"
	defer os.Remove(testFileInts)

	storageInts := NewStorage[[]int](testFileInts)
	testInts := []int{1, 2, 3, 4, 5}

	err = storageInts.Save(testInts)
	if err != nil {
		t.Errorf("Failed to save ints: %v", err)
	}

	var loadedInts []int
	err = storageInts.Load(&loadedInts)
	if err != nil {
		t.Errorf("Failed to load ints: %v", err)
	}

	if len(loadedInts) != len(testInts) {
		t.Errorf("Expected %d ints, got %d", len(testInts), len(loadedInts))
	}
}
