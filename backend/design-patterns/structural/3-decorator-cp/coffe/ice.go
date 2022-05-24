package coffe

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fc12154791502702980a046e3507ab317e48f675
	"strings"
)

// Nah tetapi ada condiment yang unik disini yaitu Ice.
// Jika sebuah coffe ditambahkan Ice maka akan terdapat Magic didalamnya.
// Magic ini terjadi jika Coffe tersebut berharga diatas 3.30 dan Ice sudah ditambahkan sebelumnya (2x ice)
// Magic tersebut adalah jika Coffe tersebut sudah ditambahkan Ice sebelumnya maka Kopi tersebut menjadi BEKU

type Ice struct {
	Coffe Coffe
}

func (i Ice) GetCost() float64 {
	return i.Coffe.GetCost() + 0.20
}

func (i Ice) GetDescription() string {
<<<<<<< HEAD
	desc := fmt.Sprintf("%v, Ice", i.Coffe.GetDescription())
	if checkRepeatingWord(desc) {
		desc = fmt.Sprintf("%v, BEKU", desc)
	}
	return desc
=======
	return description // TODO: replace this
>>>>>>> fc12154791502702980a046e3507ab317e48f675
}

// check repeating word in a string
// Use this to check repeating word in string ;)
func checkRepeatingWord(s string) bool {
	input := strings.Fields(s)
	for _, word := range input {
		if strings.Count(s, word) > 1 {
			return true
		}
	}
	return false
<<<<<<< HEAD
}
=======
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
