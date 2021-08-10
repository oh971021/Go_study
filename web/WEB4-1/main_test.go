package main // upload 넷이 잘 작동하는지 테스팅! - 이것은 테스팅 패키지 이므로 _test.go로 파일 작성.

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

//테스팅 시 뭐와 맵핑 시켜야하는가?  맨처음에 convey 부터 세팅해준다.
//가상 목업용 웹 세팅!
func TestUploadTest(t *testing.T) { // 테스팅 패키지는 양식 맞추어 주면 자동으로 패키지 생성.
	assert := assert.New(t)
	//----------파일 읽어주기---------------------------------------------------------------------------------------------------//
	path := "C:/Users/11/desktop/1.png" // 자기 업로드 파일 저장경로에서
	file, _ := os.Open(path)            // _ 무시 (err) 파일을 받아서 os에 보관. / 위치가 지정안되있으면 Formfile이 req로 받은 파일 보관 중. 폼데이터
	defer file.Close()

	os.RemoveAll("./uploads") //이전업로드저장된 파일을 모두지운다. 밑에 단계 해석 후 마지막에 볼 것.

	//Form파일 만들기---------------------------------------------------------------------------------------------------//
	buf := &bytes.Buffer{}                                                  //데이터는 버퍼에 싣려있다. 버퍼인스턴스 바이트값주소
	writer := multipart.NewWriter(buf)                                      //MIME:웹기반 메일에서 파일전송 표준포멧 multipart는 파일전송 기본포맷형식
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path)) //Base(path)는 업로드된 파일경로에서 마지막 경로만 반환. 즉 파일이름만 보냄.
	//CreateformFile==> 업로드 내용에서 파일만 잘라서 가져와라. 패스경로는 버리고
	assert.NoError(err) //검사. 폼파일 만드는데 에러가 없어야한다.

	//----데이터 넣어주기---------------------------------------------------------------------------------------------------//
	io.Copy(multi, file) //목적지(destination)은 multi, 어디서 가져오냐?file(위에서 자른 친구) 에서 가져와서
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf) //Method는 POST, target은 /uploads, Body에는 데이터가 싣려있는 buf로 보낸다
	// 목업 방식으로 req로 새로운 요청을 때려라. 위의 뉴 레코드는 가상응신인데 저것은 바디의 렌더링된 데이터
	req.Header.Set("Content-type", writer.FormDataContentType()) //가져온 file, Form 데이터형식을 알려주어야 하고, 에러시 아래 content-type지정해결

	uploadsHandler(res, req)              //콜, 데이터 전송완료
	assert.Equal(http.StatusOK, res.Code) //검사단계 : Equal은 res.code는 StatusOK와 동일해야

	//-----------------------------------------------------200번 StatusOK Pass 여부 검사 확인 -----------------------------------------------------------------//

	uploadFilePath := "./uploads/" + filepath.Base(path) //업로드 파일이름을 끊어와서 업로드와 파일이름 합성
	_, err = os.Stat(uploadFilePath)                     //Stat함수는 파일정보 알려준다 _, 파일인포는 필요없고 err만
	assert.NoError(err)                                  //Error가 없는지 검사해라~

	uploadFile, _ := os.Open(uploadFilePath) //업로드 파일패스를 불러와서 업로드 파일에 담아주고.
	originFile, _ := os.Open(path)           // 위 path원래파일과 확인/ 안짤린 긴 경로.
	defer uploadFile.Close()                 //오픈 파일 자원을 닫아준다
	defer originFile.Close()                 //오픈 파일 자원을 닫아준다
	//자원이 없어진게 아니라 업로드와 오리진 2개가 존재하고있음.
	//이거를 어떤 방식으로 읽어오느냐?

	uploadData := []byte{}      //업로드 데이터 읽어온다. 바이트 배열 방식으로
	originData := []byte{}      //오리진 데이터 읽어온다. 바이트 배열 방식으로
	uploadFile.Read(uploadData) // 합성된 데이터
	originFile.Read(originData) // 두데이터 준비 완료, 본래의 데이터

	assert.Equal(originData, uploadData) // 어썰트로 원본과 업로드 데이터가 같은지 분석해라!*/
}
