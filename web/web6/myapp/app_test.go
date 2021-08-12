package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo")
}

func TestGetUsersInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:89")
}

func TestCreateUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	//----Post로 보냈을 때 users에 POST방식전송, 정보는 Json 방식으로 전송한다
	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sorli", "last_name":"noh", "email":"vivace8530@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID) // user ID가 0이 아니다 등록기록됨

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id)) // Itoa:정수형을 문자열 변환, ID정보를  Get방식으로/user/ID를 넣어서 오도록 만듬
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)                             // := 인수입력, 함수호출시 함수로 값을 전달해주는 값
	err = json.NewDecoder(resp.Body).Decode(user2) // 정보파싱, =매개변수 입력, 함수정의에서 전달받은 인수를 함수 내부로 전달하는 변수
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID) //user1 ID는 cteate한 user.ID이고 새로운정보를 받는 user2 ID user2.ID
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	//----Post로 보냈을 때 users에 POST방식전송, 정보는 Json 방식으로 전송한다
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sorli", "last_name":"noh", "email":"vivace8530@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID) // user ID가 0이 아니다 등록기록됨

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req) // Itoa:정수형을 문자열 변환, ID정보를  Get방식으로/user/ID를 넣어서 오도록 만듬

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID:1")
}
