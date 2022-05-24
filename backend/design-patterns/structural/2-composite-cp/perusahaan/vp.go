package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
<<<<<<< HEAD
	total := vp.GetSalary()
	for _, employee := range vp.Subordinate {
		total += employee.TotalDivisonSalary()
	}
	return total
}
=======
	0 // TODO: replace this
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
