package main

import (
	"fmt"
)

var person = "andi"
var names = []string{"budi", "toni", "adi", "ado", "alif", "yudi"}

//mengembalikan string, dimana `name` menyapa semua `names`
func greetAll(person string, names []string, output chan<- string) {
	// TODO: answer here
	for _, name := range names {
		output <- "hello " + name
	}
	fmt.Println("selesai mengirim")

}

// buat size buffered channel sesuai jumlah names
func testBufferedChannel(result chan<- string) {
	size := len(names)
	c := make(chan string, size)
	go greetAll(person, names, c)
	for i := 0; i < size; i++ {
		result <- <-c
	}
	fmt.Println("selesai mengirim")
}

//goroutine greetAll dapat mengirim ke goroutine testBufferedChannel walaupun channel belum siap menerima
