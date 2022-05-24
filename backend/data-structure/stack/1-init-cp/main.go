package main

// Dari contoh yang telah diberikan, cobalah untuk membuat stack yang memiliki jumlah maksimal elemen sebanyak 10 dengan menambahkan atribut Size.
// Gunakan slice untuk menyimpan data pada stack. Pada kasus ini, data yang disimpan berupa int.
// Lengkapi function NewStack sehingga function tersebut dapat mengembalikan objek Stack dengan memiliki ukuran dari parameter.
type Stack struct {
	Size int
	Top int
	Data []int
}

func NewStack(size int) Stack {
	return Stack{
		Size: size,
		Top: -1,
		Data: make([]int, 0, size),
	}
}