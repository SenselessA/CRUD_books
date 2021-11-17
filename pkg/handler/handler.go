package handler

import (
	"github.com/SenselessA/CRUD_books/pkg/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SenselessA/CRUD_books/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	books := router.Group("books", h.GetAllBooks)
	{
		books.GET("")
	}

	book := router.Group("book")
	{
		book.GET("/:id", h.GetBook)
		book.POST("/", h.CreateBook)
		book.PUT("/", h.UpdateBook)
		book.DELETE("/", h.DeleteBook)
	}

	return router
}
