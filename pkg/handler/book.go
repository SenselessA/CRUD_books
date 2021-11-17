package handler

import (
	"net/http"

	"github.com/SenselessA/CRUD_books"
	"github.com/gin-gonic/gin"
)

// @Summary All Books
// @Tags books
// @Description get all books
// @ID get-all-books
// @Accept  json
// @Produce  json
// @Success 200 {object} []repository.Book
// @Router /books [get]
func (h *Handler) GetAllBooks(c *gin.Context) {
	books, err := h.services.GetAllBooks()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary Get Book
// @Tags book
// @Description get book by ID
// @ID get-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} []repository.Book
// @Router /book [get]
func (h *Handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := h.services.GetBook(id)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary Create Book
// @Tags book
// @Description create book
// @ID create-book
// @Accept json
// @Produce json
// @Param input body CRUD_books.Book true "book info"
// @Success 200 {object} CRUD_books.BookId
// @Router /book [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var book CRUD_books.Book

	if err := c.BindJSON(&book); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.services.AddBook(book)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary Update Book
// @Tags book
// @Description create book
// @ID update-book
// @Accept json
// @Produce json
// @Param input body CRUD_books.Book true "book info"
// @Success 200 {object} repository.Book
// @Router /book [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	var book CRUD_books.Book

	if err := c.BindJSON(&book); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedBook, err := h.services.UpdateBook(book)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// @Summary Delete Book
// @Tags book
// @Description delete book
// @ID delete-book
// @Accept  json
// @Produce  json
// @Param input body CRUD_books.BookId true "book id"
// @Success 200 {object} repository.Book
// @Router /book [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	var bookId CRUD_books.BookId

	if err := c.BindJSON(&bookId); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	deletedBook, err := h.services.DeleteBook(bookId.Id)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, deletedBook)
}
