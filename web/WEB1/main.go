package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	// 어노테이션을 붙혀주기 전에는 text file로 저장이 된다.
	// 변경 하지 않고 사용을 하면 결과값에 아래 지정 값이 나오지 않는다.
	// json 형식이 아니다. (golang != json)

	// `` = 어노테이션(Annotation)을 붙혀주는 것 = json 형식으로 표준을 잡겠다
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)

	// 헤더 작성하기
	w.Header().Add("content-type", "applicatioon/json")
	w.Header().Add("testing", "I'm testing respnsewritier")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
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
