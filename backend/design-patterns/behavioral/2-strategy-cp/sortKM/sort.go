package sortKM

//kali ini kita coba menggunakan strategy pattern untuk menentukan sort
//ada dua sort yaitu ascending dan descending sort

type Strategy interface {
	Sort([]int)
}

type SortKM struct {
	Strategy Strategy
}

func (s *SortKM) Sort(array []int) {
<<<<<<< HEAD
	s.Strategy.Sort(array)
=======
	// TODO: answer here
>>>>>>> fc12154791502702980a046e3507ab317e48f675
}

func (s *SortKM) SetStrategy(Strategy Strategy) {
	s.Strategy = Strategy
<<<<<<< HEAD
}
=======
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
