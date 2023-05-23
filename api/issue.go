package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func (rest *Rest) routeIssues(r *chi.Mux) {
	r.Route("/issues", func(r chi.Router) {
		r.Get("/", rest.listOfIssues)

		r.Get("/issue", rest.getIssue)

		r.Post("/issue", rest.publishIssue)
	})
}

func (rest *Rest) listOfIssues(w http.ResponseWriter, r *http.Request) {
	newsletter := r.URL.Query().Get("newsletter_id")

	if newsletter != "" {
		newsletter = newsletter + "/"
	}

	ctx := r.Context()
	subscriptions := rest.env.Storage.GetIssuesList(ctx, os.Stdout, "/", newsletter)
	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) getIssue(w http.ResponseWriter, r *http.Request) {
	newsletter := r.URL.Query().Get("newsletter_id")
	name := r.URL.Query().Get("name")
	path := newsletter + "/" + name

	if newsletter == "" || name == "" {
		w.WriteHeader(http.StatusForbidden)
	}

	ctx := r.Context()
	subscriptions, downlFailure := rest.env.Storage.DownloadFileIntoMemory(ctx, os.Stdout, path)

	if downlFailure != nil {
		w.WriteHeader(http.StatusNoContent)
	}

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (rest *Rest) publishIssue(w http.ResponseWriter, r *http.Request) {
	newsletter := r.URL.Query().Get("newsletter_id")
	name := r.URL.Query().Get("name")
	path := newsletter + "/" + name
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		log.Fatalln(bodyErr)
	}
	data := string(body)

	if newsletter == "" || name == "" || data == "" {
		w.WriteHeader(http.StatusForbidden)
	}

	ctx := r.Context()
	subscriptions := rest.env.Storage.StreamFileUpload(ctx, os.Stdout, path, data)

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
