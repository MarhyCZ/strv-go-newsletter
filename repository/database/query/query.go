package query

import _ "embed"

var (
	//go:embed scripts/ListEditors.sql
	ListEditors string
	//go:embed scripts/CreateEditor.sql
	CreateEditor string
	//go:embed scripts/GetEditor.sql
	GetEditor string

	//go:embed scripts/CreateNewsletter.sql
	CreateNewsletter string
	//go:embed scripts/RenameNewsletter.sql
	RenameNewsletter string
	//go:embed scripts/DeleteNewsletter.sql
	DeleteNewsletter string
	//go:embed scripts/DeleteNewsletter.sql
	ListNewsletters string
)
