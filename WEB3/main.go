package main

import (
	"net/http"
	"web/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
	// myapp 에서 handler 를 만든다.
}
