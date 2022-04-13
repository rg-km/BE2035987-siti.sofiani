package main

import "fmt"

// struct + array / slice
type queue struct {
	data [] int
}

func maain (q *queue) Enqueue(newData) int{
	fmt.Println("Enqueue: add data", newData)

}

func Enqueue(arr []int, newData int) []int {
	fmt.Println("Enqueue:", newData)

	return append(arr, newData)
}

func Dequeue(arr []int) []int {
	fmt.Println("Dequeue: ", arr[0])

	return arr[1:]
}