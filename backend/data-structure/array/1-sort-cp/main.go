// Sort array terlebih dahulu, kemudian rotasi ke kiri sesuai dengan nilai yang telah ditentukan.
//
// Contoh Sort array:
// Input: [4,5,2,1,3]
// Output: [1,2,3,4,5]

//Contoh RotateLeft:
//Input: 4, [1,2,3,4,5]
//Output: [5,1,2,3,4]

// Explanation RotateLeft:
// untuk melakukan rotasi kiri dengan nilai 4, array mengalami urutan perubahan berikut:
// [1,2,3,4,5] -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4]

package main

import "fmt"

func main() {
	var nums = []int{4, 5, 2, 1, 3}
	arrSorted := Sort(nums)
	fmt.Println(arrSorted)
	rotateLeft := RotateLeft(2, arrSorted)
	fmt.Println(rotateLeft)
}

func Sort(arr []int) []int {
	//return []int{} // TODO: replace this
	swapped != false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+i] {
				//elemen bertukar
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return arr
		}
	}
	return arr
}


func RotateLeft(d int, arr []int) []int {
	// return []int{} // TODO: replace this
	arrLenght := int(len(arr))
	arrResult := arr[d:arrLenght]
	for _, v := range arr[0:d] {
		arrResult = append(arrResult, v)
	}
	return arr
}
