package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marhycz/strv-go-newsletter/environment"
	"github.com/marhycz/strv-go-newsletter/repository/store"
)

type Rest struct {
	*chi.Mux
	env *environment.Env
}

func NewController(env *environment.Env) *Rest {
	c := &Rest{
		env: env,
	}
	c.initRouter()
	return c
}

func (rest *Rest) initRouter(){
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/subscriptions", func(r chi.Router){
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			subscriptions := rest.env.Store.GetSubscriptions(ctx)
			err := json.NewEncoder(w).Encode(subscriptions)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
		})
	
		r.Get("/{newsletter_id}/{email}", func(w http.ResponseWriter, r *http.Request) {
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
		})
	
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var sub store.Subscription
	
			json.NewDecoder(r.Body).Decode(&sub)
			ctx := r.Context()
	
			subscriptions := rest.env.Store.NewSubscription(ctx, sub.Newsletter_id, sub.Email)
	
			err := json.NewEncoder(w).Encode(subscriptions)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
		})
	
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			var sub store.Subscription
	
			json.NewDecoder(r.Body).Decode(&sub)
			ctx := r.Context()
	
			subscriptions := rest.env.Store.DeleteSubscription(ctx, sub.Id)
	
			err := json.NewEncoder(w).Encode(subscriptions)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
		})
	})

	rest.routeEditor(r)
	rest.Mux = r
}
