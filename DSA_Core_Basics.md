# 📘 Core DSA Concepts (Beginner-Friendly Guide)

This guide explains the **core foundation of DSA** in simple language
with examples.\
Read it carefully --- it will make your learning journey much easier.

------------------------------------------------------------------------

## 1. Asymptotic Analysis

When we write algorithms, we want to know **how fast or slow** they are
when the input grows very large.

Instead of measuring actual time (depends on your computer), we use
**asymptotic analysis**.

### Big-O (O)

-   **Worst-case performance** → upper bound.
-   Example: Searching in an **unsorted array** of `n` items → O(n) (you
    may check every element).

### Big-Theta (Θ)

-   **Tight bound** → both lower and upper bound are the same.
-   Example: Simple loop from `1 to n` → Θ(n).

### Big-Omega (Ω)

-   **Best-case performance** → lower bound.
-   Example: Linear search → Ω(1) (if the item is at the start).

------------------------------------------------------------------------

### Best, Average, Worst Case

-   **Best case**: fastest scenario (item found immediately).
-   **Average case**: average expected runtime over many inputs.
-   **Worst case**: slowest scenario (item found at the very end).

We usually care most about **Worst case**.

------------------------------------------------------------------------

### Amortized Analysis

Sometimes an operation seems slow, but **averaged over many
operations**, it is fast.

Examples: 1. **Dynamic arrays (like Go slices, JS arrays):** - When
full, they **double size** → that one operation is costly. - But most
operations are O(1). Average (amortized) = O(1).

2.  **Hash maps:**
    -   Most lookups = O(1).
    -   But if two keys land in the same bucket (collision), it may take
        O(n).
    -   Still, **average is O(1)**.
3.  **Union-Find (Disjoint Set Union):**
    -   With path compression, multiple operations average out to
        **almost O(1)**.

------------------------------------------------------------------------

## 2. Mathematical Tools

### Logarithms (log)

-   `log_b(n)` = "How many times must you divide `n` by `b` to reach
    1?"\

-   Example: `log2(8) = 3` because 8 → 4 → 2 → 1 (3 steps).

-   In CS, **log means log base 2** (because of binary computers).

### Series

-   Repeated sums.
-   Example: `1 + 2 + 3 + … + n = n(n+1)/2` ≈ O(n²).

### Growth Rates

We compare functions to know which grows faster as input increases: -
Constant: O(1)\
- Logarithmic: O(log n)\
- Linear: O(n)\
- Linearithmic: O(n log n)\
- Quadratic: O(n²)\
- Exponential: O(2\^n)

Rule: **Lower is better** for performance.

------------------------------------------------------------------------

## 3. Recursion

Recursion = A function that **calls itself**.

### Base Case

-   A condition to **stop recursion**, otherwise it will go infinite.

Example in JS:

``` js
function factorial(n) {
  if (n === 0) return 1; // base case
  return n * factorial(n-1);
}
```

### Recurrence Relations

Equation describing recursive runtime. - Example: Merge Sort → T(n) =
2T(n/2) + O(n).

### Master Theorem

A shortcut to solve such recurrences.

For T(n) = aT(n/b) + f(n):\
- If f(n) smaller → O(n\^(log_b a)).\
- If equal → O(n\^(log_b a) \* log n).\
- If larger → O(f(n)).

Example: Merge Sort → a=2, b=2, f(n)=O(n) → O(n log n).

### Recursion Tree

A diagram breaking recursive calls into levels. Helps visualize how many
calls happen.

------------------------------------------------------------------------

## 4. Memory Model Basics

When you write code, memory is divided mainly into **stack** and
**heap**.

### Stack (fast)

-   Stores function calls, local variables.
-   Last In, First Out (like a plate stack).\
-   Example: recursion uses stack.

### Heap (big, slower)

-   Stores dynamic data (arrays, objects, slices).\
-   You control allocation.

------------------------------------------------------------------------

### Reference vs Value Semantics

-   **Value semantics**: Copy the actual data (like numbers).
    -   Changing one copy does NOT affect the other.\
-   **Reference semantics**: Copy the reference (address).
    -   Changing one affects the other.

Go Example:

``` go
a := []int{1,2,3}
b := a
b[0] = 99
fmt.Println(a) // [99 2 3] → because slice is reference type
```

------------------------------------------------------------------------

### Mutability

-   **Mutable** = can be changed (arrays, objects).\
-   **Immutable** = cannot be changed (strings in JS, Go).

Example in JS:

``` js
let s = "hello";
s[0] = "H"; // ❌ doesn’t work
s = "Hello"; // ✅ must create a new string
```

------------------------------------------------------------------------

### Shallow vs Deep Copy

-   **Shallow copy**: copies references, not actual data. Changes affect
    both.\
-   **Deep copy**: copies entire structure independently.

JS Example:

``` js
let arr1 = [[1,2], [3,4]];
let shallow = arr1.slice(); 
shallow[0][0] = 99;
console.log(arr1); // [[99,2],[3,4]] → changed!

let deep = JSON.parse(JSON.stringify(arr1));
deep[0][0] = 77;
console.log(arr1); // [[99,2],[3,4]] → original safe
```

------------------------------------------------------------------------

# ✅ Summary

-   **Asymptotic analysis** = compare algorithm efficiency (Big-O, Θ, Ω,
    amortized).\
-   **Math tools** = logs, series, growth rates help analyze
    complexity.\
-   **Recursion** = break problems into smaller ones (base case,
    recurrence, master theorem).\
-   **Memory model** = stack vs heap, value vs reference, mutability,
    shallow vs deep copy.

These are the foundations of DSA. Master these before moving on to
advanced data structures.
