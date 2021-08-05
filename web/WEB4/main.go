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
		fmt.Fprint(w, err)                  //에러 반환
		return
	}
	defer uploadFile.Close() //핸들러 사용후 defer로 닫음

	dirname := "./uploads" //업로드된 파일 저장할 공간 만드는 부분
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return

	}

	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public"))) //파일웹서버만드는거

	http.ListenAndServe(":3000", nil)
}
