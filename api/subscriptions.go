package api

import (
	"encoding/json"
	"errors"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	sendEmail "github.com/marhycz/strv-go-newsletter/emails"
	"github.com/marhycz/strv-go-newsletter/repository/store"
	"net/http"
	"regexp"
)

var (
	ErrAlreadySubscribed = errors.New("You are already subscribed to this newsletter")
)

func (rest *Rest) routeSubscriptions(r *chi.Mux) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.With(editorOnly).Get("/", rest.listOfSubscriptions)

		r.With(
			httpin.NewInput(getSubscriptionInput{}),
		).Get("/{newsletter_id}/{email}", rest.getSubscription)

		r.Post("/subscribe/{newsletter_id}", rest.subscribe)

		r.With(
			httpin.NewInput(unsubcribeInput{}),
		).Get("/unsubscribe/{subscription_id}", rest.unsubscribe)
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
	input := ctx.Value(httpin.Input).(*getSubscriptionInput)

	subscriptions := rest.env.Store.GetSubscription(ctx, input.NewsletterID, input.Email)
	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) subscribe(w http.ResponseWriter, r *http.Request) {

	var sub store.Subscription

	json.NewDecoder(r.Body).Decode(&sub)
	ctx := r.Context()

	currSubscriptions := rest.env.Store.GetSubscription(ctx, sub.Newsletter_id, sub.Email)
	if len(currSubscriptions) > 0 {
		w.Write([]byte(ErrAlreadySubscribed.Error()))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
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
	ctx := r.Context()
	input := ctx.Value(httpin.Input).(*unsubcribeInput)

	subscriptions := rest.env.Store.DeleteSubscription(ctx, input.SubscriptionID)

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
