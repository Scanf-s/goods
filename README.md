# Goods

Data structures implemented in Go, for study and review.
The goal is to understand how each structure works by building it from scratch.

## Data Structures

### 1. Linear Data Structures

#### List
- [x] ArrayList (Dynamic Array)
- [x] SinglyLinkedList
- [x] DoublyLinkedList
- [x] CircularLinkedList

#### Stack
- [x] ArrayStack
- [x] LinkedStack

#### Queue
- [x] ArrayQueue
- [x] LinkedQueue
- [x] CircularQueue
- [x] RingBuffer
- [x] Deque

**Benchmark: Ring buffer vs Linked-list Circular queue**

- Same workload (alternating `Offer`/`Poll`), measured with Go 1.24 `testing.B.Loop` on an AMD Ryzen 5 5600X (linux/amd64)

| Implementation | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `RingBuffer` (slice-based) | 4.50 | 0 | **0** |
| `CircularQueue` (linked list) | 24.15 | 24 | 1 |

- The slice-based ring buffer runs on contiguous memory with no per-element heap allocation.
- So it stays cache-friendly and puts zero pressure on the GC.
- !!!Roughly **6x lower latency** than the linked-list queue under the same workload.!!!

```text
goos: linux
goarch: amd64
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkRingBuffer-12       271154154    4.506 ns/op    0 B/op    0 allocs/op
BenchmarkCircularQueue-12     46714058   24.15 ns/op    24 B/op    1 allocs/op
```

### 2. Tree Data Structures
- [x] BinaryTree
- [x] BinarySearchTree
- [x] Trie
- [ ] Heap

### 3. Hash-based Structures
- [ ] HashMap
- [ ] HashSet

### 4. Graph
- [ ] AdjacencyMatrix
- [ ] AdjacencyList

### 5. Cache
- [x] LRU Cache
