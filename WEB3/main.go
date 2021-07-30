package main

import (
	"net/http"
	"web/myapp"
)

func main() {
	// mux 가 들어오면 응신값을 라우팅되는 서버값이 움직인다.
	// nil 이 들어오면 퍼스널 웹서버값이 움직인다.

	// 결론 : 밑에 요 서버 지정 코드는 myapp의 newhttphandler 만을 움직이는 서버를 만드는 것.
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
	// myapp 에서 handler 를 만든다.
}
