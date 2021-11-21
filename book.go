package CRUD_books

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Isbm  string `json:"isbm"`
}

type BookId struct {
	Id    int    `json:"id"`
}

type UpdateBook struct {
	Title string `json:"title"`
	Isbm  string `json:"isbm"`
}
