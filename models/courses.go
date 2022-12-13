package models

import (
	"time"
)

// Course this will implement the model for the course
type Course struct {
	ID         uint       `json:"id"`
	ClassID    uint       `json:"class_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
	Attendance Attendance `json:"attendance"`
}
