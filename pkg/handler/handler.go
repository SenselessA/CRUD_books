package handler

import (
	"net/http"

	"github.com/SenselessA/CRUD_books/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", h.books)
	mux.HandleFunc("/books/", h.book)
	return mux
}

// books
func (h *Handler) books(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.GetAllBooks(w, r)
	}
}

// book/id
func (h *Handler) book(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetBook(w, r)
	case "POST":
		h.CreateBook(w, r)
	case "PUT":
		h.UpdateBook(w, r)
	case "DELETE":
		h.DeleteBook(w, r)
	}
}
