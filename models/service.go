package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

type ServiceConfig func(*Service) error

func WithDB(name, dsn string) ServiceConfig {
	return func(s *Service) error {
		pgUrl, err := pq.ParseURL(os.Getenv("DB_HOST"))
		if err != nil {
			log.Fatal(err)
		}

		db, err := sql.Open(name, pgUrl)
		if err != nil {
			log.Fatal(err)
		}
		s.db = db
		return nil
	}
}

func WithBook() ServiceConfig {
	return func(s *Service) error {
		s.Book = NewBookService(s.db)
		return nil
	}
}

func NewService(cfgs ...ServiceConfig) (*Service, error) {
	var s Service
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}

	return &s, nil
}

type Service struct {
	Book BookService
	db   *sql.DB
}

func (s *Service) Close() error {
	return s.db.Close()
}
