package main

import (
	"fmt"
	"net/http"
)

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "반갑다")
}

type fooHandler struct{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // "/" : 절대경로 지정 (인덱스링크)
		fmt.Fprint(w, "Hello")
	})

	http.HandleFunc("/Junseok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "world")
	})

	http.Handle("/Junseok", &fooHandler{})
	// port 는 3000번 , nil은 3000번을 찾으면 서버를 실행하라는 명령어
	http.ListenAndServe(":3000", nil)
	// ListenAndServe : 서버를 만든다.
}
