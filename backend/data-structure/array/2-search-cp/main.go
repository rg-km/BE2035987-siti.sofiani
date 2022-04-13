package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	//return []string{}, nil 
	SearchMatch := []string{}
	for a1 := range arr1 {
		for a2 := range arr2 {
			if arr1[a1] == arr2[a2] {
				SearchMatch = append(SearchMatch, arr1[a1])
			}
		}
	}
	if len(SearchMatch) == 0 {
		return nil, fmt.Errorf("no match")
	}
	return SearchMatch, nil
}
