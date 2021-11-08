package service

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/SenselessA/CRUD_books/pkg/repository"
)

type BookService struct {
	repo repository.Books
}

func NewBookService(repo repository.Books) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) GetAllBooks() ([]repository.Book, error) {
	return b.repo.GetAllBooks()
}

func (b *BookService) GetBook(id string) (repository.Book, error) {
	return b.repo.GetBook(id)
}

func (b *BookService) AddBook(book CRUD_books.Book) (int, error) {
	return b.repo.AddBook(book)
}

func (b *BookService) UpdateBook(book CRUD_books.Book) (repository.Book, error) {
	return b.repo.UpdateBook(book)
}

func (b *BookService) DeleteBook(id int) (repository.Book, error) {
	return b.repo.DeleteBook(id)
}