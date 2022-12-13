package models

import (
	"time"
)

// Class this will implement the model for the class
type Class struct {
	ID        uint      `json:"id"`
	Mentors   []*User   `json:"mentors"`
	Members   []*User   `json:"members"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Logo      string    `json:"logo"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Courses   []Course  `json:"courses"`
}
