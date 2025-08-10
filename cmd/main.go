package main

import (
	"github.com/fredianto2405/catetin-api/config"
	"github.com/fredianto2405/catetin-api/internal/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db & setup routes
	db := config.NewDB()
	r := router.SetupRouter(db)

	// run app
	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port)
}
