package main

import (
	"app/Go-study/web/WEB7/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3001", myapp.NewHandler())
}
