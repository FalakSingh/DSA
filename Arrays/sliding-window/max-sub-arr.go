package main


func maxSubArr(arr []int, k int) int {
	var maxSum int
	for i:=0; i <k; i++ {
		maxSum += arr[i]
	}
	var windowSum int = maxSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}


func main() {
 maxSubArr([]int{2,1,5,1,3,2}, 3)
}
