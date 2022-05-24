package laptop

// Ini adalah concrete implementation dari state On
type On struct {
	Laptop *Laptop
}

func (o On) Press() {
	o.Laptop.CurrentState = "Off"
	o.Laptop.ChangeState(Off{o.Laptop})
}

func (o On) CanTurnOnLaptop() bool {
	return true
}

func (o On) Sleep() {
<<<<<<< HEAD
	o.Laptop.ChangeCurrentState("Sleeping")
	o.Laptop.ChangeState(Sleeping{o.Laptop})
}
=======
	// TODO: answer here
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
