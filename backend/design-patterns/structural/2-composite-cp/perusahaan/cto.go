package perusahaan

<<<<<<< HEAD
=======
import "fmt"

>>>>>>> fc12154791502702980a046e3507ab317e48f675
type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
<<<<<<< HEAD
	total := c.GetSalary()
	for _, employee := range c.Subordinate {
		total += employee.TotalDivisonSalary()
	}
	return total
}
=======
	return 0 // TODO: replace this
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
