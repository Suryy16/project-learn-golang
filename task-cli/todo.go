package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	ID          int        `json:"ID"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type Todos []Todo

func (todos *Todos) add(description string) {
	id := len(*todos) + 1

	todo := Todo{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(ID int) error {
	if ID < 0 || ID >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(ID int) error {
	t := *todos
	if err := t.ValidateIndex(ID); err != nil {
		return err
	}

	*todos = append(t[:ID], t[ID+1:]...)

	return nil
}

func (todos *Todos) StatusChange(status string, ID int) error {
	t := *todos
	text := status[5:]
	if err := t.ValidateIndex(ID); err != nil {
		return err
	}

	t[ID].Status = text
	updateTime := time.Now()
	t[ID].UpdatedAt = &updateTime
	return nil
}

func (todos *Todos) update(description string, ID int) error {
	t := *todos

	if err := t.ValidateIndex(ID); err != nil {
		return err
	}

	t[ID].Description = description
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("id", "Description", "Status", "Created At", "Updated At")

	for id, t := range *todos {
		createdAt := t.CreatedAt.Format(time.RFC1123)
		updatedAt := ""
		if t.UpdatedAt == nil {
			updatedAt = ""
		} else {
			updatedAt = t.UpdatedAt.Format(time.RFC1123)
		}
		table.AddRow(strconv.Itoa(id), t.Description, t.Status, createdAt, updatedAt)
	}

	table.Render()
}
