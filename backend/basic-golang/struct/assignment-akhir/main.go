//dikerjakan jika waktu cukup saja
package main

import "fmt"

type Employee struct {
	Name              string
	BaseSalary, Bonus int
}

//method GetFullSalary() untuk menampilkan salary + bonus
//method AddBonus(bonus int) untuk mengubah nilai atribut bonus employee

// TODO: answer here
func (e Employee) GetFullSalary(){
	fullsalary := e.BaseSalary+e.Bonus
	fmt.Println("nama :", e.Name, "full salary :", fullsalary)
}
func (e *Employee) AddBonus(b int){
	A := e.Bonus+b
	fmt.Println("nama :", e.Name, "Bonus :", A)
}


func main() {
	employee1 := Employee{Name: "bob", BaseSalary: 4000000, Bonus: 300000}
	fmt.Println("sebelum bonus dinaikkan")
	employee1.GetFullSalary()
	fmt.Println("tambah 100000 ke bonus")
	employee1.AddBonus(100000)
	fmt.Println("setelah bonus dinaikkan")
	employee1.GetFullSalary()
}
