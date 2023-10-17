package dtomodel

import (
	"fmt"
	"time"
)

// User represents a user account
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("User ID: %d, Username: %s, Email: %s, Created At: %v\n",
		u.UserID, u.Username, u.Email, u.CreatedAt)
}
