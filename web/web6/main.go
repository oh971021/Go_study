package main

import (
	"app/Go-study/web/web6/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3001", myapp.NewHandler())
}
