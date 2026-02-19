package entity

import "time"

type Project struct {
	ID          int
	UserID      int
	Name        string
	Description string
	CreatedAt   time.Time
}
