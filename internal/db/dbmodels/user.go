package dbmodels

import (
	"fmt"
	"time"
)

// User represents a user account
type User struct {
	UserID    int       `ksql:"user_id" json:"user_id"`
	Username  string    `ksql:"username" json:"username"`
	Email     string    `ksql:"email" json:"email"`
	CreatedAt time.Time `ksql:"created_at" json:"created_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("User ID: %d, Username: %s, Email: %s, Created At: %v\n",
		u.UserID, u.Username, u.Email, u.CreatedAt)
}
