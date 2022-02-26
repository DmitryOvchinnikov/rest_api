package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dmitryovchinnikov/rest_api/models"
)

func SendError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(error)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	fmt.Println(data)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
