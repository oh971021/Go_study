// 모듈 = 패키지의 상위개념이다. (패키지를 포함하고 있는 것)
// 프로그램 = 코드를 통해 제작된 실행파일 / 실행시작 지점을 포함한 패키지.
// 즉, main패키지를 빌드하면 프로그램이된다.

package main // 코드를 묶어놓는 단위. 모든 코드는 반드시 패키지로 묶여야 한다.
// main 패키지가 아닌 다른 그외 패키지는 (실행시작지점이 없는 보조패키지)

// 메인 패키지, 보조 패키지 모두를 가지고 있는게 모듈이다.

// 프로세스 = 프로그램을 실행하면 OS를 로드(메모리에 올려서 프로그램을 실행하는 것)한다.
// ex) 계산기를 켜면 프로그램이 실행되는데 이것이 로드 된 프로세스라고 볼 수 있다. 같은 프로그램에 여러개의 프로세스가 존재한다.

import ( // 외부, 보조패키지를 가져올 때 사용한다.
	"fmt"
	"math/rand" // 경로의 마지막에 오는 패키지가 사용되는 패키지이다.
	_ "strings" // 실제 사용하지 않아도 패키지 초기화에따른 부가효과를 위해 사용할때가 있다.
)

func main() {
	var a int
	fmt.Println(rand.Int(), a)
}