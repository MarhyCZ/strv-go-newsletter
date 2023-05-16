package query

import _ "embed"

var (
	//go:embed scripts/ListUsers.sql
	ListUsers string
)
