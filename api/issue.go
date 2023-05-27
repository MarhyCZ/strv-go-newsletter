package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
	sendEmail "github.com/marhycz/strv-go-newsletter/emails"
	"github.com/marhycz/strv-go-newsletter/repository/database"
)

func (rest *Rest) routeIssues(r *chi.Mux) {
	r.Route("/issues", func(r chi.Router) {
		r.Use(editorOnly)
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
	ctx := r.Context()
	db := rest.env.Database.Database

	newsletter := r.URL.Query().Get("newsletter_id")
	name := r.URL.Query().Get("name")
	subject := r.URL.Query().Get("subject")
	email := ctx.Value("claims").(claims).Username
	path := newsletter + "/" + name
	body, bodyErr := io.ReadAll(r.Body)

	if bodyErr != nil {
		log.Fatalln(bodyErr)
	}

	if newsletter == "" || name == "" {
		w.WriteHeader(http.StatusForbidden)
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

	newsletterInt, errConv := strconv.Atoi(newsletter)

	if errConv != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	it := rest.env.Store.GetNewsletterSubscriptions(ctx, newsletterInt)
	for _, element := range it {
		pattern := regexp.MustCompile("(.*)@")
		subName := pattern.FindString(element.Email)

		sendEmail.SendNewEmail(subName, element.Email, "Newsletter: "+subject, body, element.Id)
	}

}
