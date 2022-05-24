package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
<<<<<<< HEAD
	s.Laptop.ChangeCurrentState("On")
	s.Laptop.ChangeState(On{s.Laptop})
=======
	// TODO: answer here
>>>>>>> fc12154791502702980a046e3507ab317e48f675
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
<<<<<<< HEAD
	return
}
=======
	// TODO: answer here
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
