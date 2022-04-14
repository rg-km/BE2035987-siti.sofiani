// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {
	// return "" // TODO: replace this
	var str1Map = make(map[rune]int)
	var str2Map = make(map[rune]int)

	for _, char := range str1 {
		str1Map[char]++
	}

	for _, char := range str2 {
		str2Map[char]++
	}

	if len(str1Map) != len(str2Map) {
		return "Bukan Anagram"
	}

	for key, value := range str1Map {
		if str2Map[key] != value {
			return "Bukan Anagram"
		}
	}

	return "Anagram"
}