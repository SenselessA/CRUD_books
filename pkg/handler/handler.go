package handler

import (
	"github.com/SenselessA/CRUD_books"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SenselessA/CRUD_books/docs"
	"github.com/SenselessA/CRUD_books/pkg/repository"
)

type Books interface {
	AddBook(book CRUD_books.Book) (int, error)
	GetBook(id string) (repository.Book, error)
	GetAllBooks() ([]repository.Book, error)
	DeleteBook(id string) (repository.Book, error)
	UpdateBook(book CRUD_books.Book) (repository.Book, error)
}

type User interface {
	SignUp(c *gin.Context, inp CRUD_books.SignUpInput) error
	SignIn(c *gin.Context, inp CRUD_books.SignInInput) (string, error)
	ParseToken(c *gin.Context, token string) (int64, error)
}

type Handler struct {
	booksService Books
	usersService User
}

func NewHandler(books Books, users User) *Handler {
	return &Handler{
		booksService: books,
		usersService: users,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(loggingMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.GET("/sign-in", h.SignIn)
	}

	books := router.Group("books", h.GetAllBooks)
	{
		books.GET("")
	}

	book := router.Group("book")
	{
		books.Use(h.authMiddleware())

		book.GET("/:id", h.GetBook)
		book.POST("/", h.CreateBook)
		book.PUT("/:id", h.UpdateBook)
		book.DELETE("/:id", h.DeleteBook)
	}

	return router
}
