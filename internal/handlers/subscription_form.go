package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/setxpro/subscribe-apex/internal/controllers"
)

func SubscritptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Erro ao fazer parse do form: %v", err)
			return
		}

		var s controllers.Subscription

		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao fazer parse do body: %v", err), http.StatusBadRequest)
			return
		}

		err := controllers.CreateSubscription(s.Name, s.Email)

		if err != nil {
			fmt.Fprintf(w, "Erro ao fazer parse do form: %v", err)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
