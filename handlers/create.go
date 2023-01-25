package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapi/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(person)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Person inserida com sucesso. ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
