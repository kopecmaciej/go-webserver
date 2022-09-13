package main

import (
	"fmt"
	"go-web-server/lib"
	"go-web-server/pkg/models"
	"go-web-server/pkg/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	GoServerAddr = os.Getenv("SERVER_ADDR")
)

func main() {
	lib.Open()
	lib.AutoMigrate(&models.User{})
	lib.InitRedis()

	router := mux.NewRouter()
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	handler := cors.AllowAll().Handler(router)

	var port string = GoServerAddr
	if len(port) < 1 {
		port = "4000"
	}
	fmt.Println("Server listen on port: " + port)

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start on port %s ,err: %v", port, err)
	}
}
