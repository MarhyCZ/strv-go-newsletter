package query

import _ "embed"

var (
	//go:embed scripts/ListEditors.sql
	ListEditors string

	//go:embed scripts/CreateEditor.sql
	CreateEditor string
)
