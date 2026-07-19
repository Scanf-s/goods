package lru_cache

import "testing"

func TestEviction(t *testing.T) {
	c, _ := NewLRUCache[int, int](2)
	c.Put(1, 1)
	if len(c.Storage) != 1 {
		t.Fatalf("storage has %d entries, want 1", len(c.Storage))
	}

	c.Put(2, 2)
	if len(c.Storage) != 2 {
		t.Fatalf("storage has %d entries, want 2", len(c.Storage))
	}

	c.Put(3, 3)
	if _, ok := c.Get(1); ok {
		t.Fatal("key 1 should have been evicted")
	}
	if _, ok := c.Get(2); !ok {
		t.Fatal("key 2 should exist in cache")
	}
	if len(c.Storage) != 2 {
		t.Fatalf("storage has %d entries, want 2", len(c.Storage))
	}

	c.Put(4, 4)
	if len(c.Storage) != 2 {
		t.Fatalf("storage has %d entries, want 2", len(c.Storage))
	}
	if _, ok := c.Get(3); ok {
		t.Fatal("key 3 should have been evicted")
	}
}

func TestLeetCodeSequence(t *testing.T) {
	c, _ := NewLRUCache[int, int](2)
	c.Put(1, 1)
	c.Put(2, 2)
	if v, ok := c.Get(1); !ok || v != 1 {
		t.Fatalf("Get(1) = %v,%v; want 1,true", v, ok)
	}
	c.Put(3, 3) // evicts 2
	if _, ok := c.Get(2); ok {
		t.Fatal("key 2 should have been evicted")
	}
	c.Put(4, 4) // evicts 1
	if _, ok := c.Get(1); ok {
		t.Fatal("key 1 should have been evicted")
	}
	if v, _ := c.Get(3); v != 3 {
		t.Fatalf("Get(3) = %v; want 3", v)
	}
	if v, _ := c.Get(4); v != 4 {
		t.Fatalf("Get(4) = %v; want 4", v)
	}
}

func TestUpdateExistingKeyRefreshesRecency(t *testing.T) {
	c, _ := NewLRUCache[int, int](2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(1, 10) // 1 becomes MRU
	c.Put(3, 3)  // must evict 2, not 1
	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("Get(1) = %v,%v; want 10,true", v, ok)
	}
	if _, ok := c.Get(2); ok {
		t.Fatal("key 2 should have been evicted")
	}
}

func TestCapacityOne(t *testing.T) {
	c, _ := NewLRUCache[int, int](1)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	if _, ok := c.Get(2); ok {
		t.Fatal("key 2 should have been evicted")
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Fatalf("Get(3) = %v,%v; want 3,true", v, ok)
	}
	if len(c.Storage) != 1 || c.Size != 1 {
		t.Fatalf("len=%d size=%d; want 1,1", len(c.Storage), c.Size)
	}
}

func TestStorageBounded(t *testing.T) {
	c, _ := NewLRUCache[int, int](3)
	for i := 0; i < 1000; i++ {
		c.Put(i, i)
	}
	if len(c.Storage) != 3 {
		t.Fatalf("storage grew to %d entries; want 3", len(c.Storage))
	}
	if c.Size != 3 {
		t.Fatalf("Size = %d; want 3", c.Size)
	}
}

func TestZeroValueDistinguishable(t *testing.T) {
	c, _ := NewLRUCache[string, int](2)
	c.Put("k", 0)
	if v, ok := c.Get("k"); !ok || v != 0 {
		t.Fatalf("Get(k) = %v,%v; want 0,true", v, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Fatal("missing key reported as hit")
	}
}

func TestNegativeCapacityRejected(t *testing.T) {
	if _, err := NewLRUCache[int, int](-1); err == nil {
		t.Fatal("NewLRUCache(-1) should return an error")
	}
}

func TestShrinkCapacityEvicts(t *testing.T) {
	c, _ := NewLRUCache[int, int](5)
	for i := 1; i <= 5; i++ {
		c.Put(i, i)
	}
	err := c.UpdateCapacity(-1)
	if err == nil {
		t.Fatal("negative capacity should be rejected.")
	}
	err = c.UpdateCapacity(2)
	if err != nil {
		t.Fatalf("failed to update capacity: %v", err)
	}
	c.Put(6, 6)
	if len(c.Storage) > 2 {
		t.Fatalf("after shrink to 2, storage has %d entries", len(c.Storage))
	}
}
