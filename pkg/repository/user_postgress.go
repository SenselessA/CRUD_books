package repository

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{db}
}

func (r *Users) Create(c *gin.Context, user CRUD_books.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password, registered_at) values ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.RegisteredAt)

	return err
}

func (r *Users) GetByCredentials(c *gin.Context, email, password string) (CRUD_books.User, error) {
	var user CRUD_books.User
	err := r.db.QueryRow("SELECT id, name, email, registered_at FROM users WHERE email=$1 AND password=$2", email, password).
		Scan(&user.ID, &user.Name, &user.Email, &user.RegisteredAt)

	return user, err
}
