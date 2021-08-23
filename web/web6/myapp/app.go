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
	// 여기서 id값을 string으로 출력한다.
	vars := mux.Vars(r)
	// 근데 id는 Query에서 int type으로 사용하기 때문에 int로 바꿔준다.
	// vars는 strinf이므로Atoi로 int정수형으로 바꾸면 첫번째 인티저형id와,두번째err가 나온다
	id, err := strconv.Atoi(vars["id"])
	// 변환이 잘 되지 않았따면
	if err != nil {
		// 배드리퀘스트 = 400 을 보내준다.
		w.WriteHeader(http.StatusBadRequest)
		// 그리고 화면에 400을 찍어주고, error 내용을 출력한다.
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
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User Id:", id)
	return
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler).Methods("GET")        // w.r
	mux.HandleFunc("/users", createUserHandler).Methods("POST") // encode.decode
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandeler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	return mux
}
