module tutorial/usepkg

go 1.16

require (
	github.com/guptarohit/asciigraph v0.5.2
	github.com/tuckersGo/musthaveGo/ch16/expkg v0.0.0-20210521175101-4a5fd1ed5e1b
)

// usepkg 모듈 안에 custompkg, program 두개 의 패키지가 담겨져있다.
// 두개의 패키지 중 메인패키지가 실행파일이고, 나머지는 실행 패키지의 보조 패키지