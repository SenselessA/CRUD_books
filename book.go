package CRUD_books

type Book struct {
	Id    int    `json: "-"`
	Title string `json: "title"`
	Isbm  string `json: "isbm"`
}
