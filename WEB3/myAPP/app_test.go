package myapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
testxxxx 와 같은 이름으로 짓는 이유 : 해당 메소드가 테스트메소드임을 알리는 것
testing package의 T포인터를 인자로 받음 ( 정해진 양식이라고 함 )
testing.T 를 하나의 입력으로 받고, 출력은 없다
결국 app_test.go 는 app.go의 함수들을 테스트하기 위한 모듈
*/
func TestIndexPathHandler(t *testing.T) {

	// asser 패키지를 쓰는이유 : 해당 패키지의 테스트 함수가 테스트하는데 유용하다.
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) //실제 response를 사용하지 않고 테스팅 응신방법

	indexHandler(res, req)
	fmt.Println("wow")

	// assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }

	// StatusOK와 res.Code가 같은지 비교한다.
	assert.Equal(http.StatusOK, res.Code)
	// iotil 패키지에 들어있는 readall
	// Readall(response의 body에 뿌려진 정보를 읽어와라)
	data, _ := ioutil.ReadAll(res.Body)
	// 적어놓은 헬로월드와 리스폰스의 바디에서 읽어온 것을 비교한다.
	assert.Equal("Hello World!", string(data))
}

func TestIndexPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar", nil)

	indexHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}
