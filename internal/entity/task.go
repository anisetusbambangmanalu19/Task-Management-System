package entity

import "time"

type Task struct {
	ID          int
	ProjectID   int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}
