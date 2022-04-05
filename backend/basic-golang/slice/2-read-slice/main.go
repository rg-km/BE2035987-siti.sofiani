package main

import "fmt"

// Untuk membaca slice itu sebenarnya sama seperti pada array.
// Kita cukup mengambil index nya
func main() {
	var slice []string
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	fmt.Println(slice[0])
	fmt.Println(slice[9])

	// Tapi ada tips nih untuk membaca slice ataupun array. kita bisa menggunakan
	// for range yang akan menampilkan index dan nilai yang ada pada array tersebut.
	// for loop akan dijelaskan lebih lengkap di materi looping

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
