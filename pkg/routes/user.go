package routes

import (
	"encoding/json"
	"fmt"
	"go-web-server/pkg/middleware"
	"go-web-server/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/user", middleware.Auth(GetAllUsers)).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUser{}
	var maxBytesPerBody int64 = 524228
	r.Body = http.MaxBytesReader(w, r.Body, maxBytesPerBody)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user, err := newUser.CreateUser()
	if err != nil {
		http.Error(w, "Error while saving user to database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringId := vars["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		fmt.Println(err)
	}
	user := models.User{Id: uint(id)}
	u, err := user.GetUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
    return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	stringId := vars["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		fmt.Println(err)
	}
	err = user.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}
