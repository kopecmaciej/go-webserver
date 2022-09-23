package routes

import (
	"encoding/json"
	"fmt"
	"go-web-server/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

var AuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", Login).Methods("POST")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var auth models.Authorization
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(auth)
	defer r.Body.Close()
	user, err := auth.GetValidUser()
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	token := auth.CreateSession(user.Id)
	fmt.Println(token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
