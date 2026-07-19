package lru_cache

import (
	"fmt"

	"github.com/Scanf-s/goods/cache"
)

type LRUCache[K comparable, V any] struct {

	// Capacity represents the maximum capacity of LRU cache
	Capacity int

	// Storage of the LRU cache (HashMap)
	Storage map[K]*cache.Node[K, V]

	// Head (sentinel node)
	Head *cache.Node[K, V]

	// Tail (sentinel node)
	Tail *cache.Node[K, V]

	// Size represents current size of storage
	Size int
}

func NewLRUCache[K comparable, V any](capacity int) (*LRUCache[K, V], error) {
	var defaultKey K
	var defaultValue V
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be a positive integer")
	}
	head := &cache.Node[K, V]{Key: defaultKey, Data: defaultValue}
	tail := &cache.Node[K, V]{Key: defaultKey, Data: defaultValue}
	head.Next = tail
	head.Prev = nil
	tail.Prev = head
	tail.Next = nil

	return &LRUCache[K, V]{
		Capacity: capacity,
		Storage:  make(map[K]*cache.Node[K, V], capacity),
		Head:     head,
		Tail:     tail,
		Size:     0,
	}, nil
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	var defaultValue V
	if node := c.Storage[key]; node != nil {
		detachNode(node)
		c.lruNodeUpdate(node)
		return node.Data, true
	}
	// Cache miss
	return defaultValue, false
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if node := c.Storage[key]; node != nil {
		node.Data = value
		detachNode(node)
		c.lruNodeUpdate(node)
		return
	}

	if c.Capacity == c.Size {
		oldestNode := c.Head.Next
		detachNode(oldestNode)
		delete(c.Storage, oldestNode.Key)
		c.Size--
	}

	newNode := &cache.Node[K, V]{Key: key, Data: value}
	c.Storage[key] = newNode
	c.lruNodeUpdate(newNode)
	c.Size++
}

func (c *LRUCache[K, V]) UpdateCapacity(capacity int) error {
	if capacity <= 0 {
		return fmt.Errorf("capacity must be positive, got %d", capacity)
	}
	c.Capacity = capacity
	for c.Size > c.Capacity {
		oldestNode := c.Head.Next
		detachNode(oldestNode)
		delete(c.Storage, oldestNode.Key)
		c.Size--
	}
	return nil
}

func detachNode[K comparable, V any](node *cache.Node[K, V]) {
	prevNode := node.Prev
	nextNode := node.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode
}

func (c *LRUCache[K, V]) lruNodeUpdate(node *cache.Node[K, V]) {
	latestNode := c.Tail.Prev
	node.Next = c.Tail
	c.Tail.Prev = node
	latestNode.Next = node
	node.Prev = latestNode
}
