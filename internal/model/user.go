package model

import "time"
type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}


