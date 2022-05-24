package main

import "github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	runes := []rune(s)
	length := len(runes)
	if length < 2 || length%2 != 0 {
		return false
	}
	stack := stack.NewStack()
	for _, r := range runes {
		prev, err := stack.Peek()
		if err != nil || !isPair(prev, r) {
			stack.Push(r)
			continue
		}
		stack.Pop()
	}
	return stack.IsEmpty()
}

func isPair(a, b rune) bool {
	if (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}') {
		return true
	}
	return false
}