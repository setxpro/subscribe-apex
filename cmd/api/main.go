package main

import (
	"log"
	"net/http"

	"github.com/setxpro/subscribe-apex/internal/handlers"
)

func main() {

	http.HandleFunc("POST /", handlers.SubscritptionHandler)
	http.HandleFunc("GET /", handlers.FindAllSubscriptionsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
