# GOODS - Go Data Structures & Algorithms

## Project Overview
This project is a study-focused implementation of common data structures and algorithms in Go. The goal is to understand the underlying principles, time/space complexity, and practical applications of each implementation.

## Coding Conventions

### General
- Use Go generics (`[T any]` or `[T comparable]`) for type-safe implementations
- Each data structure gets its own package under a category folder
- Include unit tests for every implementation (`*_test.go`)
- Document time and space complexity in comments

### File Structure
```
goods/
├── list/                    # Linear data structures
│   ├── list.go             # Common interfaces
│   ├── arraylist/
│   └── linkedlist/
├── stack/
├── queue/
├── tree/
├── graph/
├── heap/
├── hash/
└── algorithms/
    ├── sort/
    ├── search/
    └── ...
```

### Naming
- Package names: lowercase, single word (e.g., `arraylist`, `binarytree`)
- Struct names: PascalCase (e.g., `ArrayList`, `BinaryTree`)
- Interface names: describe behavior (e.g., `Sortable`, `Comparable`)

---

## Data Structures Roadmap

### 1. Linear Data Structures

#### List
- [x] ArrayList (Dynamic Array)
- [x] SinglyLinkedList
- [x] DoublyLinkedList
- [x] CircularLinkedList

#### Stack
- [x] ArrayStack
- [ ] LinkedStack

#### Queue
- [ ] ArrayQueue
- [ ] LinkedQueue
- [ ] CircularQueue
- [ ] Deque (Double-ended Queue)
- [ ] PriorityQueue

### 2. Tree Data Structures

#### Binary Trees
- [ ] BinaryTree (basic)
- [ ] BinarySearchTree (BST)
- [ ] AVLTree (self-balancing)
- [ ] RedBlackTree

#### Other Trees
- [ ] Trie (Prefix Tree)
- [ ] Heap (MinHeap, MaxHeap)
- [ ] B-Tree
- [ ] Segment Tree

### 3. Hash-based Structures
- [ ] HashMap
- [ ] HashSet
- [ ] BloomFilter

### 4. Graph
- [ ] AdjacencyMatrix
- [ ] AdjacencyList
- [ ] WeightedGraph

---

## Algorithms Roadmap

### 1. Sorting Algorithms
- [ ] BubbleSort - O(n²)
- [ ] SelectionSort - O(n²)
- [ ] InsertionSort - O(n²)
- [ ] MergeSort - O(n log n)
- [ ] QuickSort - O(n log n) average
- [ ] HeapSort - O(n log n)
- [ ] CountingSort - O(n + k)
- [ ] RadixSort - O(nk)

### 2. Searching Algorithms
- [ ] LinearSearch - O(n)
- [ ] BinarySearch - O(log n)
- [ ] JumpSearch - O(√n)
- [ ] InterpolationSearch - O(log log n)

### 3. Graph Algorithms
- [ ] BFS (Breadth-First Search)
- [ ] DFS (Depth-First Search)
- [ ] Dijkstra's Algorithm
- [ ] Bellman-Ford Algorithm
- [ ] Floyd-Warshall Algorithm
- [ ] Kruskal's Algorithm (MST)
- [ ] Prim's Algorithm (MST)
- [ ] Topological Sort

### 4. Dynamic Programming
- [ ] Fibonacci (memoization & tabulation)
- [ ] Longest Common Subsequence
- [ ] Knapsack Problem
- [ ] Coin Change
- [ ] Edit Distance

### 5. Other Algorithms
- [ ] Two Pointers
- [ ] Sliding Window
- [ ] Divide and Conquer
- [ ] Greedy Algorithms
- [ ] Backtracking

---

## Implementation Guidelines

### For Each Data Structure
1. Define the struct with proper fields
2. Implement core operations (insert, delete, search, etc.)
3. Add utility methods (size, isEmpty, clear, etc.)
4. Write comprehensive tests
5. Document complexity in comments

### Testing Convention
```go
func TestDataStructure_Operation(t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

### Complexity Documentation
```go
// Insert adds an element to the list.
// Time Complexity: O(1) average, O(n) worst case
// Space Complexity: O(1)
func (l *ArrayList[T]) Insert(element T) { ... }
```

---

## Current Progress
- [x] Project initialization
- [x] List interface definition
- [x] ArrayList implementation (with 15 unit tests)
- [x] SinglyLinkedList implementation
- [x] DoublyLinkedList implementation
- [x] CircularLinkedList implementation
- [x] Stack interface definition
- [ ] ArrayStack implementation (Need to add unit tests)
- [ ]