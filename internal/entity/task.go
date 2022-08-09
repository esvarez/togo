package entity

import "fmt"

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

func (t Task) String() string {
	return fmt.Sprintf("Task{ID: %s, Name: %s, Description: %s}", t.ID, t.Name, t.Description)
}
