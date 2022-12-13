package models

import (
	"database/sql"
	"time"
)

// Attendance the attendance for the user
type Attendance struct {
	ID        uint         `json:"id"`
	CourseID  uint         `json:"course_id"`
	Users     []*User      `json:"users"`
	OpenedAt  sql.NullTime `json:"opened_at"`
	ClosedAt  sql.NullTime `json:"closed_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt time.Time    `json:"deleted_at"`
}
