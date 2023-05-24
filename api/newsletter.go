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
		r.Use(editorOnly)
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
	ctx := r.Context()
	editorID := ctx.Value("claims").(claims).EditorID

	db := rest.env.Database.Database
	newsletter, err := database.CreateNewsletter(r.Context(), db, editorID, input.Name, input.Description)
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

	ctx := r.Context()
	editorID := ctx.Value("claims").(claims).EditorID

	db := rest.env.Database.Database
	newsletter, err := database.GetNewsletter(ctx, db, input.ID)
	if newsletter.EditorID != editorID {
		response := fmt.Sprintf("The newsletter with ID: %s was not found in your account.", newsletter.ID)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(response))
		return
	}

	err = database.DeleteNewsletter(ctx, db, input.ID)
	if err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not deleted.", newsletter.ID)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
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

func (rest *Rest) renameNewsletter(w http.ResponseWriter, r *http.Request) {
	input := renameNewsletterInput{}
	if err := parseRequestBody(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	editorID := ctx.Value("claims").(claims).EditorID

	db := rest.env.Database.Database
	newsletter, err := database.GetNewsletter(ctx, db, input.ID)
	if newsletter.EditorID != editorID {
		response := fmt.Sprintf("The newsletter with ID: %s was not found in your account.", newsletter.ID)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(response))
		return
	}

	err = database.RenameNewsletter(ctx, db, input.ID, input.Name)
	if err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not renamed.", newsletter.ID)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
	}

	response := fmt.Sprintf("The newsletter with ID: %s was successfully renamed.", input.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
