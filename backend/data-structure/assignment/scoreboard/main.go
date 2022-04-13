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
	total := make(map[string]int)

	for _, v := range s{
		total[v.Name] = (v.Correct * 4) - v.Wrong
	}

	if total[s[i].Name] == total[s[j].Name] {
		if s[i].Correct < s[j].Correct {
			return true
		} else if s[i].Correct == s[j].Correct {
			return false
		}
	}

	return total[s[i].Name] < total[s[j].Name]
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(sort.Reverse(s))
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
