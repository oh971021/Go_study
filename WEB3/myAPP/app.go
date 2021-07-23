package myapp

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
	FirstName string    `json:"first_name"` //어노테이션)(Annotaion) 설명을 붙히는것
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprint(w, "Hello Index!") // Fprint은 w값에 "Hello World" 값을 주어서 프린팅해라는 의미
}

type fooHandler struct{} //인스턴스 생성

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //ServeHTTP 인터페이스 구현
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Hello Foo!")
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) //인터페이스, JSON형식 결과값은 바이트 어레이
	//w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "bar"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() //라우터 클래스 만든다. mux 등록
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler) //mux

	mux.Handle("/foo", &fooHandler{}) // mux
	return mux
	//요 함수를 통해 웹서버 간단제작.
}
