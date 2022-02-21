package models

import (
	"database/sql"
	"os"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type BookService interface {
	BookDB
}

type BookDB interface {
	GetBooks() ([]Book, error)
	GetBook(id int) (*Book, error)
	AddBook(book *Book) (int64, error)
	UpdateBook(book *Book) (int64, error)
	RemoveBook(id int) (int64, error)
}

func NewBookService(db *sql.DB) BookService {
	return &bookService{
		BookDB: &bookDB{db: db},
	}
}

type bookService struct {
	BookDB
}

var _ BookDB = &bookDB{}

type bookDB struct {
	db *sql.DB
}

func (b bookDB) GetBooks() ([]Book, error) {
	var books []Book

	scheme := os.Getenv("DB_SCHEME")
	rows, err := b.db.Query("SELECT * FROM $1", scheme)
	if err != nil {
		return []Book{}, err
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return []Book{}, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b bookDB) GetBook(id int) (*Book, error) {
	var book *Book

	scheme := os.Getenv("DB_SCHEME")
	rows := b.db.QueryRow("SELECT * FROM $1 WHERE id=$2", scheme, id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b bookDB) AddBook(book *Book) (int64, error) {
	scheme := os.Getenv("DB_SCHEME")
	err := b.db.QueryRow("INSERT INTO $1 (title, author, year) VALUES($2, $3, $4) RETURNING id;", scheme, book.Title, book.Author, book.Year).Scan(&book.ID)
	if err != nil {
		return 0, err
	}

	return int64(book.ID), nil
}

func (b bookDB) UpdateBook(book *Book) (int64, error) {
	scheme := os.Getenv("DB_SCHEME")
	result, err := b.db.Exec("UPDATE $1 SET title=$2, author=$3, year=$4 WHERE id=$5 RETURNING id", scheme, &book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b bookDB) RemoveBook(id int) (int64, error) {
	scheme := os.Getenv("DB_SCHEME")
	result, err := b.db.Exec("DELETE FROM $1 WHERE id = $2", scheme, id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
