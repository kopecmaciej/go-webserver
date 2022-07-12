package main

import (
	"fmt"
	"go-web-server/lib"
	"go-web-server/pkg/models"
	"go-web-server/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	lib.Open()
	lib.AutoMigrate(&models.User{})

	router := mux.NewRouter()

	routes.UserRoutes(router)

	handler := cors.AllowAll().Handler(router)

	fmt.Println("Server listen on port 4000")
	panic(http.ListenAndServe(":4000", handler))
}
