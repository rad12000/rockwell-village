package api

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Auth!: %s</h1>", r.RequestURI)
}
