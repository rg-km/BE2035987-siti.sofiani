package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here
	getAreaFunc := func (y int) func (int)int  {
		return func(x int) int {
			return x*y
		}
	}
	areaF := getAreaFunc(4)
	res := areaF(3) // 12
	fmt.Println(res)
}
