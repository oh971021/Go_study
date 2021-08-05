package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing" // 요녀석은 xx_test.go 로 만들면 생겨나는 테스팅 모듈

	"github.com/stretchr/testify/assert"
)

// app_test.go 로 이름을 지정해줘서 TEST하는 모듈이라는 것을 지정한다.

/*
testxxxx 와 같은 이름으로 짓는 이유 : 해당 메소드가 테스트메소드임을 알리는 것
testing package의 T포인터를 인자로 받음 ( 정해진 양식이라고 함 )
testing.T 를 하나의 입력으로 받고, 출력은 없다
결국 app_test.go 는 app.go의 함수들을 테스트하기 위한 모듈
*/
func TestIndexPathHandler(t *testing.T) {

	// asser 패키지를 쓰는이유 : 해당 패키지의 테스트 함수가 테스트하는데 유용하다.
	assert := assert.New(t)

	res := httptest.NewRecorder() // 실제 response를 사용하지 않고 테스팅 응신방법
	// nil := pointer 초기값을 셋팅 해주기 위해서 입력해준다.
	// 이자리에 mux 도 들어 갈 수 있는데, 라우팅을 할 때 사용.
	req := httptest.NewRequest("GET", "/", nil)

	// mux : "/" 경로를 타겟 해서 라우팅(분배), 랜더링 부분
	// mux : 라우터 (분배의 기능, 경로 타겟팅) - 경로에 맞춰서 분배를 해주어라.
	mux := NewHttpHandler()
	// 이 단계에서의 핵심 :
	mux.ServeHTTP(res, req)

	// indexHandler(res, req)
	fmt.Println("index handler 성공")

	// assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }

	// StatusOK와 res.Code가 같은지 비교한다.
	assert.Equal(http.StatusOK, res.Code)

	// iotil 패키지에 들어있는 readall
	// iotil 은 err, return 두개의 값을 받는다. 근데 에러는 받지 않기 위해 _로 표시했다.
	// Readall(response의 body에 뿌려진 정보를 읽어와라)
	// app.go 의 index handler 에서 render data를 읽어오는 부분이다.
	data, _ := ioutil.ReadAll(res.Body)

	// 적어놓은 헬로월드와 리스폰스의 바디에서 읽어온 것을 비교한다.
	assert.Equal("Hello Index!", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("bar handler 성공")

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello bar!", string(data))
}

func TestFooHandler_Withoutjson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("foo handler 성공")

	// assert.Equal(http.StatusOK, res.Code)
	// data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	// requset 과정에서 경로에 name을 입력해준다.
	req := httptest.NewRequest("GET", "/bar?name=OhJunSeok", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("bar handler 성공")

	assert.Equal(http.StatusOK, res.Code)
	// 리스폰스된 페이지의 body 부분을 읽어온다.
	data, _ := ioutil.ReadAll(res.Body)

	// name까지 입력한 경로에 저장되고, body 에 값이 변경되어 저장된다.
	// 내가 입력한 값과 body에 입력된 값을 비교한다.
	// app.go 에 가보면 경로에 따라 hello ~~ 가 바뀌는걸 확인한다.
	assert.Equal("Hello OhJunSeok!", string(data))
}

func TestFooHandler_Withjson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법

	req := httptest.NewRequest("POST", "/foo",
		// 새로운 모듈 strings 를 사용. 그 안에 내장함수 NewReader 사용.
		strings.NewReader(`{"first_name":"junseok", "last_name":"oh", "team":"KOSK", "email":"oh971021@gmail.com"}`))

	// 라우터 mux : 등록 된 경로의 목록에 대해 들어온 요청과 일치 URL, 경로에 대한 핸들러 호출
	mux := NewHttpHandler()

	// 현재 사용하는 HTTP와 호환되는지 확인.
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("foo handler 성공")

	// IANA에 등록된 HTTP 상태 코드(statusOK는 200)와
	// response된 코드(현재 app.go에서 잘 전달되었기 때문에 200코드로 나타남)를 비교, 확인한다.
	assert.Equal(http.StatusOK, res.Code)

	// new : built-in 함수 - 새로운 메모리 공간을 할당함
	// User 타입의 user 변수이고, 새로운 메모리 공간에 생성
	user := new(User)

	// json 패키지에서 NewDecoder 내장함수를 불러옴
	// 요 함수는 json 값으로 매개변수에 들어온 값을 읽을 수 있도록 한다.
	// 결론 : user 라는 메모리 공간 안에 res.Body 에 저장된 json 값을 불러온다는 뜻인 것 같다.
	err := json.NewDecoder(res.Body).Decode(user)

	// err 변수가 nil인지 아닌지를 나타내는 함수
	assert.Nil(err)

	// 직접 입력한 값과 Json에 입력되어있는 값을 비교함
	assert.Equal("junseok", user.FirstName)
	assert.Equal("oh", user.LastName)

	// Team KOSK 에서 활용해서 넣은 부분 (app.go와 json값 모두 변경)
	assert.Equal("KOSK", user.Team)
}
