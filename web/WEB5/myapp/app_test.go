package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//---1단계 TestIndex  NewHandler()등록빌드 검증테스팅---//

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err) //에러가 없어야 한다
	assert.Equal(http.StatusOK, resp.StatusCode)
	//--2단계 Hello World 데이터 테스팅 검사
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

//--------2단계 TestUsers 등록 검사------------//
func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo") //Get UserInfo포함하고 있으면 pass 아니면 fail
}

// ---------3단계 TestGetUserInfo Get id 검사----------//
func TestGetUsersInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// --- id 1 --- //
	resp, err := http.Get(ts.URL + "/users/1")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User id:1")
}

// func TestSelf(t *esting.T {
// 	assert := assert.Newt)

// 	ts := httptest.NewServer(NewHander))
// 	defer ts.Clos()

// 	resp, err := htp.Get(ts.URL + "/sel")
// 	assert.NoError(er)
// 	assert.Equal(htt.StatusOK, resp.StatusCoe)
// 	data, _ := ioutil.ReadAll(resp.Boy)
// 	assert.Contains(string(data), "Self") //Get UserInfo포함하고 있으면 pass 아니면 fil
// }
