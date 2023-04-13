// Filename: cmd/web/data.go
package main

import (
	"net/url"

	"github.com/abelwhite/quotes/internal/models"
)

type templateData struct {
	Quote          []*models.Quote
	ErrorsFromForm map[string]string
	FormData       url.Values
	// Flash    string //flash is the key
}
