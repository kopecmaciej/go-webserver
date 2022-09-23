package main

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/lib"
	"go-web-server/pkg/models"
	"go-web-server/pkg/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.LoadConfig()

	lib.Open()
	lib.AutoMigrate(&models.User{})
	lib.InitRedis()

	router := mux.NewRouter()
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	handler := cors.AllowAll().Handler(router)

	port := config.GlobalConfig.Server.Port
	if len(port) < 1 {
		port = "4000"
	}
	fmt.Println(port)
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
