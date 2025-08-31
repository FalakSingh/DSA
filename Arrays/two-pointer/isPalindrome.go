package main

import "fmt"

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	text := "racecar"
	result := isPalindrome(text)
	fmt.Printf("%v is a Palindrome: %v", text, result)
}

