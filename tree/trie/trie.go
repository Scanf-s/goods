package trie

import "errors"

var (
	// ErrEmptyWord is returned when a caller passes an empty word.
	ErrEmptyWord = errors.New("trie: word must not be empty")

	// ErrWordNotFound is returned when Delete cannot find the given word.
	ErrWordNotFound = errors.New("trie: word not found")

	// ErrNilTrie is returned when a method runs on a nil trie pointer.
	ErrNilTrie = errors.New("trie: trie is not initialized")
)

type (
	// Node represents a single node in the trie.
	Node struct {
		// Children maps one rune to the node that follows it.
		Children map[rune]*Node

		// IsEndOfWord is true when a complete word ends at this node.
		IsEndOfWord bool
	}

	// Trie is a prefix tree over a set of strings.
	Trie struct {
		// Root is the node above the first rune of every word.
		// The root itself never stores a rune.
		Root *Node

		// size is the number of distinct words in the trie.
		size int
	}
)

// NewNode creates an empty node with an initialized child map.
func NewNode() *Node {
	return &Node{
		Children:    make(map[rune]*Node),
		IsEndOfWord: false,
	}
}

// NewTrie creates an empty trie with a root node.
func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
		size: 0,
	}
}

// Insert adds a word to the trie.
// Inserting the same word twice does not change the size of the trie.
// It returns ErrEmptyWord for an empty word and ErrNilTrie for a nil trie.
func (t *Trie) Insert(word string) error {
	if t == nil {
		return ErrNilTrie
	}
	if word == "" {
		return ErrEmptyWord
	}
	if t.Root == nil {
		t.Root = NewNode()
	}

	node := t.Root
	for _, r := range word {
		if node.Children == nil {
			node.Children = make(map[rune]*Node)
		}
		if node.Children[r] == nil {
			node.Children[r] = NewNode()
		}
		node = node.Children[r]
	}
	if !node.IsEndOfWord {
		node.IsEndOfWord = true
		t.size++
	}
	return nil
}

// Search reports whether the exact word is stored in the trie.
// A prefix of a stored word is not a match unless it was inserted itself.
func (t *Trie) Search(word string) bool {
	if word == "" {
		return false
	}
	node := t.findNode(word)
	return node != nil && node.IsEndOfWord
}

// StartsWith reports whether at least one stored word begins with the prefix.
// A stored word counts as a prefix of itself.
func (t *Trie) StartsWith(prefix string) bool {
	if t == nil || t.size == 0 {
		return false
	}
	return t.findNode(prefix) != nil
}

// Delete removes the word from the trie and prunes every node that is no
// longer part of another word.
// It returns ErrWordNotFound when the word is not stored, ErrEmptyWord for an
// empty word and ErrNilTrie for a nil trie.
func (t *Trie) Delete(word string) error {
	if t == nil {
		return ErrNilTrie
	}
	if word == "" {
		return ErrEmptyWord
	}
	if t.Root == nil {
		return ErrWordNotFound
	}

	// Keep the path so unused nodes can be removed from the bottom up.
	runes := []rune(word)
	path := make([]*Node, 1, len(runes)+1)
	path[0] = t.Root
	for _, r := range runes {
		child := path[len(path)-1].Children[r]
		if child == nil {
			return ErrWordNotFound
		}
		path = append(path, child)
	}

	last := path[len(path)-1]
	if !last.IsEndOfWord {
		return ErrWordNotFound
	}
	last.IsEndOfWord = false
	t.size--

	for i := len(runes) - 1; i >= 0; i-- {
		child := path[i+1]
		if child.IsEndOfWord || len(child.Children) != 0 {
			break
		}
		delete(path[i].Children, runes[i])
	}
	return nil
}

// WordsWithPrefix returns every stored word that begins with the prefix.
// The order of the returned words is not defined, so a caller that needs a
// stable order has to sort the result.
// It returns an empty result when no word matches.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	start := t.findNode(prefix)
	if start == nil || t.size == 0 {
		return nil
	}

	var words []string
	var visit func(*Node, string)
	visit = func(node *Node, word string) {
		if node.IsEndOfWord {
			words = append(words, word)
		}
		for r, child := range node.Children {
			visit(child, word+string(r))
		}
	}
	visit(start, prefix)
	return words
}

// Size returns the number of distinct words in the trie.
func (t *Trie) Size() int {
	if t == nil {
		return 0
	}
	return t.size
}

// IsEmpty reports whether the trie holds no word at all.
func (t *Trie) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes every word from the trie.
func (t *Trie) Clear() {
	if t == nil {
		return
	}
	t.Root = NewNode()
	t.size = 0
}

// findNode follows a path of runes from the root. A missing path returns nil.
func (t *Trie) findNode(path string) *Node {
	if t == nil || t.Root == nil {
		return nil
	}
	node := t.Root
	for _, r := range path {
		node = node.Children[r]
		if node == nil {
			return nil
		}
	}
	return node
}
