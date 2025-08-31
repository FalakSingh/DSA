package main

import "fmt"

func zerosToEnd(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	writeIndex := 0
	for readIndex := 0; readIndex < len(arr); readIndex++ {
		if arr[readIndex] != 0 {
			arr[writeIndex] = arr[readIndex]
			writeIndex++
		}
	}

	for writeIndex < len(arr) {
		arr[writeIndex] = 0
		writeIndex++
	}
	return arr
}

func main() {
	array := []int{0, 1, 2, 0, 3, 4}
	array = zerosToEnd(array)
	fmt.Println(array)
}
