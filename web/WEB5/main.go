package main

import (
	"app/Go-study/web/WEB5/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
