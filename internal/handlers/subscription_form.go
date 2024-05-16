package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/setxpro/subscribe-apex/internal/controllers"
	"github.com/setxpro/subscribe-apex/pkg/email"
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

		message := fmt.Sprintf("Parabéns %s, sua assinatura foi concluída. Em breve receberá um e-mail com o link de acesso.", s.Name)

		// Enviando o email
		if err := email.SendEmail(email.SentEmail{
			To:                   s.Email,
			From:                 "patrick.anjos@bagaggio.com.br",
			Html:                 message,
			Subject:              "Assinatura concluída",
			Message:              "",
			Base64Attachment:     "",
			Base64AttachmentName: "",
		}); err != nil {
			fmt.Printf("Erro ao enviar o email: %v\n", err)
		} else {
			fmt.Println("Email enviado com sucesso!")
		}

		if err != nil {
			fmt.Fprintf(w, "Erro ao fazer parse do form: %v", err)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func FindAllSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	subscriptions, err := controllers.FindAllSubscriptions()

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar assinaturas: %v", err), http.StatusInternalServerError)
		return
	}

	// cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Define o status HTTP 200 OK explicitamente
	w.WriteHeader(http.StatusOK)

	// Codifica a resposta em formato JSON e a envia
	if err := json.NewEncoder(w).Encode(subscriptions); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao codificar resposta JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
