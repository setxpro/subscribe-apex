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

	http.HandleFunc("POST /", handlers.SubscritptionHandler)
	http.HandleFunc("GET /", handlers.FindAllSubscriptionsHandler)

	fmt.Printf("Servidor iniciado em http://localhost:%s\n", os.Getenv("API_PORT"))

	if err := http.ListenAndServe(":"+os.Getenv("API_PORT"), nil); err != nil {
		log.Fatal(err)
	}

}
