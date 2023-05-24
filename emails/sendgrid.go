package sendEmail

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	htmlTemplate "github.com/marhycz/strv-go-newsletter/emails/templates"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendEmail struct {
	client *sendgrid.Client
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func SendNewEmail(receiverName, receiverAddress, subject string, md []byte, subId string) *SendEmail {
	from := mail.NewEmail("VSE GO project 2023", "vse.goproject2023@gmail.com")
	to := mail.NewEmail(receiverName, receiverAddress)

	html := mdToHTML(md)
	str := string(html)

	email := htmlTemplate.ComposeEmail(str, "localhost:8080/subscriptions/unsubscribe/"+subId)
	plainTextContent := str

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, email)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	fb := &SendEmail{
		client: client,
	}

	return fb
}
