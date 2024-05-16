package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SentEmail struct {
	To                   string `json:"to"`
	From                 string `json:"from"`
	Html                 string `json:"html"`
	Subject              string `json:"subject"`
	Base64Attachment     string `json:"base64Attachment"`
	Base64AttachmentName string `json:"base64AttachmentName"`
	Message              string `json:"message"`
}

func NewSentEmail(to string, from string, html string, subject string, base64Attachment string, base64AttachmentName string, message string) SentEmail {

	return SentEmail{
		To:                   to,
		From:                 from,
		Html:                 html,
		Subject:              subject,
		Base64Attachment:     base64Attachment,
		Base64AttachmentName: base64AttachmentName,
		Message:              message,
	}
}

func SendEmail(sentEmail SentEmail) error {

	url := os.Getenv("EMAIL_API_URL")

	// Codifica a estrutura SentEmail para JSON
	payload, err := json.Marshal(sentEmail)

	if err != nil {
		return fmt.Errorf("erro ao codificar para JSON: %v", err)
	}

	// Requisição POST com os dados JSON
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return fmt.Errorf("erro ao fazer a requisição POST: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("falha na requisição: status %d", resp.StatusCode)
	}

	return nil
}
