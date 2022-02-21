package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/dmitryovchinnikov/rest_api/models"
	"github.com/dmitryovchinnikov/rest_api/utils"
	"github.com/gorilla/mux"
)

func NewBooks(s models.BookService, r *mux.Router) *Books {
	return &Books{
		s,
		r,
	}
}

type Books struct {
	service models.BookService
	router  *mux.Router
}

func (b *Books) GetBooks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var errDB models.Error

		books, err := b.service.GetBooks()
		if err != nil {
			errDB.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, errDB)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (b *Books) GetBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var errDB models.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errDB.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, errDB)
			return
		}

		book, err := b.service.GetBook(id)
		if err != nil {
			errDB.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, errDB)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, &book)
	}
}

func (b *Books) AddBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var errDB models.Error

		var book *models.Book
		err := json.NewDecoder(r.Body).Decode(book)
		if err != nil {
			log.Fatal(err)
		}

		if book.Author == "" || book.Title == "" || book.Year == "" {
			errDB.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, errDB)
			return
		}

		id, err := b.service.AddBook(book)
		if err != nil {
			errDB.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, errDB)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, id)
	}
}

func (b *Books) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errDB models.Error

		var book *models.Book
		err := json.NewDecoder(r.Body).Decode(book)
		if err != nil {
			log.Fatal(err)
		}

		if book.Author == "" || book.Title == "" || book.Year == "" {
			errDB.Message = "Enter all fields."
			utils.SendError(w, http.StatusBadRequest, errDB)
			return
		}

		id, err := b.service.UpdateBook(book)
		spew.Dump(err)
		if err != nil {
			errDB.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, errDB)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, id)
	}
}

func (b *Books) RemoveBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errDB models.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}

		idRemoved, err := b.service.RemoveBook(id)
		if err != nil {
			errDB.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, errDB)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, idRemoved)
	}
}
