package repository

import (
	"database/sql"
	"fmt"

	"github.com/SenselessA/CRUD_books"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Book struct {
	Id    int `db:"book_id"`
	Title string
	Isbm  string
}

type BooksRepository struct {
	db *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (r *BooksRepository) GetAllBooks() ([]Book, error) {
	var result []Book

	query := fmt.Sprintf("SELECT book_id, title, isbm FROM %s", bookTable)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *BooksRepository) GetBook(id string) (Book, error) {
	var result Book
	query := fmt.Sprintf("SELECT book_id, title, isbm FROM %s WHERE book_id = %s", bookTable, id)

	err := r.db.Get(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *BooksRepository) AddBook(newBook CRUD_books.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, isbm) VALUES ($1, $2) RETURNING book_id", bookTable)

	row := r.db.QueryRow(query, newBook.Title, newBook.Isbm)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BooksRepository) UpdateBook(book CRUD_books.Book) (Book, error) {
	var result Book
	query := fmt.Sprintf("UPDATE %s SET (title, isbm) = ($1, $2) WHERE book_id = %d RETURNING *", bookTable, book.Id)

	err := r.db.Get(&result, query, book.Title, book.Isbm)
	if err != nil {
		logrus.Println(err)
		return result, err
	}

	return result, nil
}

func (r *BooksRepository) DeleteBook(id string) (Book, error) {
	var result Book

	query := fmt.Sprintf("DELETE FROM %s WHERE book_id = %s RETURNING *", bookTable, id)

	err := r.db.Get(&result, query)
	// if all ok, err = sql: no rows in result set ?????????
	if err != nil || err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}
