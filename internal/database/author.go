package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	db   *sql.DB
	ID   string
	Name string
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

func (a *Author) Create(Name string) (*Author, error) {
	id := uuid.New().String()
	_, err := a.db.Exec("INSERT INTO author (id, name) VALUES ($1, $2)", id, Name)

	if err != nil {
		return &Author{}, err
	}

	return &Author{ID: id, Name: Name}, nil
}

func (a *Author) FindById(id string) (Author, error) {
	var author Author
	err := a.db.QueryRow("SELECT id, name FROM author WHERE id = $1", id).Scan(&author.ID, &author.Name)

	if err != nil {
		return Author{}, err
	}

	return author, nil
}

func (a *Author) FindAll() ([]Author, error) {
	var authors []Author
	rows, err := a.db.Query("SELECT id, name FROM author")

	if err != nil {
		return []Author{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var author Author
		err := rows.Scan(&author.ID, &author.Name)

		if err != nil {
			return []Author{}, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a *Author) FindByBookId(bookId string) (Author, error) {
	var author Author
	err := a.db.QueryRow("SELECT a.id, a.name FROM author a INNER JOIN book b ON b.author_id = a.id WHERE b.id = $1", bookId).Scan(&author.ID, &author.Name)

	if err != nil {
		return Author{}, err
	}

	return author, nil
}
