package main

import (
	"fmt"

	// "golang.org/x/text/width"
)

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

// TODO: answer here
type Rectangle struct {
	width int
	Length int
}
func (r *Rectangle) SetWidthPointer(width int){
	fmt.Println("hasil :", r.width + width) 
}
func (r Rectangle) SetWidthValue(width int) {
	fmt.Println("hasil :", r.width + width) 
}
func main() {
	var r Rectangle
	r.width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.width)
	r.SetWidthValue(70)
}
