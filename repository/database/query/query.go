package query

import _ "embed"

var (
	//go:embed scripts/ListEditors.sql
	ListEditors string

	//go:embed scripts/CreateEditor.sql
	CreateEditor string

	//go:embed scripts/getEditor.sql
	GetEditor string
)
