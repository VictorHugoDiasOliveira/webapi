package handlers

import (
	"encoding/json"
	"net/http"
	"webapi/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	persons, err := models.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}
