package myapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("bar handler 성공")

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello bar!", string(data))
}

func TestFooHandler_Withjson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req)
	fmt.Println("foo handler 성공")

	assert.Equal(http.StatusBadRequest, res.Code)
}
