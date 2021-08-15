package main

import (
	"app/Go-study/Web2/TODO/app"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3002", n)

}
