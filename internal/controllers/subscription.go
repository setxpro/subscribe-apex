package controllers

import "github.com/setxpro/subscribe-apex/internal/db"

type Subscription struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateSubscription(name string, email string) error {

	subscription := Subscription{Name: name, Email: email}

	return db.Insert("subscription", subscription)
}

func FindAllSubscriptions() ([]Subscription, error) {

	var subscriptions []Subscription

	err := db.FindAll("subscription", &subscriptions)

	return subscriptions, err
}
