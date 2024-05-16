package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/setxpro/subscribe-apex/internal/handlers"
)

func main() {

	// Carrega as variáveis
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL := os.Getenv("MONGODB_URL")
	fmt.Println(mongoURL)

	http.HandleFunc("POST /", handlers.SubscritptionHandler)
	http.HandleFunc("GET /", handlers.FindAllSubscriptionsHandler)

	if err := http.ListenAndServe(":"+os.Getenv("API_PORT"), nil); err != nil {
		log.Fatal(err)
	}

}
