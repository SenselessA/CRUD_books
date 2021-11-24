package handler

import (
	"net/http"
	"strconv"

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
	books, err := h.booksService.GetAllBooks()
	if err != nil {
		c.Status(http.StatusInternalServerError)
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

	book, err := h.booksService.GetBook(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
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
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := h.booksService.AddBook(book)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary Update Book
// @Tags book
// @Description update book
// @ID update-book
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param input body CRUD_books.UpdateBook true "book info"
// @Success 200 {object} repository.Book
// @Router /book/{id} [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.Status(http.StatusBadRequest)

		return
	}

	stringId, err := strconv.Atoi(idParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var bookNewInfo CRUD_books.UpdateBook

	if err := c.BindJSON(&bookNewInfo); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	updatedBook, err := h.booksService.UpdateBook(CRUD_books.Book{
		Id:    stringId,
		Title: bookNewInfo.Title,
		Isbm:  bookNewInfo.Isbm,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
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
// @Param id path int true "Book ID"
// @Success 200 {object} repository.Book
// @Router /book/{id} [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	deletedBook, err := h.booksService.DeleteBook(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, deletedBook)
}
