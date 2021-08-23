package main

import (
	"app/Go-study/web/WEB8/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3001", myapp.NewHandler())
}
