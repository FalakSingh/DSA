# Advanced DSA Fundamentals: Complexity Analysis, Mathematical Tools, Recursion & Memory Models

A comprehensive guide to mastering theoretical foundations of Data Structures and Algorithms including asymptotic analysis, mathematical tools, recursion patterns, and memory models.

## Table of Contents

1. [Asymptotic Analysis: Big-O, Θ, Ω](#1-asymptotic-analysis-big-o-θ-ω)
2. [Best/Average/Worst Case Analysis](#2-bestaverageworse-case-analysis)
3. [Amortized Analysis](#3-amortized-analysis)
4. [Mathematical Tools](#4-mathematical-tools)
5. [Recursion Deep Dive](#5-recursion-deep-dive)
6. [Memory Model Basics](#6-memory-model-basics)
7. [Practice Problems & Analysis](#7-practice-problems--analysis)

---

## 1. Asymptotic Analysis: Big-O, Θ, Ω

Asymptotic analysis provides a mathematical framework to analyze algorithm performance as input size approaches infinity, abstracting away machine-specific constants and implementation details.

### 1.1 Big-O Notation (Upper Bound)

**Definition:** f(n) = O(g(n)) if there exist positive constants c and n₀ such that:
0 ≤ f(n) ≤ c·g(n) for all n ≥ n₀

Big-O represents the **worst-case** upper bound of an algorithm's growth rate.

#### Mathematical Representation
```
O(g(n)) = {f(n) : ∃ positive constants c and n₀ such that 0 ≤ f(n) ≤ c·g(n) ∀ n ≥ n₀}
```

#### Common Big-O Complexities (from best to worst)
1. **O(1)** - Constant time
2. **O(log n)** - Logarithmic time  
3. **O(n)** - Linear time
4. **O(n log n)** - Linearithmic time
5. **O(n²)** - Quadratic time
6. **O(n³)** - Cubic time
7. **O(2ⁿ)** - Exponential time
8. **O(n!)** - Factorial time

#### JavaScript Examples

```javascript
// O(1) - Constant Time
function getFirstElement(arr) {
    return arr[0]; // Single operation regardless of array size
}

// O(n) - Linear Time
function findElement(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) {
            return i;
        }
    }
    return -1;
}

// O(n²) - Quadratic Time
function bubbleSort(arr) {
    const n = arr.length;
    for (let i = 0; i < n; i++) {          // Outer loop: n iterations
        for (let j = 0; j < n - i - 1; j++) { // Inner loop: up to n iterations
            if (arr[j] > arr[j + 1]) {
                [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
            }
        }
    }
    return arr;
}

// O(log n) - Logarithmic Time
function binarySearch(arr, target) {
    let left = 0, right = arr.length - 1;
    
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) return mid;
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}
```

#### Go Examples

```go
package main

import "fmt"

// O(1) - Constant Time
func getFirstElement(arr []int) int {
    if len(arr) > 0 {
        return arr[0]
    }
    return -1
}

// O(n) - Linear Time
func findElement(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i
        }
    }
    return -1
}

// O(n²) - Quadratic Time
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// O(log n) - Logarithmic Time
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

### 1.2 Big-Omega Notation (Lower Bound)

**Definition:** f(n) = Ω(g(n)) if there exist positive constants c and n₀ such that:
0 ≤ c·g(n) ≤ f(n) for all n ≥ n₀

Big-Omega represents the **best-case** lower bound of an algorithm's growth rate.

#### Mathematical Representation
```
Ω(g(n)) = {f(n) : ∃ positive constants c and n₀ such that 0 ≤ c·g(n) ≤ f(n) ∀ n ≥ n₀}
```

#### Examples

**Insertion Sort:**
- Best case: Ω(n) - when array is already sorted
- Worst case: O(n²) - when array is reverse sorted

```javascript
function insertionSort(arr) {
    for (let i = 1; i < arr.length; i++) {
        let key = arr[i];
        let j = i - 1;
        
        // Best case: already sorted, inner loop runs 0 times = Ω(n)
        // Worst case: reverse sorted, inner loop runs i times = O(n²)
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j--;
        }
        arr[j + 1] = key;
    }
    return arr;
}
```

### 1.3 Big-Theta Notation (Tight Bound)

**Definition:** f(n) = Θ(g(n)) if and only if f(n) = O(g(n)) AND f(n) = Ω(g(n))

This means: ∃ positive constants c₁, c₂, and n₀ such that:
0 ≤ c₁·g(n) ≤ f(n) ≤ c₂·g(n) for all n ≥ n₀

#### Mathematical Representation
```
Θ(g(n)) = {f(n) : ∃ positive constants c₁, c₂ and n₀ such that 0 ≤ c₁·g(n) ≤ f(n) ≤ c₂·g(n) ∀ n ≥ n₀}
```

#### Examples

**Merge Sort:**
- Always Θ(n log n) regardless of input (best, average, worst case)

```javascript
function mergeSort(arr) {
    if (arr.length <= 1) return arr;
    
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));     // T(n/2)
    const right = mergeSort(arr.slice(mid));       // T(n/2)
    
    return merge(left, right);                     // Θ(n)
}

function merge(left, right) {
    const result = [];
    let i = 0, j = 0;
    
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            result.push(left[i++]);
        } else {
            result.push(right[j++]);
        }
    }
    
    return result.concat(left.slice(i)).concat(right.slice(j));
}

// Recurrence: T(n) = 2T(n/2) + Θ(n)
// Solution: T(n) = Θ(n log n)
```

### 1.4 Notation Comparison Summary

| Notation | Meaning | Analogy | Use Case |
|----------|---------|---------|----------|
| **O(g(n))** | Upper bound (≤) | "At most" | Worst-case analysis |
| **Ω(g(n))** | Lower bound (≥) | "At least" | Best-case analysis |
| **Θ(g(n))** | Tight bound (=) | "Exactly" | Average-case analysis |

---

## 2. Best/Average/Worst Case Analysis

### 2.1 Understanding Different Cases

#### Best Case
- Most favorable input for the algorithm
- Lower bound on running time
- Represented by Ω notation

#### Average Case  
- Expected performance over all possible inputs
- Usually requires probabilistic analysis
- Represented by Θ notation

#### Worst Case
- Least favorable input for the algorithm
- Upper bound on running time
- Represented by O notation
- Most commonly analyzed in practice

### 2.2 Comprehensive Example: Quick Sort

```javascript
function quickSort(arr, low = 0, high = arr.length - 1) {
    if (low < high) {
        const pi = partition(arr, low, high);
        
        quickSort(arr, low, pi - 1);     // Left subarray
        quickSort(arr, pi + 1, high);   // Right subarray
    }
    return arr;
}

function partition(arr, low, high) {
    const pivot = arr[high];  // Choose last element as pivot
    let i = low - 1;          // Index of smaller element
    
    for (let j = low; j < high; j++) {
        if (arr[j] < pivot) {
            i++;
            [arr[i], arr[j]] = [arr[j], arr[i]];
        }
    }
    [arr[i + 1], arr[high]] = [arr[high], arr[i + 1]];
    return i + 1;
}
```

**Quick Sort Analysis:**

1. **Best Case:** Ω(n log n)
   - Pivot always divides array into two equal halves
   - Recurrence: T(n) = 2T(n/2) + Θ(n)

2. **Average Case:** Θ(n log n)
   - On average, pivot provides reasonably balanced partitions
   - Expected depth of recursion is log n

3. **Worst Case:** O(n²)
   - Pivot is always the smallest or largest element
   - Recurrence: T(n) = T(n-1) + T(0) + Θ(n) = T(n-1) + Θ(n)

#### Go Implementation with Analysis

```go
func quickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        
        quickSort(arr, low, pi-1)    // Left subarray
        quickSort(arr, pi+1, high)   // Right subarray
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]  // Choose last element as pivot
    i := low - 1        // Index of smaller element
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

// Best Case Input: [4, 2, 6, 1, 3, 5, 7] - balanced partitions
// Worst Case Input: [1, 2, 3, 4, 5, 6, 7] - already sorted
```

### 2.3 Case Analysis Examples

#### Linear Search
```javascript
function linearSearch(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) {
            return i;
        }
    }
    return -1;
}

// Best Case: Ω(1) - element found at first position
// Average Case: Θ(n/2) = Θ(n) - element found at middle
// Worst Case: O(n) - element at last position or not found
```

#### Binary Search Tree Operations
```javascript
class TreeNode {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

function search(root, target) {
    if (!root || root.val === target) {
        return root;
    }
    
    if (target < root.val) {
        return search(root.left, target);
    } else {
        return search(root.right, target);
    }
}

// Best Case: Ω(1) - target is at root
// Average Case: Θ(log n) - balanced tree
// Worst Case: O(n) - degenerate tree (linked list)
```

---

## 3. Amortized Analysis

Amortized analysis provides the average performance of each operation in the worst case over a sequence of operations, even when individual operations may be expensive.

### 3.1 Key Concepts

**Amortized Cost** = Total Cost of n Operations / n

Unlike average-case analysis (which considers typical inputs), amortized analysis considers worst-case sequences but averages the cost over multiple operations.

### 3.2 Methods of Amortized Analysis

#### 3.2.1 Aggregate Method
Calculate total cost of n operations and divide by n.

#### 3.2.2 Accounting Method (Banker's Method)
Assign different charges to operations, storing credits for future expensive operations.

#### 3.2.3 Potential Method
Use a potential function that maps data structure states to real numbers.

### 3.3 Dynamic Arrays (Resizable Arrays)

**The Classic Example:** Dynamic arrays that double in size when full.

#### JavaScript Implementation

```javascript
class DynamicArray {
    constructor() {
        this.data = new Array(1);
        this.size = 0;
        this.capacity = 1;
    }
    
    push(element) {
        // If array is full, double the capacity
        if (this.size === this.capacity) {
            this.resize();
        }
        
        this.data[this.size] = element;
        this.size++;
    }
    
    resize() {
        const oldCapacity = this.capacity;
        this.capacity *= 2;  // Double the capacity
        
        const newData = new Array(this.capacity);
        
        // Copy all elements - O(n) operation
        for (let i = 0; i < this.size; i++) {
            newData[i] = this.data[i];
        }
        
        this.data = newData;
        console.log(`Resized from ${oldCapacity} to ${this.capacity}`);
    }
    
    get(index) {
        if (index < 0 || index >= this.size) {
            throw new Error("Index out of bounds");
        }
        return this.data[index];
    }
    
    getSize() {
        return this.size;
    }
}

// Usage and analysis
const dynArray = new DynamicArray();

// Adding elements: 1, 2, 3, 4, 5, 6, 7, 8, 9
for (let i = 1; i <= 9; i++) {
    dynArray.push(i);
}

/*
Sequence of operations:
- Insert 1: resize to capacity 1, cost = 1
- Insert 2: resize to capacity 2, copy 1 element, cost = 2  
- Insert 3: resize to capacity 4, copy 2 elements, cost = 3
- Insert 4: no resize, cost = 1
- Insert 5: resize to capacity 8, copy 4 elements, cost = 5
- Insert 6, 7, 8: no resize, cost = 1 each
- Insert 9: resize to capacity 16, copy 8 elements, cost = 9

Total cost = 1 + 2 + 3 + 1 + 5 + 1 + 1 + 1 + 9 = 24
Amortized cost per operation = 24/9 ≈ 2.67 = O(1)
*/
```

#### Go Implementation

```go
package main

import "fmt"

type DynamicArray struct {
    data     []int
    size     int
    capacity int
}

func NewDynamicArray() *DynamicArray {
    return &DynamicArray{
        data:     make([]int, 1),
        size:     0,
        capacity: 1,
    }
}

func (da *DynamicArray) Push(element int) {
    // If array is full, double the capacity
    if da.size == da.capacity {
        da.resize()
    }
    
    da.data[da.size] = element
    da.size++
}

func (da *DynamicArray) resize() {
    oldCapacity := da.capacity
    da.capacity *= 2
    
    newData := make([]int, da.capacity)
    
    // Copy all elements - O(n) operation
    for i := 0; i < da.size; i++ {
        newData[i] = da.data[i]
    }
    
    da.data = newData
    fmt.Printf("Resized from %d to %d\n", oldCapacity, da.capacity)
}

func (da *DynamicArray) Get(index int) (int, error) {
    if index < 0 || index >= da.size {
        return 0, fmt.Errorf("index out of bounds")
    }
    return da.data[index], nil
}

func (da *DynamicArray) Size() int {
    return da.size
}
```

#### Amortized Analysis Using Potential Method

**Potential Function:** Φ(D) = 2×size - capacity

- When array is exactly full: Φ(D) = 2×capacity - capacity = capacity
- When array is half full: Φ(D) = 2×(capacity/2) - capacity = 0

**Analysis:**
1. **Normal insertion** (no resize): actual cost = 1, potential change = 2
   - Amortized cost = 1 + 2 = 3

2. **Insertion with resize**: actual cost = size + 1, potential change = 0 - size = -size
   - Amortized cost = (size + 1) + (-size) = 1

**Result:** Every operation has O(1) amortized cost.

### 3.4 Hash Tables with Chaining

```javascript
class HashTable {
    constructor(initialSize = 8) {
        this.buckets = new Array(initialSize).fill(null).map(() => []);
        this.size = 0;
        this.capacity = initialSize;
        this.loadFactorThreshold = 0.75;
    }
    
    hash(key) {
        let hash = 0;
        for (let i = 0; i < key.length; i++) {
            hash = (hash + key.charCodeAt(i) * i) % this.capacity;
        }
        return hash;
    }
    
    put(key, value) {
        const index = this.hash(key);
        const bucket = this.buckets[index];
        
        // Check if key already exists
        for (let i = 0; i < bucket.length; i++) {
            if (bucket[i][0] === key) {
                bucket[i][1] = value;  // Update existing
                return;
            }
        }
        
        // Add new key-value pair
        bucket.push([key, value]);
        this.size++;
        
        // Check if resize is needed
        if (this.size > this.capacity * this.loadFactorThreshold) {
            this.resize();
        }
    }
    
    resize() {
        const oldBuckets = this.buckets;
        this.capacity *= 2;
        this.buckets = new Array(this.capacity).fill(null).map(() => []);
        this.size = 0;
        
        // Rehash all elements - O(n) operation
        for (const bucket of oldBuckets) {
            for (const [key, value] of bucket) {
                this.put(key, value);  // This won't trigger another resize
            }
        }
    }
    
    get(key) {
        const index = this.hash(key);
        const bucket = this.buckets[index];
        
        for (const [k, v] of bucket) {
            if (k === key) return v;
        }
        return null;
    }
}

// Amortized Analysis:
// - Normal operations: O(1) average
// - Resize operations: O(n) but infrequent
// - Amortized cost: O(1) per operation
```

### 3.5 Union-Find (Disjoint Set Union)

```javascript
class UnionFind {
    constructor(n) {
        this.parent = Array.from({length: n}, (_, i) => i);
        this.rank = new Array(n).fill(0);
    }
    
    find(x) {
        if (this.parent[x] !== x) {
            // Path compression - makes future finds faster
            this.parent[x] = this.find(this.parent[x]);
        }
        return this.parent[x];
    }
    
    union(x, y) {
        const rootX = this.find(x);
        const rootY = this.find(y);
        
        if (rootX === rootY) return;
        
        // Union by rank - attach smaller tree under larger tree
        if (this.rank[rootX] < this.rank[rootY]) {
            this.parent[rootX] = rootY;
        } else if (this.rank[rootX] > this.rank[rootY]) {
            this.parent[rootY] = rootX;
        } else {
            this.parent[rootY] = rootX;
            this.rank[rootX]++;
        }
    }
    
    connected(x, y) {
        return this.find(x) === this.find(y);
    }
}

// Amortized Analysis:
// - Individual find/union operations can be O(log n)
// - With path compression + union by rank: O(α(n)) amortized
// - α(n) is inverse Ackermann function, practically constant
```

---

## 4. Mathematical Tools

### 4.1 Logarithms

#### Definition and Properties

**Definition:** log_b(x) = y if and only if b^y = x

#### Fundamental Properties

```javascript
// 1. Product Rule: log(xy) = log(x) + log(y)
function logProduct(x, y) {
    return Math.log(x) + Math.log(y); // equals Math.log(x * y)
}

// 2. Quotient Rule: log(x/y) = log(x) - log(y)
function logQuotient(x, y) {
    return Math.log(x) - Math.log(y); // equals Math.log(x / y)
}

// 3. Power Rule: log(x^n) = n * log(x)
function logPower(x, n) {
    return n * Math.log(x); // equals Math.log(Math.pow(x, n))
}

// 4. Change of Base Formula: log_b(x) = log_c(x) / log_c(b)
function logBase(x, base) {
    return Math.log(x) / Math.log(base);
}

// 5. Base Conversion
const log2 = (x) => Math.log(x) / Math.log(2);
const log10 = (x) => Math.log10(x);
const ln = (x) => Math.log(x); // Natural logarithm (base e)
```

#### Common Logarithm Identities

| Identity | Value | Application |
|----------|--------|-------------|
| log₂(n) | Height of balanced binary tree | Tree algorithms |
| log₁₀(n) | Number of decimal digits | Input/output analysis |
| ln(n) | Natural logarithm | Continuous mathematics |
| log(1) = 0 | Any base | Base case in recursion |
| log(b) = 1 | Same base b | Logarithm definition |

#### Applications in Algorithm Analysis

```javascript
// Binary Search - O(log n) because we halve the search space
function binarySearchAnalysis(n) {
    // Maximum comparisons = ceil(log₂(n))
    return Math.ceil(log2(n));
}

// Height of complete binary tree with n nodes
function treeHeight(n) {
    return Math.floor(log2(n));
}

// Heap operations complexity
function heapOperations(n) {
    return {
        insert: log2(n),      // O(log n)
        extractMax: log2(n),  // O(log n)
        build: n              // O(n) - not O(n log n)!
    };
}
```

### 4.2 Series and Summations

#### Arithmetic Series
**Sum:** 1 + 2 + 3 + ... + n = n(n+1)/2 = Θ(n²)

```javascript
function arithmeticSum(n) {
    return n * (n + 1) / 2;
}

// Application: Nested loop analysis
function nestedLoopComplexity(n) {
    let operations = 0;
    for (let i = 1; i <= n; i++) {
        for (let j = 1; j <= i; j++) {
            operations++; // This runs 1+2+3+...+n = n(n+1)/2 times
        }
    }
    return operations; // Θ(n²)
}
```

#### Geometric Series
**Sum:** 1 + r + r² + ... + r^n = (r^(n+1) - 1)/(r - 1) for r ≠ 1

```javascript
function geometricSum(r, n) {
    if (r === 1) return n + 1;
    return (Math.pow(r, n + 1) - 1) / (r - 1);
}

// Application: Binary tree node count
function binaryTreeNodes(height) {
    // Total nodes = 2⁰ + 2¹ + 2² + ... + 2^h = 2^(h+1) - 1
    return Math.pow(2, height + 1) - 1;
}

// Application: Dynamic array resize cost
function dynamicArrayCost(n) {
    // Resize costs: 1 + 2 + 4 + 8 + ... + n = 2n - 1 = O(n)
    let cost = 0;
    let size = 1;
    while (size <= n) {
        cost += size;
        size *= 2;
    }
    return cost;
}
```

#### Harmonic Series
**Sum:** H_n = 1 + 1/2 + 1/3 + ... + 1/n ≈ ln(n) + γ (Euler's constant)

```javascript
function harmonicSum(n) {
    let sum = 0;
    for (let i = 1; i <= n; i++) {
        sum += 1 / i;
    }
    return sum; // Approximately ln(n) + 0.577...
}

// Application: Analysis of quicksort average case
function quicksortAverageComparisons(n) {
    // Expected comparisons ≈ 2n * H_n ≈ 2n ln(n)
    return 2 * n * (Math.log(n) + 0.5772); // Approximate
}
```

### 4.3 Growth Rates and Dominance

#### Growth Rate Hierarchy (from slowest to fastest)

```javascript
function compareGrowthRates(n) {
    return {
        constant: 1,                    // O(1)
        logarithmic: Math.log2(n),     // O(log n)
        linear: n,                     // O(n)
        linearithmic: n * Math.log2(n), // O(n log n)
        quadratic: n * n,              // O(n²)
        cubic: n * n * n,              // O(n³)
        exponential: Math.pow(2, n),   // O(2ⁿ)
        factorial: factorial(n)         // O(n!)
    };
}

function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}

// Dominance rules
function analyzeComplexity(n) {
    // When comparing: f(n) + g(n) = O(max(f(n), g(n)))
    const linear = n;
    const quadratic = n * n;
    const logarithmic = Math.log2(n);
    
    // n² + n + log n = O(n²) because n² dominates
    return Math.max(quadratic, linear, logarithmic); 
}
```

#### Little-o and Little-omega Notations

```javascript
// f(n) = o(g(n)) means f(n) grows strictly slower than g(n)
// f(n) = ω(g(n)) means f(n) grows strictly faster than g(n)

function littleOExample() {
    // n = o(n²) because lim(n→∞) n/n² = lim(n→∞) 1/n = 0
    // log n = o(n) because logarithms grow slower than linear
    
    const examples = {
        "log n is o(n)": true,
        "n is o(n log n)": true,
        "n log n is o(n²)": true,
        "n² is NOT o(n²)": true  // Same growth rate
    };
    
    return examples;
}
```

---

## 5. Recursion Deep Dive

### 5.1 Anatomy of Recursion

#### The Three Laws of Recursion
1. **Base Case:** Condition that stops recursion
2. **Recursive Case:** Function calls itself with modified parameters
3. **Progress:** Each call moves toward the base case

```javascript
function recursiveTemplate(n, ...args) {
    // Law 1: Base Case - prevents infinite recursion
    if (baseCondition(n)) {
        return baseValue;
    }
    
    // Law 2: Recursive Case - function calls itself
    // Law 3: Progress - parameter moves toward base case
    return combineResults(
        processCurrentLevel(n, args),
        recursiveTemplate(modifyParameter(n), ...newArgs)
    );
}
```

### 5.2 Base Case Design Patterns

#### Single Base Case
```javascript
function factorial(n) {
    if (n <= 1) return 1;        // Base case
    return n * factorial(n - 1); // Recursive case
}
```

#### Multiple Base Cases
```javascript
function fibonacci(n) {
    if (n <= 0) return 0;        // Base case 1
    if (n === 1) return 1;       // Base case 2
    return fibonacci(n - 1) + fibonacci(n - 2); // Recursive case
}
```

#### Go Examples
```go
// Single base case
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// Multiple base cases
func fibonacci(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

### 5.3 Recurrence Relations

#### Definition
A recurrence relation defines a sequence where each term depends on previous terms.

**General Form:** T(n) = aT(n/b) + f(n)

#### Common Patterns

**1. Linear Recurrences**
```javascript
// T(n) = T(n-1) + c
function linearRecurrence(n) {
    if (n === 0) return 1;
    return linearRecurrence(n - 1) + 1; // T(n) = O(n)
}
```

**2. Divide and Conquer Recurrences**
```javascript
// T(n) = 2T(n/2) + n
function mergeSort(arr) {
    if (arr.length <= 1) return arr;
    
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));     // T(n/2)
    const right = mergeSort(arr.slice(mid));       // T(n/2)
    
    return merge(left, right);                     // O(n)
}
```

**3. Tree Recurrences**
```javascript
// T(n) = 2T(n-1) + c (binary recursion)
function fibonacciNaive(n) {
    if (n <= 1) return n;
    return fibonacciNaive(n - 1) + fibonacciNaive(n - 2); // T(n) = O(2^n)
}
```

### 5.4 Master Theorem

**For recurrences of the form:** T(n) = aT(n/b) + f(n)

Where a ≥ 1, b > 1, and f(n) is asymptotically positive.

#### The Three Cases

**Case 1:** If f(n) = O(n^(log_b(a) - ε)) for some ε > 0
- **Solution:** T(n) = Θ(n^(log_b(a)))
- **Intuition:** Recursive calls dominate

**Case 2:** If f(n) = Θ(n^(log_b(a)) × log^k(n)) for k ≥ 0  
- **Solution:** T(n) = Θ(n^(log_b(a)) × log^(k+1)(n))
- **Intuition:** Work is evenly distributed

**Case 3:** If f(n) = Ω(n^(log_b(a) + ε)) for some ε > 0, and af(n/b) ≤ cf(n) for some c < 1
- **Solution:** T(n) = Θ(f(n))
- **Intuition:** Current level dominates

#### Master Theorem Examples

```javascript
// Example 1: Merge Sort
// T(n) = 2T(n/2) + n
// a=2, b=2, f(n)=n
// n^(log₂2) = n¹ = n, so f(n) = Θ(n)
// Case 2: T(n) = Θ(n log n)

// Example 2: Binary Search
// T(n) = T(n/2) + 1
// a=1, b=2, f(n)=1
// n^(log₂1) = n⁰ = 1, so f(n) = Θ(1)
// Case 2: T(n) = Θ(log n)

function masterTheoremAnalysis(a, b, fComplexity, n) {
    const logBA = Math.log(a) / Math.log(b);
    const threshold = Math.pow(n, logBA);
    
    console.log(`a=${a}, b=${b}, n^(log_b a) = n^${logBA.toFixed(2)} ≈ ${threshold.toFixed(2)}`);
    
    if (fComplexity < threshold) {
        console.log(`Case 1: T(n) = Θ(n^${logBA.toFixed(2)})`);
    } else if (Math.abs(fComplexity - threshold) < 0.1) {
        console.log(`Case 2: T(n) = Θ(n^${logBA.toFixed(2)} × log n)`);
    } else {
        console.log(`Case 3: T(n) = Θ(f(n))`);
    }
}
```

### 5.5 Recursion Tree Method

Visual method to solve recurrence relations by drawing the recursion tree.

```javascript
// Example: T(n) = 2T(n/2) + n
function recursionTreeAnalysis(n, level = 0) {
    const indent = "  ".repeat(level);
    
    if (n <= 1) {
        console.log(`${indent}T(${n}) = 1`);
        return 1;
    }
    
    console.log(`${indent}T(${n}) = 2T(${Math.floor(n/2)}) + ${n}`);
    
    const leftCost = recursionTreeAnalysis(Math.floor(n/2), level + 1);
    const rightCost = recursionTreeAnalysis(Math.floor(n/2), level + 1);
    
    const totalCost = leftCost + rightCost + n;
    console.log(`${indent}Total cost at level ${level}: ${totalCost}`);
    
    return totalCost;
}

/*
Tree visualization for T(n) = 2T(n/2) + n:

Level 0:    T(n) cost = n
           /          \
Level 1: T(n/2)     T(n/2)    cost = n/2 + n/2 = n
        /    \      /    \
Level 2: T(n/4) T(n/4) T(n/4) T(n/4)  cost = 4×(n/4) = n

Total levels = log₂(n)
Cost per level = n
Total cost = n × log₂(n) = O(n log n)
*/
```

### 5.6 Optimizing Recursion

#### Tail Recursion
```javascript
// Non-tail recursive (accumulates on return)
function factorialNormal(n) {
    if (n <= 1) return 1;
    return n * factorialNormal(n - 1); // Operation after recursive call
}

// Tail recursive (accumulates on the way down)
function factorialTail(n, accumulator = 1) {
    if (n <= 1) return accumulator;
    return factorialTail(n - 1, n * accumulator); // Recursive call is last operation
}
```

#### Memoization (Dynamic Programming)
```javascript
function fibonacciMemo(n, memo = new Map()) {
    if (n <= 1) return n;
    
    if (memo.has(n)) {
        return memo.get(n); // Return cached result
    }
    
    const result = fibonacciMemo(n - 1, memo) + fibonacciMemo(n - 2, memo);
    memo.set(n, result); // Cache the result
    
    return result;
}

// Time complexity: O(n) instead of O(2^n)
// Space complexity: O(n) for memoization table
```

#### Go Memoization Example
```go
func fibonacciMemo(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
    memo[n] = result
    
    return result
}

func main() {
    memo := make(map[int]int)
    fmt.Println(fibonacciMemo(40, memo)) // Much faster than naive recursion
}
```

---

## 6. Memory Model Basics

### 6.1 Stack vs Heap Memory

#### Stack Memory
- **Purpose:** Stores local variables, function parameters, return addresses
- **Characteristics:** 
  - LIFO (Last In, First Out) structure
  - Fast allocation/deallocation
  - Limited size (typically MB range)
  - Thread-local (each thread has own stack)
  - Automatically managed

#### Heap Memory  
- **Purpose:** Stores dynamically allocated objects
- **Characteristics:**
  - Random access memory pool
  - Slower allocation/deallocation
  - Large size (limited by system RAM)
  - Shared across threads
  - Manual or garbage-collected management

#### Memory Layout Visualization

```
High Address
┌─────────────────┐
│     Stack       │ ← Local variables, function calls
│       ↓         │ ← Grows downward
├─────────────────┤
│                 │ ← Free space
│                 │
│       ↑         │ ← Grows upward  
│     Heap        │ ← Dynamically allocated objects
├─────────────────┤
│   Data Segment  │ ← Global/static variables
├─────────────────┤
│  Code Segment   │ ← Program instructions
└─────────────────┘
Low Address
```

#### JavaScript Memory Example

```javascript
function memoryExample() {
    // Stack allocated
    let primitive = 42;           // Number stored on stack
    let localString = "hello";    // Reference on stack, value may be interned
    
    // Heap allocated
    let object = {                // Reference on stack, object on heap
        name: "Alice",
        age: 30,
        hobbies: ["reading", "coding"] // Array also on heap
    };
    
    let array = new Array(1000);  // Large array on heap
    
    function innerFunction(param) { // New stack frame created
        let localVar = param * 2;  // Stack allocated
        return localVar;           // Stack frame destroyed on return
    }
    
    return innerFunction(primitive);
} // All local variables and stack frame destroyed here

// Object and array remain on heap until garbage collected
```

#### Go Memory Example

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func memoryExample() *Person {
    // Stack allocated (in Go, compiler decides)
    localInt := 42
    
    // This might be stack or heap allocated (escape analysis)
    localPerson := Person{Name: "Alice", Age: 30}
    
    // This will be heap allocated (escapes function scope)
    heapPerson := &Person{Name: "Bob", Age: 25}
    
    // Slice header on stack, underlying array on heap
    slice := make([]int, 1000)
    
    fmt.Println(localInt, localPerson.Name, len(slice))
    
    return heapPerson // Returns pointer to heap-allocated object
}

func main() {
    person := memoryExample()
    fmt.Println(person.Name) // Accessing heap-allocated memory
}
```

### 6.2 Reference vs Value Semantics

#### Value Semantics
- **Definition:** Copy the actual value when assigning or passing to functions
- **Behavior:** Independent copies, modifications don't affect original
- **Memory:** May involve copying data

#### Reference Semantics  
- **Definition:** Copy the reference/pointer when assigning or passing
- **Behavior:** Shared data, modifications affect all references
- **Memory:** Efficient for large objects

#### JavaScript Examples

```javascript
// Value semantics with primitives
function valueSemantics() {
    let a = 5;
    let b = a;    // Copy the value
    b = 10;       // Doesn't affect 'a'
    
    console.log(a, b); // Output: 5, 10
}

// Reference semantics with objects
function referenceSemantics() {
    let obj1 = { x: 5 };
    let obj2 = obj1;     // Copy the reference
    obj2.x = 10;         // Modifies the shared object
    
    console.log(obj1.x, obj2.x); // Output: 10, 10
}

// Function parameter behavior
function modifyPrimitive(value) {
    value = 100;         // Doesn't affect original
}

function modifyObject(obj) {
    obj.x = 100;         // Modifies original object
}

function modifyReference(obj) {
    obj = { x: 200 };    // Doesn't affect original reference
}

let num = 42;
let object = { x: 5 };

modifyPrimitive(num);
console.log(num);        // Output: 42 (unchanged)

modifyObject(object);
console.log(object.x);   // Output: 100 (changed)

let obj2 = { x: 5 };
modifyReference(obj2);
console.log(obj2.x);     // Output: 5 (unchanged)
```

#### Go Examples

```go
// Value semantics (default for structs)
type Point struct {
    X, Y int
}

func modifyPointValue(p Point) {
    p.X = 100  // Modifies copy, not original
}

func modifyPointReference(p *Point) {
    p.X = 100  // Modifies original through pointer
}

func main() {
    // Value semantics
    p1 := Point{X: 5, Y: 10}
    p2 := p1           // Copy the struct
    p2.X = 20          // Doesn't affect p1
    
    fmt.Println(p1.X, p2.X) // Output: 5, 20
    
    // Reference semantics with pointers
    ptr1 := &Point{X: 5, Y: 10}
    ptr2 := ptr1       // Copy the pointer
    ptr2.X = 20        // Modifies the shared object
    
    fmt.Println(ptr1.X, ptr2.X) // Output: 20, 20
    
    // Function calls
    p := Point{X: 5, Y: 10}
    modifyPointValue(p)
    fmt.Println(p.X)   // Output: 5 (unchanged)
    
    modifyPointReference(&p)
    fmt.Println(p.X)   // Output: 100 (changed)
}

// Slices have reference semantics for underlying array
func modifySlice(s []int) {
    s[0] = 100         // Modifies underlying array
    s = append(s, 999) // May create new array, doesn't affect original slice
}
```

### 6.3 Mutability

#### Immutable Objects
- Cannot be modified after creation
- Thread-safe by default
- Functional programming paradigm

#### Mutable Objects
- Can be modified after creation
- May require synchronization in concurrent programs
- Imperative programming paradigm

```javascript
// Immutable approach (JavaScript)
const immutableExample = {
    // Create new object instead of modifying
    updatePerson: (person, newAge) => {
        return { ...person, age: newAge }; // Spread creates new object
    },
    
    // Pure function - no side effects
    addToArray: (arr, element) => {
        return [...arr, element]; // Creates new array
    }
};

// Usage
const person = { name: "Alice", age: 25 };
const updatedPerson = immutableExample.updatePerson(person, 26);
// person is unchanged, updatedPerson is new object

// Mutable approach (traditional)
const mutableExample = {
    updatePerson: (person, newAge) => {
        person.age = newAge; // Modifies original object
    },
    
    addToArray: (arr, element) => {
        arr.push(element); // Modifies original array
    }
};
```

### 6.4 Shallow vs Deep Copy

#### Shallow Copy
- Copies only the first level of the object
- Nested objects are still referenced (not copied)
- Fast and memory efficient

#### Deep Copy  
- Recursively copies all levels of nested objects
- Completely independent copies
- Slower and more memory intensive

```javascript
// Shallow copy examples
function shallowCopyExamples() {
    const original = {
        name: "Alice",
        address: {
            city: "New York",
            country: "USA"
        },
        hobbies: ["reading", "coding"]
    };
    
    // Method 1: Object.assign()
    const shallow1 = Object.assign({}, original);
    
    // Method 2: Spread operator
    const shallow2 = { ...original };
    
    // Method 3: Array.from() for arrays
    const arr = [1, [2, 3], 4];
    const shallowArr = Array.from(arr);
    
    // Problem with shallow copy
    shallow1.address.city = "San Francisco"; // Affects original!
    console.log(original.address.city); // "San Francisco"
    
    return { shallow1, shallow2, shallowArr };
}

// Deep copy examples
function deepCopyExamples() {
    const original = {
        name: "Alice",
        address: {
            city: "New York", 
            country: "USA"
        },
        hobbies: ["reading", "coding"]
    };
    
    // Method 1: JSON (limited - no functions, dates become strings)
    const deep1 = JSON.parse(JSON.stringify(original));
    
    // Method 2: Recursive deep copy
    function deepCopy(obj) {
        if (obj === null || typeof obj !== "object") {
            return obj;
        }
        
        if (obj instanceof Date) {
            return new Date(obj.getTime());
        }
        
        if (obj instanceof Array) {
            return obj.map(item => deepCopy(item));
        }
        
        if (typeof obj === "object") {
            const copy = {};
            Object.keys(obj).forEach(key => {
                copy[key] = deepCopy(obj[key]);
            });
            return copy;
        }
    }
    
    const deep2 = deepCopy(original);
    
    // Test deep copy
    deep2.address.city = "Los Angeles"; // Doesn't affect original
    console.log(original.address.city); // Still "New York"
    
    return { deep1, deep2 };
}

// Performance comparison
function copyPerformanceTest(obj, iterations = 10000) {
    console.time("Shallow Copy");
    for (let i = 0; i < iterations; i++) {
        const copy = { ...obj };
    }
    console.timeEnd("Shallow Copy");
    
    console.time("Deep Copy (JSON)");
    for (let i = 0; i < iterations; i++) {
        const copy = JSON.parse(JSON.stringify(obj));
    }
    console.timeEnd("Deep Copy (JSON)");
}
```

#### Go Copy Examples

```go
package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name    string
    Address *Address
}

type Address struct {
    City    string
    Country string
}

// Shallow copy - shares pointer to Address
func (p Person) ShallowCopy() Person {
    return Person{
        Name:    p.Name,
        Address: p.Address, // Same pointer!
    }
}

// Deep copy - creates new Address
func (p Person) DeepCopy() Person {
    newAddress := &Address{
        City:    p.Address.City,
        Country: p.Address.Country,
    }
    
    return Person{
        Name:    p.Name,
        Address: newAddress,
    }
}

// Deep copy using JSON (less efficient)
func (p Person) DeepCopyJSON() Person {
    data, _ := json.Marshal(p)
    var copy Person
    json.Unmarshal(data, &copy)
    return copy
}

func main() {
    original := Person{
        Name: "Alice",
        Address: &Address{
            City:    "New York",
            Country: "USA",
        },
    }
    
    // Shallow copy test
    shallow := original.ShallowCopy()
    shallow.Address.City = "San Francisco" // Affects original!
    fmt.Println("Original after shallow:", original.Address.City) // "San Francisco"
    
    // Deep copy test  
    original.Address.City = "New York" // Reset
    deep := original.DeepCopy()
    deep.Address.City = "Los Angeles" // Doesn't affect original
    fmt.Println("Original after deep:", original.Address.City) // "New York"
    fmt.Println("Deep copy:", deep.Address.City) // "Los Angeles"
}
```

---

## 7. Practice Problems & Analysis

### 7.1 Complexity Analysis Practice

#### Problem 1: Nested Loops Analysis
```javascript
function analyzeNestedLoops() {
    const examples = [
        {
            name: "Simple Nested",
            code: function(n) {
                let count = 0;
                for (let i = 0; i < n; i++) {        // n iterations
                    for (let j = 0; j < n; j++) {    // n iterations each
                        count++;                     // n × n = n² operations
                    }
                }
                return count; // O(n²)
            },
            complexity: "O(n²)"
        },
        
        {
            name: "Triangular Nested", 
            code: function(n) {
                let count = 0;
                for (let i = 0; i < n; i++) {        // n iterations
                    for (let j = 0; j <= i; j++) {   // 1+2+3+...+n = n(n+1)/2
                        count++;
                    }
                }
                return count; // O(n²) but with better constant
            },
            complexity: "O(n²)"
        },
        
        {
            name: "Logarithmic Inner Loop",
            code: function(n) {
                let count = 0;
                for (let i = 1; i <= n; i++) {       // n iterations
                    for (let j = 1; j <= n; j *= 2) { // log n iterations each
                        count++;
                    }
                }
                return count; // n × log n
            },
            complexity: "O(n log n)"
        }
    ];
    
    return examples;
}
```

#### Problem 2: Recursive Complexity Analysis  
```javascript
function recursiveComplexityExamples() {
    // T(n) = T(n-1) + O(1) → O(n)
    function linearRecursion(n) {
        if (n <= 0) return 0;
        return 1 + linearRecursion(n - 1);
    }
    
    // T(n) = 2T(n-1) + O(1) → O(2^n)
    function exponentialRecursion(n) {
        if (n <= 0) return 1;
        return exponentialRecursion(n - 1) + exponentialRecursion(n - 1);
    }
    
    // T(n) = T(n/2) + O(1) → O(log n)  
    function logarithmicRecursion(n) {
        if (n <= 1) return 1;
        return 1 + logarithmicRecursion(Math.floor(n / 2));
    }
    
    // T(n) = 2T(n/2) + O(n) → O(n log n)
    function divideAndConquer(arr) {
        if (arr.length <= 1) return arr;
        
        const mid = Math.floor(arr.length / 2);
        const left = divideAndConquer(arr.slice(0, mid));  // T(n/2)
        const right = divideAndConquer(arr.slice(mid));    // T(n/2)
        
        return merge(left, right); // O(n)
    }
    
    return {
        linear: linearRecursion,
        exponential: exponentialRecursion,
        logarithmic: logarithmicRecursion,
        divideConquer: divideAndConquer
    };
}
```

### 7.2 Amortized Analysis Practice

#### Problem: Dynamic Array with Shrinking
```javascript
class DynamicArrayWithShrinking {
    constructor() {
        this.data = new Array(1);
        this.size = 0;
        this.capacity = 1;
    }
    
    push(element) {
        if (this.size === this.capacity) {
            this.resize(this.capacity * 2);
        }
        this.data[this.size] = element;
        this.size++;
    }
    
    pop() {
        if (this.size === 0) return undefined;
        
        const element = this.data[this.size - 1];
        this.size--;
        
        // Shrink when 1/4 full to avoid thrashing
        if (this.size > 0 && this.size === Math.floor(this.capacity / 4)) {
            this.resize(Math.floor(this.capacity / 2));
        }
        
        return element;
    }
    
    resize(newCapacity) {
        const oldData = this.data;
        this.data = new Array(newCapacity);
        this.capacity = newCapacity;
        
        for (let i = 0; i < this.size; i++) {
            this.data[i] = oldData[i];
        }
    }
    
    // Amortized analysis using potential method
    static analyzePotential() {
        // Potential function: Φ(D) = |2×size - capacity|
        // This ensures we have enough "credits" for expensive operations
        
        const analysis = {
            pushNormal: "Amortized cost = 1 + 2 = 3",
            pushResize: "Amortized cost = size + 1 - size = 1", 
            popNormal: "Amortized cost = 1 - 2 = -1",
            popResize: "Amortized cost = size/2 + 1 - size/2 = 1"
        };
        
        return analysis;
    }
}
```

### 7.3 Master Theorem Practice

```javascript
function masterTheoremProblems() {
    const problems = [
        {
            recurrence: "T(n) = 4T(n/2) + n",
            analysis: {
                a: 4, b: 2, f: "n",
                logBA: 2, // log₂(4) = 2
                comparison: "n vs n²",
                case: 1,
                solution: "T(n) = Θ(n²)"
            }
        },
        
        {
            recurrence: "T(n) = 2T(n/2) + n log n", 
            analysis: {
                a: 2, b: 2, f: "n log n",
                logBA: 1, // log₂(2) = 1  
                comparison: "n log n vs n",
                case: 3,
                solution: "T(n) = Θ(n log n)"
            }
        },
        
        {
            recurrence: "T(n) = 3T(n/4) + n",
            analysis: {
                a: 3, b: 4, f: "n", 
                logBA: Math.log(3)/Math.log(4), // ≈ 0.79
                comparison: "n vs n^0.79",
                case: 3,
                solution: "T(n) = Θ(n)"
            }
        },
        
        {
            recurrence: "T(n) = T(n/3) + T(2n/3) + n",
            analysis: {
                note: "Cannot use Master Theorem directly - subproblems are unequal",
                method: "Use recursion tree or substitution method",
                solution: "T(n) = Θ(n log n)"
            }
        }
    ];
    
    return problems;
}

// Solve Master Theorem problems programmatically
function solveMasterTheorem(a, b, fType, n = 1000) {
    const logBA = Math.log(a) / Math.log(b);
    const threshold = Math.pow(n, logBA);
    
    let fValue;
    switch(fType) {
        case 'constant': fValue = 1; break;
        case 'linear': fValue = n; break;
        case 'quadratic': fValue = n * n; break;
        case 'nlogn': fValue = n * Math.log(n); break;
        case 'logarithmic': fValue = Math.log(n); break;
        default: fValue = n; // Default to linear
    }
    
    console.log(`Recurrence: T(n) = ${a}T(n/${b}) + f(n)`);
    console.log(`n^(log_${b}(${a})) = n^${logBA.toFixed(3)} ≈ ${threshold.toFixed(2)}`);
    console.log(`f(n) = ${fValue.toFixed(2)}`);
    
    if (fValue < threshold * 0.9) { // Case 1 with some epsilon
        console.log(`Case 1: f(n) = O(n^${logBA.toFixed(3)} - ε)`);
        console.log(`Solution: T(n) = Θ(n^${logBA.toFixed(3)})`);
    } else if (Math.abs(fValue - threshold) < threshold * 0.2) { // Case 2
        console.log(`Case 2: f(n) = Θ(n^${logBA.toFixed(3)})`);
        console.log(`Solution: T(n) = Θ(n^${logBA.toFixed(3)} log n)`);
    } else { // Case 3
        console.log(`Case 3: f(n) = Ω(n^${logBA.toFixed(3)} + ε)`);
        console.log(`Solution: T(n) = Θ(f(n))`);
    }
}
```

### 7.4 Memory Model Practice

#### Problem: Understanding Reference vs Value Semantics
```javascript
function memoryModelQuiz() {
    console.log("=== Memory Model Quiz ===\n");
    
    // Question 1: Primitive vs Object
    console.log("Question 1: What gets printed?");
    let a = 5;
    let b = a;
    b = 10;
    console.log(`a = ${a}, b = ${b}`); // Answer: a = 5, b = 10
    
    let obj1 = { x: 5 };
    let obj2 = obj1;
    obj2.x = 10;
    console.log(`obj1.x = ${obj1.x}, obj2.x = ${obj2.x}`); // Answer: obj1.x = 10, obj2.x = 10
    
    // Question 2: Function Parameters
    console.log("\nQuestion 2: Function parameter behavior");
    function modifyValues(num, obj, arr) {
        num = 100;        // Creates new local variable
        obj.prop = 200;   // Modifies original object
        arr = [4, 5, 6];  // Creates new local reference
        arr.push(7);      // Modifies local array, not original
    }
    
    let number = 50;
    let object = { prop: 10 };
    let array = [1, 2, 3];
    
    modifyValues(number, object, array);
    console.log(`number = ${number}`);     // 50 (unchanged)
    console.log(`object.prop = ${object.prop}`); // 200 (changed)
    console.log(`array = [${array}]`);     // [1,2,3] (unchanged)
    
    // Question 3: Array Methods
    console.log("\nQuestion 3: Array method behavior");
    let original = [1, 2, 3];
    let mapped = original.map(x => x * 2);    // Returns new array
    let pushed = original.push(4);           // Modifies original, returns length
    
    console.log(`original = [${original}]`);  // [1,2,3,4]
    console.log(`mapped = [${mapped}]`);      // [2,4,6]
    console.log(`pushed = ${pushed}`);        // 4 (length after push)
}

// Run the quiz
memoryModelQuiz();
```

#### Go Memory Model Practice
```go
package main

import "fmt"

func memoryModelPractice() {
    fmt.Println("=== Go Memory Model Practice ===")
    
    // Question 1: Value vs Pointer semantics
    type Point struct { X, Y int }
    
    p1 := Point{1, 2}
    p2 := p1        // Value copy
    p2.X = 10
    fmt.Printf("p1.X = %d, p2.X = %d\n", p1.X, p2.X) // 1, 10
    
    ptr1 := &Point{1, 2}
    ptr2 := ptr1    // Pointer copy
    ptr2.X = 10
    fmt.Printf("ptr1.X = %d, ptr2.X = %d\n", ptr1.X, ptr2.X) // 10, 10
    
    // Question 2: Slice behavior
    slice1 := []int{1, 2, 3}
    slice2 := slice1        // Copies slice header, shares underlying array
    slice2[0] = 10
    fmt.Printf("slice1 = %v, slice2 = %v\n", slice1, slice2) // [10 2 3], [10 2 3]
    
    slice3 := make([]int, len(slice1))
    copy(slice3, slice1)    // Deep copy of elements
    slice3[0] = 20
    fmt.Printf("slice1 = %v, slice3 = %v\n", slice1, slice3) // [10 2 3], [20 2 3]
    
    // Question 3: Function parameters
    modifyValue := func(x int) { x = 100 }           // No effect on original
    modifyPointer := func(x *int) { *x = 100 }       // Modifies original
    modifySlice := func(s []int) { s[0] = 100 }      // Modifies underlying array
    
    val := 42
    modifyValue(val)
    fmt.Printf("After modifyValue: %d\n", val)       // 42
    
    modifyPointer(&val)
    fmt.Printf("After modifyPointer: %d\n", val)     // 100
    
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Printf("After modifySlice: %v\n", slice)     // [100 2 3]
}
```

## Summary

This comprehensive guide covers the fundamental theoretical concepts essential for mastering Data Structures and Algorithms:

### Key Takeaways

1. **Asymptotic Analysis** provides mathematical tools to compare algorithm efficiency
   - Big-O: Upper bound (worst case)
   - Big-Omega: Lower bound (best case)  
   - Big-Theta: Tight bound (average case)

2. **Amortized Analysis** shows that expensive operations can have low average cost
   - Dynamic arrays: O(1) amortized insertion despite occasional O(n) resize
   - Hash tables: O(1) amortized operations with proper load factor management

3. **Mathematical Tools** enable precise algorithm analysis
   - Logarithms: Essential for tree heights and divide-and-conquer analysis
   - Series: Calculate total costs of nested operations
   - Growth rates: Compare algorithm scalability

4. **Recursion** requires careful analysis of base cases and recurrence relations
   - Master Theorem: Powerful tool for divide-and-conquer recurrences
   - Recursion trees: Visual method for complex recurrences
   - Optimization: Memoization and tail recursion improve performance

5. **Memory Models** affect both performance and correctness
   - Stack vs Heap: Different allocation strategies with trade-offs
   - Value vs Reference: Determines copying behavior and side effects
   - Shallow vs Deep Copy: Critical for data integrity in complex structures

Understanding these concepts provides the foundation for analyzing any algorithm's efficiency and choosing appropriate data structures for specific problems. Practice applying these tools to real algorithms to develop intuition for performance analysis and optimization.