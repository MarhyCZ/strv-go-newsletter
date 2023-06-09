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
	//go:embed scripts/ListNewsletters.sql
	ListNewsletters string
	//go:embed scripts/GetNewsletter.sql
	GetNewsletter string

	//go:embed scripts/CreatePasswordReset.sql
	CreatePasswordReset string
	//go:embed scripts/ResetPassword.sql
	ResetPassword string
	//go:embed scripts/GetPasswordReset.sql
	GetPasswordReset string
)
