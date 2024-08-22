package viewhandler

import (
	"lagom/handlers"
	"lagom/views/home"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return handlers.Render(w, r, home.Index())
}
