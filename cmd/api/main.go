package main

import (
	"log"
	"net/http"

	"github.com/setxpro/subscribe-apex/internal/handlers"
)

func main() {

	http.HandleFunc("/", handlers.SubscritptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
