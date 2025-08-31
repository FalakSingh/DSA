package main

import "fmt"

func removeDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	writeIndex := 1

	for readIndex := 1; readIndex < len(arr); readIndex++ {
		if arr[readIndex] != arr[readIndex-1] {
			arr[writeIndex] = arr[readIndex]
			writeIndex++
		}
	}
	
	return arr[:writeIndex]
}
func main() {
	sortedArray := []int{1, 1, 1, 2, 3, 3, 4, 5, 5}
	sortedArray = removeDuplicates(sortedArray)
	fmt.Println(sortedArray)
}
