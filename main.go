package main

import (
	"log"
	"net/http"

	"github.com/dmitryovchinnikov/rest_api/config"
	"github.com/dmitryovchinnikov/rest_api/controller"
	"github.com/dmitryovchinnikov/rest_api/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := config.Load()
	dbCfg := cfg.Database
	service, err := models.NewService(
		models.WithDB(dbCfg.Dialect(), dbCfg.DSN()),
		models.WithBook(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(s *models.Service) {
		err := s.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(service)

	r := mux.NewRouter()
	c := controller.NewBooks(service.Book, r)

	r.HandleFunc("/books", c.GetBooks()).Methods("GET")
	r.HandleFunc("/books/{id}", c.GetBook()).Methods("GET")
	r.HandleFunc("/books", c.AddBook()).Methods("POST")
	r.HandleFunc("/books", c.UpdateBook()).Methods("POST")
	r.HandleFunc("/books/{id}", c.RemoveBook()).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", &controller.WithCORS{Router: r}))
}
