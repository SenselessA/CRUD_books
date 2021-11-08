package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SenselessA/CRUD_books"
)

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	books, err := h.services.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.URL.Query().Get("id")
	book, err := h.services.GetBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book CRUD_books.Book

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.services.AddBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(map[string]interface{}{"id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book CRUD_books.Book

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedBook, err := h.services.UpdateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bookId struct {
		Id int
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &bookId); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deletedBook, err := h.services.DeleteBook(bookId.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(deletedBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
}
