package main

import (
	"log"
	"net/http"

	"github.com/kskr24/workspacehub/internal/config"
	"github.com/kskr24/workspacehub/internal/db"
	"github.com/kskr24/workspacehub/internal/routes"
)

func main() {
	config.LoadEnv()

	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	defer db.DB.Close()

	r := routes.RegisterRoutes()
	log.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
