package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/schema"
)

var Decoder = schema.NewDecoder()

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidData() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid request data"))
}

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJson(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				writeJson(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err, "path", r.URL.Path)
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func handleComponent(component func() templ.Component) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return Render(w, r, component())
	}
}
