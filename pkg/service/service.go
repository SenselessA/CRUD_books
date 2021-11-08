package service

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/SenselessA/CRUD_books/pkg/repository"
)

type Book interface {
	GetAllBooks() ([]repository.Book, error)
	GetBook(id string) (repository.Book, error)
	AddBook(newBook CRUD_books.Book) (int, error)
	UpdateBook(book CRUD_books.Book) (repository.Book, error)
	DeleteBook(id int) (repository.Book, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Books),
	}
}
