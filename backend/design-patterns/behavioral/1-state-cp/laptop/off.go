package laptop

// Ini adalah concrete implementation dari state Off
type Off struct {
	Laptop *Laptop
}

func (o Off) Press() {
	if !o.Laptop.IsThereBattery() {
		return
	}
	// chane laptop currentstate to "On" by memory
	o.Laptop.ChangeState(&On{o.Laptop})
}

func (o Off) CanTurnOnLaptop() bool {
	return false
}

func (o Off) Sleep() {
	return
<<<<<<< HEAD
}
=======
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
