package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi) // 대문자로 시작하는 이름이 export

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
