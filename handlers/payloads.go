package handlers

import "time"

type UserPayload struct {
	ID          uint      `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	IsSuperuser bool      `json:"is_superuser"`
	IsMentor    bool      `json:"is_mentor"`
	LastLogin   time.Time `json:"last_login"`
	DateJoined  time.Time `json:"date_joined"`
	Token       string    `json:"token"`
}
