package repository

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Books interface {
	GetAllBooks() ([]Book, error)
	GetBook(id string) (Book, error)
	AddBook(newBook CRUD_books.Book) (int, error)
	UpdateBook(book CRUD_books.Book) (Book, error)
	DeleteBook(id string) (Book, error)
}

type BooksRepo struct {
	Books
}

func NewBooks(db *sqlx.DB) *BooksRepo {
	return &BooksRepo{
		Books: NewBooksPostgres(db),
	}
}
