package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
	sendEmail "github.com/marhycz/strv-go-newsletter/emails"
	"github.com/marhycz/strv-go-newsletter/repository/store"
)

func (rest *Rest) routeSubscriptions(r *chi.Mux) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", rest.listOfSubscriptions)

		r.Get("/{newsletter_id}/{email}", rest.getSubscription)

		r.Post("/", rest.subscribe)

		r.Get("/unsubscribe/{subscription_id}", rest.unsubscribe)
	})
}

func (rest *Rest) listOfSubscriptions(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	subscriptions := rest.env.Store.GetSubscriptions(ctx)
	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (rest *Rest) getSubscription(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id, error := strconv.Atoi(chi.URLParam(r, "newsletter_id"))

	if error != nil {
		fmt.Println("Error during conversion")
		return
	}
	subscriptions := rest.env.Store.GetSubscription(ctx, id, chi.URLParam(r, "email"))
	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) subscribe(w http.ResponseWriter, r *http.Request) {

	var sub store.Subscription

	json.NewDecoder(r.Body).Decode(&sub)
	ctx := r.Context()

	body := []byte(`# Welcome aboard!  
	---
	
	You have successfully subscribed to our newsletter,  
	you can unsubscribe any time with the link  
	at the bottom of each e-mail we send you.
	
	**Have a great day!**`)

	subscriptions, subErr := rest.env.Store.NewSubscription(ctx, sub.Newsletter_id, sub.Email)

	pattern := regexp.MustCompile("(.*)@")
	subName := pattern.FindString(sub.Email)

	sendEmail.SendNewEmail(subName, sub.Email, "Newsletter: Successfully subscribed to newsletter", body, subscriptions)

	err := json.NewEncoder(w).Encode(subErr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) unsubscribe(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "subscription_id")
	ctx := r.Context()

	subscriptions := rest.env.Store.DeleteSubscription(ctx, id)

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
