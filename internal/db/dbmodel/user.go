package dbmodel

import (
	"fmt"
	"time"
)

// User represents a user account
type User struct {
	UserID    int       `ksql:"user_id" json:"user_id"`
	Username  string    `ksql:"username" json:"username"`
	Password  string    `ksql:"password" json:"password"`
	Email     string    `ksql:"email" json:"email"`
	CreatedAt time.Time `ksql:"created_at,skipInserts,skipUpdates" json:"created_at"`
	UpdatedAt time.Time `ksql:"updated_at,skipInserts,skipUpdates" json:"updated_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("User ID: %d, Username: %s, Email: %s, Created At: %v\n",
		u.UserID, u.Username, u.Email, u.CreatedAt)
}
