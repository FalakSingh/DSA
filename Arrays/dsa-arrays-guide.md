# Data Structures & Algorithms: Array Operations Guide

A comprehensive guide to mastering essential array operations in DSA including traversal, in-place updates, prefix/suffix sums, and difference arrays.

## Table of Contents

1. [Array Traversal](#1-array-traversal)
2. [In-Place Updates](#2-in-place-updates)
3. [Prefix Sum Arrays](#3-prefix-sum-arrays)
4. [Suffix Sum Arrays](#4-suffix-sum-arrays)
5. [Difference Arrays](#5-difference-arrays)
6. [Practice Problems](#6-practice-problems)

---

## 1. Array Traversal

Array traversal is the fundamental operation of accessing each element in an array sequentially to perform operations like searching, printing, or processing data.

### 1.1 Basic Traversal Patterns

#### Forward Traversal
**Time Complexity:** O(n)  
**Space Complexity:** O(1)

**JavaScript Example:**
```javascript
function forwardTraversal(arr) {
    for (let i = 0; i < arr.length; i++) {
        console.log(`Index ${i}: ${arr[i]}`);
    }
}

// Alternative using forEach
function forwardTraversalForEach(arr) {
    arr.forEach((element, index) => {
        console.log(`Index ${index}: ${element}`);
    });
}

const numbers = [1, 2, 3, 4, 5];
forwardTraversal(numbers);
```

**Go Example:**
```go
package main

import "fmt"

func forwardTraversal(arr []int) {
    for i := 0; i < len(arr); i++ {
        fmt.Printf("Index %d: %d\n", i, arr[i])
    }
}

// Alternative using range
func forwardTraversalRange(arr []int) {
    for index, value := range arr {
        fmt.Printf("Index %d: %d\n", index, value)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    forwardTraversal(numbers)
}
```

#### Reverse Traversal
**JavaScript Example:**
```javascript
function reverseTraversal(arr) {
    for (let i = arr.length - 1; i >= 0; i--) {
        console.log(`Index ${i}: ${arr[i]}`);
    }
}

const numbers = [1, 2, 3, 4, 5];
reverseTraversal(numbers); // Output: 5, 4, 3, 2, 1
```

**Go Example:**
```go
func reverseTraversal(arr []int) {
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Printf("Index %d: %d\n", i, arr[i])
    }
}
```

### 1.2 Advanced Traversal Techniques

#### Two-Pointer Technique
Used for problems involving pairs, palindromes, or sorted arrays.

**JavaScript Example:**
```javascript
function twoPointerTraversal(arr) {
    let left = 0;
    let right = arr.length - 1;
    
    while (left < right) {
        console.log(`Left: ${arr[left]}, Right: ${arr[right]}`);
        left++;
        right--;
    }
}

// Check if array is palindrome
function isPalindrome(arr) {
    let left = 0;
    let right = arr.length - 1;
    
    while (left < right) {
        if (arr[left] !== arr[right]) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}
```

**Go Example:**
```go
func isPalindrome(arr []int) bool {
    left := 0
    right := len(arr) - 1
    
    for left < right {
        if arr[left] != arr[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```

#### Sliding Window Traversal
Efficient for subarray problems with fixed or variable window sizes.

**JavaScript Example:**
```javascript
// Fixed window size
function maxSubarraySum(arr, k) {
    if (k > arr.length) return null;
    
    
    let maxSum = 0;
    let windowSum = 0;
    
    // Calculate sum of first window
    for (let i = 0; i < k; i++) {
        windowSum += arr[i];
    }
    maxSum = windowSum;
    
    // Slide the window
    for (let i = k; i < arr.length; i++) {

        windowSum = windowSum - arr[i - k] + arr[i];
        maxSum = Math.max(maxSum, windowSum);
    }
    
    return maxSum;
}

const numbers = [1, 4, 2, 10, 23, 3, 1, 0, 20];
console.log(maxSubarraySum(numbers, 4)); // Output: 39
```

**Go Example:**
```go
func maxSubarraySum(arr []int, k int) int {
    if k > len(arr) {
        return -1
    }
    
    maxSum := 0
    windowSum := 0
    
    // Calculate sum of first window
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    maxSum = windowSum
    
    // Slide the window
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    
    return maxSum
}
```

---

## 2. In-Place Updates

In-place algorithms modify the input data structure without using additional space proportional to the input size. They use O(1) extra space.

### 2.1 What Makes an Algorithm "In-Place"?

An in-place algorithm:
- Modifies the original data structure
- Uses only a constant amount of extra space O(1)
- Doesn't create copies of the input data
- May use a few variables for indices, temporary storage

### 2.2 Common In-Place Operations

#### Array Reversal
**JavaScript Example:**
```javascript
function reverseArrayInPlace(arr) {
    let left = 0;
    let right = arr.length - 1;
    
    while (left < right) {
        // Swap elements
        let temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
        
        left++;
        right--;
    }
    return arr;
}

// Using ES6 destructuring for swap
function reverseArrayInPlaceES6(arr) {
    let left = 0;
    let right = arr.length - 1;
    
    while (left < right) {
        [arr[left], arr[right]] = [arr[right], arr[left]];
        left++;
        right--;
    }
    return arr;
}

const numbers = [1, 2, 3, 4, 5];
console.log(reverseArrayInPlace(numbers)); // [5, 4, 3, 2, 1]
```

**Go Example:**
```go
func reverseArrayInPlace(arr []int) {
    left := 0
    right := len(arr) - 1
    
    for left < right {
        // Swap elements
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    reverseArrayInPlace(numbers)
    fmt.Println(numbers) // [5 4 3 2 1]
}
```

#### Remove Duplicates from Sorted Array
**JavaScript Example:**
```javascript
function removeDuplicates(arr) {
    if (arr.length === 0) return 0;
    
    let writeIndex = 1;
    
    for (let readIndex = 1; readIndex < arr.length; readIndex++) {
        if (arr[readIndex] !== arr[readIndex - 1]) {
            arr[writeIndex] = arr[readIndex];
            writeIndex++;
        }
    }
    
    // Trim array to new length
    arr.length = writeIndex;
    return writeIndex;
}

const sortedArray = [1, 1, 2, 2, 3, 4, 4, 5];
const newLength = removeDuplicates(sortedArray);
console.log(sortedArray); // [1, 2, 3, 4, 5]
console.log(newLength);   // 5
```
**Go Example:**
```go
func removeDuplicates(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    
    writeIndex := 1
    
    for readIndex := 1; readIndex < len(arr); readIndex++ {
        if arr[readIndex] != arr[readIndex-1] {
            arr[writeIndex] = arr[readIndex]
            writeIndex++
        }
    }
    
    return writeIndex
}
```

#### Move Zeros to End
**JavaScript Example:**
```javascript
function moveZerosToEnd(arr) {
    let writeIndex = 0;
    
    // Move all non-zero elements to the front
    for (let readIndex = 0; readIndex < arr.length; readIndex++) {
        if (arr[readIndex] !== 0) {
            arr[writeIndex] = arr[readIndex];
            writeIndex++;
        }
    }
    
    // Fill remaining positions with zeros
    while (writeIndex < arr.length) {
        arr[writeIndex] = 0;
        writeIndex++;
    }
    
    return arr;
}

const numbers = [0, 1, 0, 3, 12];
console.log(moveZerosToEnd(numbers)); // [1, 3, 12, 0, 0]
```

**Go Example:**
```go
func moveZerosToEnd(arr []int) {
    writeIndex := 0
    
    // Move all non-zero elements to the front
    for readIndex := 0; readIndex < len(arr); readIndex++ {
        if arr[readIndex] != 0 {
            arr[writeIndex] = arr[readIndex]
            writeIndex++
        }
    }
    
    // Fill remaining positions with zeros
    for writeIndex < len(arr) {
        arr[writeIndex] = 0
        writeIndex++
    }
}
```

---

## 3. Prefix Sum Arrays

Prefix sum is a preprocessing technique that allows for efficient range sum queries. The prefix sum array at index i contains the sum of all elements from index 0 to i.

### 3.1 Basic Prefix Sum Implementation

**JavaScript Example:**
```javascript
function buildPrefixSum(arr) {
    const prefixSum = new Array(arr.length);
    prefixSum[0] = arr[0];
    
    for (let i = 1; i < arr.length; i++) {
        prefixSum[i] = prefixSum[i - 1] + arr[i];
    }
    
    return prefixSum;
}

// Query range sum from index l to r (inclusive)
function rangeSum(prefixSum, l, r) {
    if (l === 0) {
        return prefixSum[r];
    }
    return prefixSum[r] - prefixSum[l - 1];
}

const arr = [1, 2, 3, 4, 5];
const prefixSum = buildPrefixSum(arr);
console.log(prefixSum); // [1, 3, 6, 10, 15]

// Sum of elements from index 1 to 3
console.log(rangeSum(prefixSum, 1, 3)); // 9 (2 + 3 + 4)
```

**Go Example:**
```go
func buildPrefixSum(arr []int) []int {
    n := len(arr)
    prefixSum := make([]int, n)
    prefixSum[0] = arr[0]
    
    for i := 1; i < n; i++ {
        prefixSum[i] = prefixSum[i-1] + arr[i]
    }
    
    return prefixSum
}

func rangeSum(prefixSum []int, l, r int) int {
    if l == 0 {
        return prefixSum[r]
    }
    return prefixSum[r] - prefixSum[l-1]
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    prefixSum := buildPrefixSum(arr)
    fmt.Println(prefixSum) // [1 3 6 10 15]
    
    // Sum of elements from index 1 to 3
    fmt.Println(rangeSum(prefixSum, 1, 3)) // 9
}
```

### 3.2 In-Place Prefix Sum
**JavaScript Example:**
```javascript
function buildPrefixSumInPlace(arr) {
    for (let i = 1; i < arr.length; i++) {
        arr[i] += arr[i - 1];
    }
    return arr;
}

const arr = [1, 2, 3, 4, 5];
buildPrefixSumInPlace(arr);
console.log(arr); // [1, 3, 6, 10, 15]
```

**Go Example:**
```go
func buildPrefixSumInPlace(arr []int) {
    for i := 1; i < len(arr); i++ {
        arr[i] += arr[i-1]
    }
}
```

### 3.3 Applications of Prefix Sum

#### Subarray with Given Sum
**JavaScript Example:**
```javascript
function subarrayWithSum(arr, targetSum) {
    const prefixSumMap = new Map();
    prefixSumMap.set(0, -1); // Handle case where subarray starts from index 0
    
    let currentSum = 0;
    
    for (let i = 0; i < arr.length; i++) {
        currentSum += arr[i];
        
        if (prefixSumMap.has(currentSum - targetSum)) {
            const startIndex = prefixSumMap.get(currentSum - targetSum) + 1;
            return [startIndex, i];
        }
        
        prefixSumMap.set(currentSum, i);
    }
    
    return [-1, -1]; // Not found
}

const arr = [1, 4, 2, 5, 3];
console.log(subarrayWithSum(arr, 7)); // [1, 2] (elements 4, 2, 1 sum to 7)
```

**Go Example:**
```go
func subarrayWithSum(arr []int, targetSum int) [2]int {
    prefixSumMap := make(map[int]int)
    prefixSumMap[0] = -1
    
    currentSum := 0
    
    for i, val := range arr {
        currentSum += val
        
        if startIdx, exists := prefixSumMap[currentSum-targetSum]; exists {
            return [2]int{startIdx + 1, i}
        }
        
        prefixSumMap[currentSum] = i
    }
    
    return [2]int{-1, -1} // Not found
}
```

---

## 4. Suffix Sum Arrays

Suffix sum is similar to prefix sum but calculates the sum from a given index to the end of the array.

### 4.1 Basic Suffix Sum Implementation

**JavaScript Example:**
```javascript
function buildSuffixSum(arr) {
    const n = arr.length;
    const suffixSum = new Array(n);
    suffixSum[n - 1] = arr[n - 1];
    
    for (let i = n - 2; i >= 0; i--) {
        suffixSum[i] = suffixSum[i + 1] + arr[i];
    }
    
    return suffixSum;
}

// Query suffix sum from index i to end
function suffixSumFromIndex(suffixSum, i) {
    return suffixSum[i];
}

const arr = [1, 2, 3, 4, 5];
const suffixSum = buildSuffixSum(arr);
console.log(suffixSum); // [15, 14, 12, 9, 5]

// Sum from index 2 to end
console.log(suffixSumFromIndex(suffixSum, 2)); // 12 (3 + 4 + 5)
```

**Go Example:**
```go
func buildSuffixSum(arr []int) []int {
    n := len(arr)
    suffixSum := make([]int, n)
    suffixSum[n-1] = arr[n-1]
    
    for i := n - 2; i >= 0; i-- {
        suffixSum[i] = suffixSum[i+1] + arr[i]
    }
    
    return suffixSum
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    suffixSum := buildSuffixSum(arr)
    fmt.Println(suffixSum) // [15 14 12 9 5]
}
```

### 4.2 In-Place Suffix Sum
**JavaScript Example:**
```javascript
function buildSuffixSumInPlace(arr) {
    for (let i = arr.length - 2; i >= 0; i--) {
        arr[i] += arr[i + 1];
    }
    return arr;
}

const arr = [1, 2, 3, 4, 5];
buildSuffixSumInPlace(arr);
console.log(arr); // [15, 14, 12, 9, 5]
```

**Go Example:**
```go
func buildSuffixSumInPlace(arr []int) {
    for i := len(arr) - 2; i >= 0; i-- {
        arr[i] += arr[i+1]
    }
}
```

### 4.3 Combined Prefix and Suffix Sums

**JavaScript Example:**
```javascript
function findMaxSubarrayExcludingIndex(arr) {
    const n = arr.length;
    
    // Build prefix sum
    const prefixSum = new Array(n);
    prefixSum[0] = arr[0];
    for (let i = 1; i < n; i++) {
        prefixSum[i] = prefixSum[i - 1] + arr[i];
    }
    
    // Build suffix sum
    const suffixSum = new Array(n);
    suffixSum[n - 1] = arr[n - 1];
    for (let i = n - 2; i >= 0; i--) {
        suffixSum[i] = suffixSum[i + 1] + arr[i];
    }
    
    let maxSum = -Infinity;
    let excludeIndex = -1;
    
    for (let i = 0; i < n; i++) {
        let currentSum = 0;
        
        if (i > 0) currentSum += prefixSum[i - 1];
        if (i < n - 1) currentSum += suffixSum[i + 1];
        
        if (currentSum > maxSum) {
            maxSum = currentSum;
            excludeIndex = i;
        }
    }
    
    return { maxSum, excludeIndex };
}
```

---

## 5. Difference Arrays

Difference arrays are used to efficiently handle multiple range update queries. Instead of updating every element in a range, we update only the boundaries.

### 5.1 Basic Difference Array Implementation

**JavaScript Example:**
```javascript
class DifferenceArray {
    constructor(arr) {
        this.n = arr.length;
        this.diff = new Array(this.n + 1).fill(0);
        this.original = [...arr];
        
        // Initialize difference array
        this.diff[0] = arr[0];
        for (let i = 1; i < this.n; i++) {
            this.diff[i] = arr[i] - arr[i - 1];
        }
    }
    
    // Update range [l, r] by adding value
    rangeUpdate(l, r, value) {
        this.diff[l] += value;
        if (r + 1 < this.n) {
            this.diff[r + 1] -= value;
        }
    }
    
    // Get the final array after all updates
    getFinalArray() {
        const result = new Array(this.n);
        result[0] = this.diff[0];
        
        for (let i = 1; i < this.n; i++) {
            result[i] = result[i - 1] + this.diff[i];
        }
        
        return result;
    }
}

// Example usage
const arr = [1, 2, 3, 4, 5];
const diffArray = new DifferenceArray(arr);

// Add 10 to range [1, 3]
diffArray.rangeUpdate(1, 3, 10);

// Add 5 to range [0, 2]
diffArray.rangeUpdate(0, 2, 5);

console.log(diffArray.getFinalArray()); // [6, 17, 18, 14, 5]
```

**Go Example:**
```go
type DifferenceArray struct {
    diff []int
    n    int
}

func NewDifferenceArray(arr []int) *DifferenceArray {
    n := len(arr)
    diff := make([]int, n+1)
    
    // Initialize difference array
    diff[0] = arr[0]
    for i := 1; i < n; i++ {
        diff[i] = arr[i] - arr[i-1]
    }
    
    return &DifferenceArray{diff: diff, n: n}
}

func (da *DifferenceArray) RangeUpdate(l, r, value int) {
    da.diff[l] += value
    if r+1 < da.n {
        da.diff[r+1] -= value
    }
}

func (da *DifferenceArray) GetFinalArray() []int {
    result := make([]int, da.n)
    result[0] = da.diff[0]
    
    for i := 1; i < da.n; i++ {
        result[i] = result[i-1] + da.diff[i]
    }
    
    return result
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    diffArray := NewDifferenceArray(arr)
    
    // Add 10 to range [1, 3]
    diffArray.RangeUpdate(1, 3, 10)
    
    // Add 5 to range [0, 2]
    diffArray.RangeUpdate(0, 2, 5)
    
    fmt.Println(diffArray.GetFinalArray()) // [6 17 18 14 5]
}
```

### 5.2 Range Update with Multiple Queries

**JavaScript Example:**
```javascript
function rangeUpdates(arr, queries) {
    const n = arr.length;
    const diff = new Array(n + 1).fill(0);
    
    // Initialize difference array
    diff[0] = arr[0];
    for (let i = 1; i < n; i++) {
        diff[i] = arr[i] - arr[i - 1];
    }
    
    // Process all queries
    for (const [l, r, value] of queries) {
        diff[l] += value;
        if (r + 1 < n) {
            diff[r + 1] -= value;
        }
    }
    
    // Build final array
    const result = new Array(n);
    result[0] = diff[0];
    for (let i = 1; i < n; i++) {
        result[i] = result[i - 1] + diff[i];
    }
    
    return result;
}

const arr = [1, 2, 3, 4, 5];
const queries = [
    [1, 3, 10],  // Add 10 to range [1, 3]
    [0, 2, 5],   // Add 5 to range [0, 2]
    [2, 4, -3]   // Subtract 3 from range [2, 4]
];

console.log(rangeUpdates(arr, queries)); // [6, 17, 15, 11, 2]
```

**Go Example:**
```go
func rangeUpdates(arr []int, queries [][3]int) []int {
    n := len(arr)
    diff := make([]int, n+1)
    
    // Initialize difference array
    diff[0] = arr[0]
    for i := 1; i < n; i++ {
        diff[i] = arr[i] - arr[i-1]
    }
    
    // Process all queries
    for _, query := range queries {
        l, r, value := query[0], query[1], query[2]
        diff[l] += value
        if r+1 < n {
            diff[r+1] -= value
        }
    }
    
    // Build final array
    result := make([]int, n)
    result[0] = diff[0]
    for i := 1; i < n; i++ {
        result[i] = result[i-1] + diff[i]
    }
    
    return result
}
```

### 5.3 2D Difference Arrays

For handling range updates in 2D matrices:

**JavaScript Example:**
```javascript
class DifferenceArray2D {
    constructor(matrix) {
        this.rows = matrix.length;
        this.cols = matrix[0].length;
        this.diff = Array(this.rows + 1).fill().map(() => 
            Array(this.cols + 1).fill(0)
        );
    }
    
    // Update rectangle from (r1, c1) to (r2, c2) by adding value
    rangeUpdate(r1, c1, r2, c2, value) {
        this.diff[r1][c1] += value;
        this.diff[r2 + 1][c1] -= value;
        this.diff[r1][c2 + 1] -= value;
        this.diff[r2 + 1][c2 + 1] += value;
    }
    
    // Get the final matrix after all updates
    getFinalMatrix(originalMatrix) {
        const result = Array(this.rows).fill().map(() => 
            Array(this.cols).fill(0)
        );
        
        // Calculate prefix sum
        for (let i = 0; i < this.rows; i++) {
            for (let j = 0; j < this.cols; j++) {
                this.diff[i][j] += 
                    (i > 0 ? this.diff[i - 1][j] : 0) +
                    (j > 0 ? this.diff[i][j - 1] : 0) -
                    (i > 0 && j > 0 ? this.diff[i - 1][j - 1] : 0);
                
                result[i][j] = originalMatrix[i][j] + this.diff[i][j];
            }
        }
        
        return result;
    }
}
```

---

## 6. Practice Problems

### Problem 1: Range Sum Query - Immutable
```javascript
// Given an array, answer multiple range sum queries efficiently
class NumArray {
    constructor(nums) {
        this.prefixSum = [0];
        for (let i = 0; i < nums.length; i++) {
            this.prefixSum[i + 1] = this.prefixSum[i] + nums[i];
        }
    }
    
    sumRange(left, right) {
        return this.prefixSum[right + 1] - this.prefixSum[left];
    }
}
```

### Problem 2: Product of Array Except Self
```javascript
function productExceptSelf(nums) {
    const n = nums.length;
    const result = new Array(n);
    
    // Calculate prefix products
    result[0] = 1;
    for (let i = 1; i < n; i++) {
        result[i] = result[i - 1] * nums[i - 1];
    }
    
    // Calculate suffix products and multiply with prefix
    let suffix = 1;
    for (let i = n - 1; i >= 0; i--) {
        result[i] *= suffix;
        suffix *= nums[i];
    }
    
    return result;
}
```

### Problem 3: Corporate Flight Bookings (Difference Array)
```javascript
function corpFlightBookings(bookings, n) {
    const diff = new Array(n + 1).fill(0);
    
    for (const [first, last, seats] of bookings) {
        diff[first - 1] += seats;
        diff[last] -= seats;
    }
    
    const result = [];
    let currentSum = 0;
    for (let i = 0; i < n; i++) {
        currentSum += diff[i];
        result.push(currentSum);
    }
    
    return result;
}
```

### Problem 4: Maximum Subarray (Kadane's Algorithm)
```javascript
function maxSubArray(nums) {
    let maxSoFar = nums[0];
    let maxEndingHere = nums[0];
    
    for (let i = 1; i < nums.length; i++) {
        maxEndingHere = Math.max(nums[i], maxEndingHere + nums[i]);
        maxSoFar = Math.max(maxSoFar, maxEndingHere);
    }
    
    return maxSoFar;
}
```

## Time and Space Complexities Summary

| Operation | Time Complexity | Space Complexity | Use Case |
|-----------|----------------|------------------|----------|
| Array Traversal | O(n) | O(1) | Basic operations, searching |
| In-place Updates | O(n) | O(1) | Memory-efficient modifications |
| Prefix Sum Build | O(n) | O(n) | Range sum queries |
| Prefix Sum Query | O(1) | O(1) | Fast range sum retrieval |
| Difference Array Update | O(1) per update | O(n) | Multiple range updates |
| Difference Array Build Final | O(n) | O(1) additional | Get final array after updates |

## Key Takeaways

1. **Array Traversal** forms the foundation of most array algorithms
2. **In-place algorithms** save memory by modifying the original data structure
3. **Prefix sums** enable O(1) range sum queries after O(n) preprocessing
4. **Suffix sums** are useful for problems requiring sums from any index to the end
5. **Difference arrays** excel at handling multiple range update operations efficiently

Remember to practice these patterns regularly and understand when to apply each technique based on the problem requirements!