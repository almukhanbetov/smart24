package main

import (
	"log"
	"os"

	"github.com/almukhanbetov/smart24/backend/internal/db"
	"github.com/almukhanbetov/smart24/backend/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	r := routes.Setup(pool)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("API started on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
