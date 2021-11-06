# My study note by Golang

##### 사용툴 : VSCode, Golang, HTML, CSS, JS

### 1. 환경 셋팅

- <a href="http://golang.org/dl">Golang install </a>

![image](https://github.com/pandora0667/TILD/blob/master/screenshot/go/windows-install/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202020-07-04%20%EC%98%A4%EC%A0%84%2012.31.07.png?raw=true)
![Go 설치 1](https://user-images.githubusercontent.com/84692769/138222131-6343de57-4e17-47bf-b2f2-feb1cc42551c.jpg)
![Go 설치 2](https://user-images.githubusercontent.com/84692769/138222140-18095dfe-d3a3-4fc8-a291-8df4caedcd0a.jpg)
![Go 설치 3](https://user-images.githubusercontent.com/84692769/138222148-ce492f07-884e-430c-87de-0c4832a76a20.jpg)


### 2. 환경 변수 설정하기
##### &nbsp;&nbsp; - GOROOT : Go가 설치된 디렉토리, GO 표준 패키지 사용 용도
##### &nbsp;&nbsp; - GOPATH : 사용자가 작업하게 되는 공간 루트, 깃허브에서 패키지를 사용하게 될 경우 저장되는 경로
##### &nbsp;&nbsp; - 설정경로 : 제어판 > 시스템 및 보안 > 시스템 >　고급 시스템 설정 > 환경변수
![image](https://user-images.githubusercontent.com/84692769/138221516-7e877a8c-93b6-4d91-9915-b5c665e3aa84.png)   ![image](https://user-images.githubusercontent.com/84692769/138221539-7ab1bade-3cea-40e2-8858-6574304b1097.png) <h5> 시스템 변수에 GOROOT 설정된 것을 확인할 수 있다. </h5>

<h5>시스템 변수에 Path에서 편집을 눌러서 경로 확인 가능 : (설치된 경로) C:\Go\bin 확인 </h5>

![image](https://user-images.githubusercontent.com/84692769/138221616-e9a53794-f19a-4d45-8a0f-6d81f1bb2a21.png)   ![image](https://user-images.githubusercontent.com/84692769/138221625-fdd3e3f3-4943-4d65-a2c1-b4729f7b381a.png) 

<h5>마무리로 cmd에서 go version / go env 로 설치, 환경변수가 잘 되었는지 확인한다.</h5>

![Go 설치 4](https://user-images.githubusercontent.com/84692769/138222153-b35f2b24-1a8e-4a52-9326-a6c5b8bfc5aa.jpg)
