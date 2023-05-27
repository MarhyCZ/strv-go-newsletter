package api

import (
	"encoding/json"
	"fmt"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"net/http"
)

func (rest *Rest) routeNewsletter(r *chi.Mux) {
	// RESTy routes for "newsletters" resource
	r.Route("/newsletter", func(r chi.Router) {
		r.Use(editorOnly)
		r.Post("/", rest.createNewsletter)
		r.Get("/", rest.listEditorNewsletters)
		r.With(
			httpin.NewInput(deleteNewsletterInput{}),
		).Delete("/{newsletter_id}", rest.deleteNewsletter)
		r.Put("/{newsletter_id}", rest.renameNewsletter)
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
	ctx := r.Context()
	input := ctx.Value(httpin.Input).(*deleteNewsletterInput)
	editorID := ctx.Value("claims").(claims).EditorID

	db := rest.env.Database.Database
	newsletter, err := database.GetNewsletter(ctx, db, input.NewsletterID)

	if newsletter == nil {
		response := fmt.Sprintf("The newsletter with ID: %s does not exist.", input.NewsletterID)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(response))
		return
	}
	if newsletter.EditorID != editorID || err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not found in your account.", newsletter.ID)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(response))
		return
	}

	err = database.DeleteNewsletter(ctx, db, newsletter.ID)
	if err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not deleted.", newsletter.ID)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
		return
	}

	response := fmt.Sprintf("The newsletter with ID: %s was successfully deleted.", newsletter.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (rest *Rest) listEditorNewsletters(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	editorID := ctx.Value("claims").(claims).EditorID
	db := rest.env.Database.Database
	newsletters, err := database.ListEditorNewsletters(ctx, db, editorID)
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

	newsletterID, err := uuid.Parse(chi.URLParam(r, "newsletter_id"))

	if err != nil {
		response := fmt.Sprintf("Wrong newsletter ID in URL")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(response))
		return
	}

	ctx := r.Context()
	editorID := ctx.Value("claims").(claims).EditorID

	db := rest.env.Database.Database
	newsletter, err := database.GetNewsletter(ctx, db, newsletterID)
	if newsletter.EditorID != editorID || err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not found in your account.", newsletter.ID)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(response))
		return
	}

	err = database.RenameNewsletter(ctx, db, newsletterID, input.Name)
	if err != nil {
		response := fmt.Sprintf("The newsletter with ID: %s was not renamed.", newsletter.ID)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
		return
	}

	response := fmt.Sprintf("The newsletter with ID: %s was successfully renamed.", newsletterID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
