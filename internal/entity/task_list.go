package entity

import "fmt"

type TaskList struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (t TaskList) String() string {
	return fmt.Sprintf("TaskList{ID: %s, Title: %s}", t.ID, t.Title)
}
