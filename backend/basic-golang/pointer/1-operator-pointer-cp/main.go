package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	// TODO: answer here
	n := &name
	a := &age
	i := &isMarried
	fmt.Println("Nilai dari alamat memory", n ,"adalah",*n)
	fmt.Println("Nilai dari alamat memory", a ,"adalah",*a)
	fmt.Println("Nilai dari alamat memory", i ,"adalah",*i)
}
