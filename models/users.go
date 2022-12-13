package models

import (
	"database/sql"
	"time"
)

// User is implement for user
type User struct {
	ID          uint          `json:"id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"-"`
	IsSuperuser bool          `json:"is_superuser"`
	IsMentor    bool          `json:"is_mentor"`
	LastLogin   sql.NullTime  `json:"last_login"`
	DateJoined  time.Time     `json:"date_joined"`
	Mentors     []*Class      `json:"-"`
	Members     []*Class      `json:"-"`
	Attendances []*Attendance `json:"-"`
}
