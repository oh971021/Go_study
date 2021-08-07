package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.new(t)

	// test server 용 서버
	ts := httptest.NewServer(NewHandler())
	// 서버는 항상 defer 로 닫아줘야한다.
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.Code)

	date, _ := iotil.ReadAll(res.Body)
	assert.Equal("Hello World", string(body))
}
