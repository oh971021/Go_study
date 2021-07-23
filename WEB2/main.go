// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type fooHandler struct{}

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello foo!")
// }

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello %s!", name)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World")
// 	})

// 	mux.HandleFunc("/bar", barHandler)

// 	mux.Handle("/foo", &fooHandler{})

// 	http.ListenAndServe(":3000", mux)
// }

// // func second() {
// // 	// 라우터 클래스 생성, mux 등록

// // 	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// // 	// 	fmt.Fprint(w, "Hello bar")
// // 	// })

// // 	// http.Handle("/foo", &fooHandler{})

// // 	// port 는 3000번 , nil은 3000번을 찾으면 서버를 실행하라는 명령어
// // 	// ListenAndServe : 서버를 만든다.
// // }

package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	mux := http.NewServeMux() //Mux 라우터를 통한 구현! Mux 클래스를 만들어주라
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler) //mux

	mux.Handle("/foo", &fooHandler{}) //mux

	http.ListenAndServe(":3000", mux) // mux 라우트 등록
	//요 함수를 통해 웹서버 간단제작.
}
