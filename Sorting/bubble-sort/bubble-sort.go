package main

import "fmt"

func main() {

	arr := []int{7, 12, 9, 11, 3, 2, 11, 6, 14, 1}
	for i := 0; i < len(arr); i++ {
		var swapped bool
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}

	fmt.Print(arr)

}
