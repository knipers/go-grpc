package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Book struct {
	db       *sql.DB
	ID       string
	Title    string
	Type     string
	AuthorID string
}

func NewBook(db *sql.DB) *Book {
	return &Book{db: db}
}

func (b *Book) Create(Title string, Type string, AuthorID string) (*Book, error) {
	id := uuid.New().String()
	_, err := b.db.Exec("INSERT INTO book (id, title, type, author_id) VALUES ($1, $2, $3, $4)", id, Title, Type, AuthorID)

	if err != nil {
		return &Book{}, err
	}

	return &Book{ID: id, Title: Title, Type: Type, AuthorID: AuthorID}, nil
}

func (b *Book) FindById(id string) (Book, error) {
	var book Book
	err := b.db.QueryRow("SELECT id, title, type, author_id FROM book WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Type, &book.AuthorID)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (b *Book) FindAll() ([]Book, error) {
	var books []Book
	rows, err := b.db.Query("SELECT id, title, type, author_id FROM book")

	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Type, &book.AuthorID)

		if err != nil {
			return []Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b *Book) FindByAuthorID(AuthorID string) ([]Book, error) {
	var books []Book
	rows, err := b.db.Query("SELECT id, title, type, author_id FROM book WHERE author_id = $1", AuthorID)

	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Type, &book.AuthorID)

		if err != nil {
			return []Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}
