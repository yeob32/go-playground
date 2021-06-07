package main

import "fmt"

// 함수 밖에서는 모든 선언이 키워드(var, func, 기타 등등)로 시작
// := 구문 사용 불가
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
