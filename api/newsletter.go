package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"net/http"
)

func (rest *Rest) routeNewsletter(r *chi.Mux) {
	// RESTy routes for "newsletters" resource
	r.Route("/newsletter", func(r chi.Router) {
		r.Post("/", rest.createNewsletter)
		r.Delete("/{newsletter_id}", rest.deleteNewsletter)
		r.Get("/", rest.listEditorNewsletters)
	})
}

func (rest *Rest) createNewsletter(w http.ResponseWriter, r *http.Request) {
	input := createNewsletterInput{}
	if err := parseRequestBody(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := rest.env.Database.Database
	newsletter, err := database.CreateNewsletter(r.Context(), db, input.EditorID, input.Name, input.Description)
	if err != nil {
		panic(err)
	}

	response := fmt.Sprintf("The newsletter with ID: %s was successfully created.", newsletter.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (rest *Rest) deleteNewsletter(w http.ResponseWriter, r *http.Request) {
	input := deleteNewsletterInput{}
	if err := parseRequestBody(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := rest.env.Database.Database
	err := database.DeleteNewsletter(r.Context(), db, input.ID)
	if err != nil {
		panic(err)
	}

	response := fmt.Sprintf("The newsletter with ID: %s was successfully deleted.", input.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (rest *Rest) listEditorNewsletters(w http.ResponseWriter, r *http.Request) {
	input := listEditorNewslettersInput{}
	if err := parseRequestBody(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := rest.env.Database.Database
	newsletters, err := database.ListEditorNewsletters(r.Context(), db, input.EditorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(newsletters)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

}
