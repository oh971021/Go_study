package myapp

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponsWriter, r *http.Request) {
	fmt.Fprint("Hello")
}

// Make a new myapp handlerg
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	return mux
}
