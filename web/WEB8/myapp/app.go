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
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
}

func getUserInfoHandeler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
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

// Delete User
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

// -------------------- updateUserHandler ------------------------ //

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// updateUserHandler를 작동시키면 JSON 데이터가 오도록 User Struct를 만들어 준다.
	// 그 Struct를 새로운 인스턴스로 생성해준다.
	updateUser := new(User)

	// JSON 형식의 데이터가 들어가있는 인스턴스의 Body를 읽어온다.
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// updateUser의 ID정보가 user에 잘 들어갔는지 확인한다.
	user, ok := userMap[updateUser.ID]
	if !ok {
		// 잘 들어간 것이 확인 된다면 Header에 StatusOK을 찍어주고
		w.WriteHeader(http.StatusOK)
		// Update User ID: {id[]}를 출력해준다.
		fmt.Fprint(w, "No User id:", updateUser.ID)
		return
	}

	// ------------------ Data 갱신 -------------------- //

	// 각각의 정보에 해당하는 매개값(id)를 통해 출력값(body)을 갱신해준다.
	// Data가 잘 들어가면 StatusOK~~ data를 출력해준다.

	// ---- Client 요청시 비워있지 않은 경우만 변환한다 ---- //
	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	// Header에는 Type과 Json 형식을 사용하겠다고 입력해준다.
	userMap[updateUser.ID] = user // userMap에서 ID가 존재하면 ID를 변경해준다. user는 포인터타입이므로 바꿀필요가없다.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// ------------------------------------------------- //

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler).Methods("GET")        // w.r
	mux.HandleFunc("/users", createUserHandler).Methods("POST") // encode.decode

	// PUT update 핸들러를 등록해준다.
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")

	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandeler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	return mux
}
