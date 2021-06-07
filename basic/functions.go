package main

// func add(x, y int) int 가능
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked return" 주어진 반환 값을을 반환
}
