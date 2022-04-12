package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	// check pertama
	// jika total score di index i lebih besar dari total score di index j
	// kembalikan true
	// jika total score di index j lebih besar dari total score di index i
	// kembalikan false

	scoreI := (s[i]).Correct * 4) - (s[i]).Wrong *1)
	scoreI := (s[j]).Correct * 4) - (s[j]).Wrong *1)

	if scoreI > scoreJ {
		return true
	} else{
		return false
	}

}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}
