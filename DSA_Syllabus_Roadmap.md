# Data Structures & Algorithms — Complete Syllabus & Roadmap
_Last updated: 2025-08-23_

> This syllabus is **language-agnostic** (use C++/Java/Python/Go/JS) and designed for **interview prep + foundations**.  
> Mark items as you complete them. Each section lists: **Core**, **Must‑know problems/patterns**, and **Optional/Advanced**.

---

## 0) How to Use This
- Aim for **steady practice** (5–10 problems per topic).
- For every topic: **(1) Learn → (2) Implement from scratch → (3) Solve problems → (4) Review mistakes**.
- Track complexity: always annotate **time/space**, and typical **trade‑offs**.
- Maintain small **code templates** (I/O, BFS/DFS, union‑find, heap, segtree, DP patterns, etc.).

---

## 1) Prerequisites & Core Analysis
**Core**
- Asymptotic analysis: **Big‑O / Θ / Ω**, best/avg/worst, **amortized analysis** (dynamic arrays, hash maps, union‑find).
- **Mathematical tools**: logarithms, series, growth rates.
- **Recursion**: base case, recurrence relations, **Master Theorem**, recursion tree.
- **Memory model basics**: stack vs heap, reference vs value semantics, mutability, shallow vs deep copy.

**Must‑know patterns**
- Derive/solve simple recurrences.
- Prove correctness at a high level (loop invariants, greedy choice property).

**Optional/Advanced**
- Potential method for amortized analysis.
- Probabilistic analysis (linearity of expectation).

---

## 2) Arrays & Two Pointers
**Core**
- Traversal, in‑place updates, prefix/suffix sums & **difference arrays**.
- **Searching**: linear, **binary search** (+ variants: first/last, lower/upper bound).
- **Sorting**: bubble/selection/insertion (learning only), **merge sort**, **quick sort** (Lomuto/Hoare), **counting sort** (when applicable).
- **Partitioning** (Dutch National Flag), **Quickselect** (K‑th order statistic).
- **Two pointers & Sliding window** (fixed/variable), **Monotonic deque** (sliding window max).

**Must‑know problems**
- Rotate/reverse array, remove duplicates, merge intervals, meeting rooms, 3‑sum/4‑sum.
- **Kadane’s algorithm**; product subarray; subarray sum = K (prefix sums + hashmap).
- Binary search on **answer/space** (min capacity to ship packages, Koko bananas, etc.).

**Optional/Advanced**
- **Sqrt decomposition** (range queries).
- **Mo’s algorithm** (offline query processing).

---

## 3) Strings
**Core**
- Basics: immutability, builders, substring, subsequence.
- **Pattern matching**: **KMP** (LPS table), **Z‑algorithm**, **Rabin–Karp** (rolling hash).
- **Frequency/Anagram** patterns, palindromes, expand‑around‑center.
- **Edit distance** (Levenshtein), LCS (ties into DP).

**Must‑know problems**
- Longest substring without repeats (hash set/map).
- Group anagrams, valid anagram, minimum window substring.
- Repeated substring, string compression, longest palindromic substring.

**Optional/Advanced**
- **Aho–Corasick** (multi‑pattern), **Suffix array/tree**, **Suffix automaton**, **Rolling hash** pitfalls (mod/overflow).
- **Trie** (below) for prefix problems.

---

## 4) Hashing (HashMap/Set) & Internals
**Core**
- Hash functions, collision resolution (**chaining**, open addressing: linear/quadratic/probing).
- Load factor, resizing, expected O(1) operations.
- Ordered vs unordered maps/sets.

**Must‑know problems**
- Two‑sum / three‑sum variants, first non‑repeating char, subarray with sum K, longest consecutive sequence.

**Optional/Advanced**
- **Rolling hash** for strings; double hashing.
- Probabilistic DS: **Bloom filter**, Count‑Min Sketch (awareness level).

---

## 5) Linked Lists
**Core**
- Singly, doubly, circular; slow/fast pointers; in‑place reversal; merge operations.

**Must‑know problems**
- Reverse list, reverse in K‑groups; detect & remove cycle (Floyd).
- Merge two sorted lists; add numbers (as lists); remove Nth from end.

**Optional/Advanced**
- Skip list (concepts), LFU/LRU (see “Design” section).

---

## 6) Stacks, Queues, Deques
**Core**
- Stack for parentheses, **Next Greater/Smaller Element** (monotonic stack).
- Queue & **circular queue**; **Deque** for sliding window maximum.
- **Min/Max stack** with auxiliary structure.

**Must‑know problems**
- Valid parentheses, evaluate RPN, daily temperatures, histogram largest rectangle, trapping rain water.

---

## 7) Trees & Binary Search Trees (BST)
**Core**
- Binary tree traversals: **pre/in/post‑order** (recursive & iterative), **level‑order** (BFS).
- Properties: height, depth, diameter, balance.
- **BST** operations & invariants; validate BST.

**Must‑know problems**
- LCA (parent pointers / **binary lifting** prep), diameter, path sum, serialize/deserialize, vertical/horizontal order traversal.

**Optional/Advanced**
- **Self‑balancing BSTs** (AVL, Red‑Black) at conceptual level.
- **Order‑statistics tree** (select/rank).

---

## 8) Heaps & Priority Queues
**Core**
- Min/max heap; heapify; push/pop; k‑way merge; top‑K patterns.

**Must‑know problems**
- K largest/smallest elements, kth element in stream, merge k sorted lists/arrays, sliding window median.

**Optional/Advanced**
- D‑ary heaps, indexed heaps; heap vs tree‑map trade‑offs.

---

## 9) Tries (Prefix Trees)
**Core**
- Insert/search/delete words; prefix count; word break synergy with DP.

**Must‑know problems**
- Autocomplete, longest word with all prefixes, replace words, maximum XOR pair (bitwise trie).

**Optional/Advanced**
- Compressed trie / radix tree; ternary search tree.

---

## 10) Disjoint Set Union (Union–Find)
**Core**
- **Union by rank/size** + **path compression**; connected components; cycle detection in undirected graphs.

**Must‑know problems**
- Redundant connection, number of provinces, Kruskal’s MST.

**Optional/Advanced**
- DSU with rollback (offline), DSU on tree (small‑to‑large).

---

## 11) Range Query Data Structures
**Core**
- **Prefix sums**; **Difference arrays** (range updates).
- **Sparse table** (idempotent queries like RMQ).
- **Fenwick Tree (BIT)**: point update/range sum; inverse prefix sum trick.
- **Segment Tree**: point & range queries; **lazy propagation**.

**Must‑know problems**
- Range sum/min/max; range add / range min; frequency of numbers; kth order statistic with Fenwick (advanced).

**Optional/Advanced**
- Segment tree beats (very advanced), persistent segment tree.
- Heavy‑light decomposition (with segtree) for tree path queries.

---

## 12) Graphs
**Core**
- Representations: adjacency list/matrix; directed vs undirected; weighted vs unweighted.
- Traversals: **BFS**, **DFS**; parent tracking; component count; bipartite check.
- **Topological sort** (Kahn/DFS); detect cycles (dir/undir).
- **Shortest paths**:
  - Unweighted: BFS
  - Weighted non‑negative: **Dijkstra** (binary heap)
  - With negatives: **Bellman–Ford**
  - All‑pairs: **Floyd–Warshall**
  - **0‑1 BFS**
- **Minimum Spanning Tree**: **Kruskal** (DSU), **Prim** (PQ).
- **Bridges & articulation points** (Tarjan); **SCCs** (Kosaraju/Tarjan).
- **Tree algorithms**: **LCA (binary lifting / Euler + RMQ)**, subtree sizes, tree DP.

**Must‑know problems**
- Number of islands/graphs of grids, course schedule, clone graph, network delay time, city with smallest neighbors at threshold.
- Count paths in DAG, critical connections, restore array from adjacent pairs.

**Optional/Advanced**
- **Max flow / Min cut**: Ford–Fulkerson/Edmonds–Karp, **Dinic**; applications (bipartite matching/assignment).
- Eulerian path/cycle; Hamiltonian path (NP‑hard awareness).
- Centroid decomposition; **HLD** (with segtree).

---

## 13) Dynamic Programming (DP)
**Core**
- **Bottom‑up vs Top‑down (memoization)**; state definition; transitions; base cases.
- Patterns:
  - 1D DP: climbing stairs/house robber/frog jumps
  - **Knapsack family**: 0/1, unbounded, bounded, subset sum/partition
  - **LIS** (O(n log n) patience sorting + reconstruction)
  - **Grid DP**: unique paths with obstacles, min path sum
  - Sequence DP: **LCS**, **Edit distance**, **Palindrome DP** (LPS/partition)
  - Interval DP: burst balloons, matrix chain multiplication
  - Bitmask DP: traveling salesman, assign tasks
  - DP on trees (rerooting, subtree DP)

**Must‑know problems**
- Coin change (min coins & #ways), decode ways, longest palindromic subsequence, buy and sell stock variants.

**Optional/Advanced**
- Optimizations: **divide & conquer DP**, **Knuth optimization**, **convex hull trick**, **monotonic queue optimization**.
- **Digit DP**, **SOS DP**, DP with bitsets, DP + binary search on answer.

---

## 14) Mathematics & Number Theory for DSA
**Core**
- **GCD**, Extended Euclid; **fast exponentiation**; **modular arithmetic** (mod inverse, Fermat’s little theorem).
- **Primes**: sieve of Eratosthenes, segmented sieve; prime factorization basics.
- **Combinatorics**: nCr, Pascal, factorials + inverse factorials under mod.

**Must‑know problems**
- LCM/GCD queries, modular inverse applications, counting with constraints.

**Optional/Advanced**
- **Chinese Remainder Theorem**, matrix exponentiation (linear recurrences), inclusion–exclusion principle.

---

## 15) Greedy & Divide‑and‑Conquer
**Core**
- Proving greedy choice & optimal substructure; exchange arguments.
- Classic greedy: activity selection, interval scheduling, Huffman coding, fractional knapsack.
- Divide‑and‑conquer: merge sort, quick sort, count inversions, closest pair of points.

**Must‑know problems**
- Min arrows to burst balloons, jump game(s), gas station, task scheduling with cooldown.

---

## 16) Backtracking & Search
**Core**
- Systematic search with pruning: permutations/combinations/subsets; N‑Queens; Sudoku; word search.
- Branch‑and‑bound awareness.

**Must‑know problems**
- Subset sum (exact), combination sum variants, phone keypad, restore IP addresses.

---

## 17) Computational Geometry (Basics)
**Core**
- Points/vectors, **orientation (cross product)**, segment intersection, polygon area (shoelace).
- **Convex hull** (Graham scan / Andrew monotone chain).
- Bounding boxes; line sweep (interval/segment problems).

**Optional/Advanced**
- Closest pair of points (D&C), rotating calipers (diameter, width), half‑plane intersection (awareness).

---

## 18) Design & Misc Patterns (Interview‑heavy)
**Core**
- **LRU cache** (hash map + doubly linked list).
- **Top‑K** (heap), streaming median (two heaps).
- Task scheduler, rate limiter (conceptual).
- Randomized algorithms: reservoir sampling; randomized quickselect.

**Optional/Advanced**
- Consistent hashing (awareness), bloom‑filter applications.

---

## 19) Testing, Templates & Debugging
- Unit‑test key DS (list, heap, union‑find, segtree).
- Reusable templates: BFS/DFS, Dijkstra, DSU, KMP, segtree/BIT, LIS (patience), binary lifting.
- Debug checklists: off‑by‑one, overflow, index bounds, visited resets, edge cases (empty, single, extremes).

---

## 20) Practice Strategy & Milestones
**Daily/Weekly Plan**
- Daily: 1–3 problems mixing **new topic + old review**.
- Weekly: 1 timed contest or mock interview; **retrospective of mistakes**.
- Maintain a **notes repo** with patterns & pitfalls, and a **“redo later”** list.

**Milestones**
1. Arrays/Strings/Hashing/LL/Stacks/Queues → basic 100 problems.
2. Trees/BST/Heaps/Tries/Union‑Find → intermediate 80 problems.
3. Graphs + DP core → 120 problems.
4. Advanced topics (range queries, optimizations, flows, geometry) → as needed.

---

## 21) Quick Checklist (tick ✓ as you master)
- [ ] Big‑O, amortized, recursion/Master theorem  
- [ ] Arrays: two pointers, sliding window, binary search (on answer), partition/quickselect  
- [ ] Strings: KMP, Z, RK, palindrome, edit distance  
- [ ] Hashing internals & rolling hash pitfalls  
- [ ] Linked lists: reverse (iter/rec), cycle, K‑group  
- [ ] Stacks/Queues/Deque, monotonic stack/queue  
- [ ] Trees/BST: traversals, LCA, diameter, path sums  
- [ ] Heaps/PQ: top‑K, stream median  
- [ ] Tries: prefix, word break, max XOR  
- [ ] Union‑Find with rank + path compression  
- [ ] Range queries: prefix/diff, sparse table, BIT, segtree + lazy  
- [ ] Graphs: BFS/DFS, topo sort, Dijkstra, Bellman‑Ford, Floyd‑Warshall, 0‑1 BFS  
- [ ] Graphs: bridges/AP, SCC, MST (Kruskal/Prim), tree DP, binary lifting  
- [ ] DP: knapsack family, LIS (n log n), LCS/edit, grid DP, interval DP, bitmask DP  
- [ ] DP optimizations: D&C, Knuth, CHT, monotonic queue  
- [ ] Math/NT: gcd/exgcd, modpow, mod inverse, sieve, nCr under mod  
- [ ] Backtracking: n‑queens, permutations/combinations  
- [ ] Greedy: activity selection, Huffman, interval scheduling  
- [ ] Geometry: orientation, hull, sweep line (basics)  
- [ ] Design: LRU, top‑K, reservoir sampling  
- [ ] Templates + tests + review loop

---

## 22) Bookmarks & References (pick a few)
- **CLRS** (Cormen et al.) — theory & proofs  
- **The Algorithm Design Manual** (Skiena) — heuristics & practice  
- **CP‑Algorithms (e-maxx)** — implementations & proofs  
- **Grokking Algorithms** — visual/introductory  
- Your preferred judge: LeetCode / Codeforces / AtCoder for practice

---

### Final Notes
- Treat this like a **syllabus**: go topic‑by‑topic, keep notes, and **revisit weak spots**.
- Mastery = knowing **when** to apply a pattern + writing it correctly **from memory** under time.
