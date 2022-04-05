package main

import "fmt"

func main() {
	// print selamat pagi menggunakan anonymous function
	// TODO: answer here
	func (){
		fmt.Println("selamat pagi")
	}()
	fmt.Println()

	func (nama string)  {
		fmt.Println(nama)
	}("ikhsan")
}
