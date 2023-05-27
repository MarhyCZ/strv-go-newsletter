package api

import (
	"encoding/json"
	"fmt"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	sendEmail "github.com/marhycz/strv-go-newsletter/emails"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func (rest *Rest) routeIssues(r *chi.Mux) {
	r.Route("/issues", func(r chi.Router) {
		r.Use(editorOnly)
		r.With(
			httpin.NewInput(listOfIssuesInput{}),
		).Get("/", rest.listOfIssues)
		r.With(
			httpin.NewInput(getIssueInput{}),
		).Get("/{newsletter_id}/{name}", rest.getIssue)
		r.With(
			httpin.NewInput(publishIssueInput{}),
		).Post("/", rest.publishIssue)
	})
}

func (rest *Rest) listOfIssues(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := ctx.Value(httpin.Input).(*listOfIssuesInput)
	newsletter := input.NewsletterID.String()

	if newsletter != "" {
		newsletter = newsletter + "/"
	}

	subscriptions, listErr := rest.env.Storage.GetIssuesList(ctx, os.Stdout, "/", newsletter)

	if listErr != nil {
		w.WriteHeader(http.StatusNoContent)
	}

	err := json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
func (rest *Rest) getIssue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := ctx.Value(httpin.Input).(*getIssueInput)
	newsletter := input.NewsletterID.String()
	path := newsletter + "/" + input.Name

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
	ctx := r.Context()
	db := rest.env.Database.Database

	input := ctx.Value(httpin.Input).(*publishIssueInput)
	newsletter := input.NewsletterID.String()
	email := ctx.Value("claims").(claims).Username
	path := newsletter + "/" + input.Name
	body, bodyErr := io.ReadAll(r.Body)

	if bodyErr != nil {
		log.Fatalln(bodyErr)
	}

	/* cE, editorCookieErr := r.Cookie("Editor")
	if editorCookieErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} */

	editor, EditErr := database.GetEditorByEmail(ctx, db, email)
	if EditErr != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	editorsNewsletters, newsErr := database.ListEditorNewsletters(ctx, db, editor.ID)
	if newsErr != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Println(editorsNewsletters)

	createIssue := rest.env.Storage.StreamFileUpload(ctx, os.Stdout, path, body)
	encErr := json.NewEncoder(w).Encode(createIssue)
	if encErr != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	it := rest.env.Store.GetNewsletterSubscriptions(ctx, input.NewsletterID)
	for _, element := range it {
		pattern := regexp.MustCompile("(.*)@")
		subName := pattern.FindString(element.Email)

		sendEmail.SendNewEmail(subName, element.Email, "Newsletter: "+input.Subject, body, element.Id)
	}

}
