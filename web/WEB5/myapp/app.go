package myapp

import (
	"fmt"
	"net/http"

	// mux 는 경로를 타겟팅해서 분배한다.
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprint(w, "Hello World") // Fprint은 w값에 "Hello World" 값을 주어서 프린팅해라는 의미
}

func usersHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprint(w, "Get UserInfo by /Users/{id}") // Fprint은 w값에 "Hello World" 값을 주어서 프린팅해라는 의미
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                   //r은 request vars가 Id 파싱을 자동으로 해준다
	fmt.Fprint(w, "User id:", vars["id"]) // vars[id] 형식으로 작성해야 파싱값이 User Id:로 들어가 w로 출력된다
}

// func selfHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Self")
// }

func NewHandler() http.Handler {
	mux := mux.NewRouter() //gorilla/mux 패키지 자동 임포트

	// 각각 경로를 추적하는 mux들..
	mux.HandleFunc("/", indexHandler) //하위경로 미정의시는 상위경로가 자동호출된다
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)
	// mux.HandleFunc("/self", selfHandler)
	return mux
}
