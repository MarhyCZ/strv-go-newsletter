package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marhycz/strv-go-newsletter/repository/store"
)

func (rest *Rest) routeSubscriptions(r *chi.Mux) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", rest.listOfSubscriptions)

		r.Get("/{newsletter_id}/{email}", rest.getSubscription)

		r.Post("/", rest.subscribe)

		r.Delete("/", rest.unsubscribe)
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

	subscriptions := rest.env.Store.NewSubscription(ctx, sub.Newsletter_id, sub.Email)

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) unsubscribe(w http.ResponseWriter, r *http.Request) {

	var sub store.Subscription

	json.NewDecoder(r.Body).Decode(&sub)
	ctx := r.Context()

	subscriptions := rest.env.Store.DeleteSubscription(ctx, sub.Id)

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
