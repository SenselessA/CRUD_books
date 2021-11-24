package CRUD_books

import (
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("user with such credentials not found")

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
}