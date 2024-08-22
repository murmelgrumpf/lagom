//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.HandlerFunc {
	fs := http.FileServerFS(publicFS)
	return func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}
}
