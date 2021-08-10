package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) { //상대방보내준파일 리드 부분
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil { //파일이 미전송 에러가 있다면
		w.WriteHeader(http.StatusBadRequest) //잘못된 파일이 전송되었을때
		fmt.Fprint(w, err)                   //에러 반환
		return
	}
	defer uploadFile.Close() //핸들러 사용후 defer로 닫음

	dirname := "./uploads"                                     //업로드된 파일 저장할 공간 만드는 부분
	os.MkdirAll(dirname, 0777)                                 //읽고 쓰고 지울수 있는 디렉토리 만들어라
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename) //%s폴더명/%s는 파일명
	file, err := os.Create(filepath)                           //비어있는 새로운 파일을
	defer file.Close()                                         //파일을 만들고 닫아줘야함

	if err != nil { //미전송된 파일 에러가 발생한다면
		w.WriteHeader(http.StatusInternalServerError) //내부상태 에러값
		fmt.Fprint(w, err)
		return
	}

	io.Copy(file, uploadFile)    //파일패스에서 들고와서 비어있는 파일에 업로드파일을 복사해줌
	w.WriteHeader(http.StatusOK) //err없이 잘 되면 ok신호
	fmt.Fprint(w, filepath)      //filepath는 upload경로를 알려준다(어디에 업로드되는지)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public"))) //파일웹서버만드는거

	http.ListenAndServe(":3000", nil)
}
