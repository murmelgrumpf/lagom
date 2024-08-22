//+build dev
//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.HandlerFunc {
	fmt.Println("building static files for development")
	fs := http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("max-age", "0")
		fs.ServeHTTP(w, r)
	}
}
