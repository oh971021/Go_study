package myapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		strings.NewReader(`{"first_name":"junseok", "last_name":"oh", "email":"oh971021@gmail.com"}`))
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

//---5단계 TestDeleteUser 테스팅
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Delete는 기본메소드가 아니다. 다만, Do(req) request에 메소드정의, Id는 임의로/Users/1, Body 값은 없으므로 nil로 한다.
	// NewRequest 는 ( <메소드> , URL , body ) 요렇게 3가지 값을 받아서 Method 를 만들어 준다.
	// 여기서 메소드화 시켜주고, 이 친구를 이제 아래의 Default 에서 적용 시키는 것.
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil) // 바디문에는 nil로 맞추겠다!

	// Do(req): Do(Method)="DELETE" 메소드를 정의
	// 요 친구가 DELETE 메소드를 생성해주는 듯 하다.
	// Default 클라이언트 라는 레퍼런스를 가져온다.
	// 그 안에 Do 라는 친구를 사용하면 request를 받을 수 있다.
	// 그 값을 받아서 response 랑 error 를 반환하는 친구.
	resp, err := http.DefaultClient.Do(req)

	// err 는 없어야 한다.
	assert.NoError(err)

	// 그리고 body의 값이 잘 들어왔는지 확인을 해본다.
	assert.Equal(http.StatusOK, resp.StatusCode)

	// 아래 log로 찍어보기위해 ReadAll(Body)로 읽어온다, delete할것이 없었다
	data, _ := ioutil.ReadAll(resp.Body)

	// log로 찍어본다.
	// 이건 언제 실행하는지 확인하는 듯 함.
	log.Print(string(data))

	// "지울게 없었다 즉 유저가 없었다"는 메세지를 포함해야 한다, 위 임의등록시킨 1번ID를 등록시킨 적이 없다
	// 지금 body 에는 nil 값. 즉, 아무 값도 없기 때문에 No User Id 가 출력 되는 것이다.
	assert.Contains(string(data), "No User Id:1")

	//[등록] user map id 1을 [등록]하는 부분, Post전송방식
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"junseok", "last_name":"oh", "email":"oh971021@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	//---위 POST방식 Json 정보를 서버가 user정보를 받아서 user정보를 리턴하는 부분
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user) //서버가 보낸 데이터를 읽어온다 = 매개변수 입력
	assert.NoError(err)                           //아무런 문제err가 없다.
	assert.NotEqual(0, user.ID)                   //user ID

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	log.Print(string(data))
	assert.Contains(string(data), "Deleted User Id:1")
}

// func TestDeleteUsers(t *testing.T) {
// 	assert := assert.New(t)

// 	ts := httptest.NewServer(NewHandler())
// 	defer ts.Close()

// 	req, _ := http.NewRequest("DELETE", ts.URL+"users/1", nil)
// 	resp, err := http.DefaultClient.Do(req)
// 	assert.NoError(err)
// 	assert.Equal(http.StatusOK, resp.StatusCode)
// 	data, _ := ioutil.ReadAll(resp.Body)
// 	assert.Contains(string(data), "No User ID:1")

// 	//----Post로 보냈을 때 users에 POST방식전송, 정보는 Json 방식으로 전송한다
// 	resp, err = http.Post(ts.URL+"/users", "application/json",
// 		strings.NewReader(`{"first_name":"sorli", "last_name":"noh", "email":"vivace8530@gmail.com"}`))
// 	assert.NoError(err)
// 	assert.Equal(http.StatusCreated, resp.StatusCode)

// 	user := new(User)
// 	err = json.NewDecoder(resp.Body).Decode(user)
// 	assert.NoError(err)
// 	assert.NotEqual(0, user.ID) // user ID가 0이 아니다 등록기록됨

// 	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
// 	resp, err = http.DefaultClient.Do(req) // Itoa:정수형을 문자열 변환, ID정보를  Get방식으로/user/ID를 넣어서 오도록 만듬

// 	assert.NoError(err)
// 	assert.Equal(http.StatusOK, resp.StatusCode)
// 	data, _ = ioutil.ReadAll(resp.Body)
// 	assert.Contains(string(data), "Deleted User ID:1")
// }
