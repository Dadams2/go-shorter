package migrations

import "embed"

// Migrations contains the SQL migrations bundled with the application.
//
//go:embed *.sql
var Migrations embed.FS
