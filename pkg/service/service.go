package service

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/SenselessA/CRUD_books/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Book interface {
	GetAllBooks() ([]repository.Book, error)
	GetBook(id string) (repository.Book, error)
	AddBook(newBook CRUD_books.Book) (int, error)
	UpdateBook(book CRUD_books.Book) (repository.Book, error)
	DeleteBook(id string) (repository.Book, error)
}

type User interface {
	SignUp(c *gin.Context, inp CRUD_books.SignUpInput) error
	SignIn(c *gin.Context, inp CRUD_books.SignInInput) (string, error)
	ParseToken(c *gin.Context, token string) (int64, error)
}

type BooksService struct {
	Book
}

type UsersService struct {
	User
}

func NewBooks(repos *repository.BooksRepo) *BooksService {
	return &BooksService{
		Book: NewBookService(repos.Books),
	}
}
