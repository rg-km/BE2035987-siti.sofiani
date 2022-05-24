package main

import "fmt"

var Solution []string

func TowerOfHanoi(n int, fromRod string, auxRod string, toRod string) {
	if n == 0 {
		return
	}

	// TODO: answer here
	TowerOfHanoi(n-1, fromRod, auxRod, toRod)
	Solution = append(Solution, fmt.Sprintf("Move disk %d from rod %s to rod %s", n, fromRod, toRod))
	// TODO: answer here
	TowerOfHanoi(n-1, auxRod,toRod, fromRod)
}

func main() {
	n := 4 // Number of disks
	TowerOfHanoi(n, "A", "B", "C")
	for i := 0; i < len(Solution); i++ {
		fmt.Println(Solution[i])
	}

}