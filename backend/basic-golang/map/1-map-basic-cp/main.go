package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	nama := map [string] string {
		"jawabarat": "bandung",
		"dkijakarta": "jakarta",
		"jawatengah":"semarang"}

	for key, value := range nama {
		fmt.Println("nama provinsi: ", key, " ibukota: ", value)
	}
}
