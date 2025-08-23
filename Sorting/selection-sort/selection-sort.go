package main

func main() {
	arr := []int{7, 12, 9, 11, 3, 2, 11, 6, 14, 1}

	minIndex := 0
	for i := range arr {
		for j := range arr {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		println(i, minIndex)
	}
}
