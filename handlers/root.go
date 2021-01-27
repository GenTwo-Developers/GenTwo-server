package handlers

import (
	"encoding/json"
	"net/http"
)

type Root struct {
	Status  int
	Message string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	root := Root{http.StatusOK, "Hi there!"}
	w.Header().Set("Content-Type", "Application/json")
	response, err := json.Marshal(root)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
