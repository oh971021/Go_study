package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandeler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // vars는 strinf이므로Atoi로 int정수형으로 바꾸면 첫번째 인티저형id와,두번째err가 나온다
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user) // json.Marshal는 Go밸류를 JSON 문자열로 변환-->슬라이스나 바이트 기반으로 go value를 반환
	fmt.Fprint(w, string(data))
}

// User struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// 요 Structure 객체는 목적지가 어딘가?? = userMap으로 간다이
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) //decode는 JSON문자열 go value반환 ,encode는 반대 go value바를 JSON문자열
	if err != nil {                             // 에러가 있따면 잘못된 JSON형식
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	// new(User) 의 값을 받아서 저장하는 Map을 만든다.
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler).Methods("GET")        // w.r
	mux.HandleFunc("/users", createUserHandler).Methods("POST") // encode.decode
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandeler)
	return mux
}
