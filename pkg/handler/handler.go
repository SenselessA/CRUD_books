package handler

import (
	"github.com/SenselessA/CRUD_books/pkg/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SenselessA/CRUD_books/docs"
)

type Handler struct {
	Book *service.BooksService
	User *service.UsersService
}

func NewHandler(BooksService *service.BooksService, UsersService *service.UsersService) *Handler {
	return &Handler{
		Book: BooksService,
		User: UsersService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

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
		book.GET("/:id", h.GetBook)
		book.POST("/", h.CreateBook)
		book.PUT("/:id", h.UpdateBook)
		book.DELETE("/:id", h.DeleteBook)
	}

	return router
}
