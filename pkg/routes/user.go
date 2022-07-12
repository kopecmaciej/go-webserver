package routes

import (
	"encoding/json"
	"fmt"
	"go-web-server/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", GetAllUsers).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUser{}
	r.Body = http.MaxBytesReader(w, r.Body, 524228)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
  user, err := newUser.CreateUser()
	if err != nil {
		http.Error(w, "Error while saving user to database", http.StatusInternalServerError)
		return
	}

	jsonUsers, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUsers)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	vars := mux.Vars(r)
	stringId := vars["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		fmt.Println(err)
	}
	u, err := user.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	jsonUser, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonUsers, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUsers)
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
