// Package trie implements a prefix tree (trie) that stores a set of strings.
//
// The trie walks a word rune by rune, so it works with any UTF-8 string and
// not only with ASCII letters. Every node keeps its children in a map that is
// keyed by a single rune, and a node is marked when a complete word ends there.
//
// The contract below is what the tests in this package expect:
//
//   - An empty word is never a valid argument. Insert and Delete reject it with
//     ErrEmptyWord, and Search always reports false for it.
//   - An empty prefix is a prefix of every word. StartsWith("") is therefore
//     true for a trie that holds at least one word, and WordsWithPrefix("")
//     returns every word in the trie.
//   - Words are compared by their exact runes, so "Apple" and "apple" are two
//     different words.
//   - Delete removes the word and also removes every node that no longer
//     belongs to another word.
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
	// TODO: implement
	return nil
}

// Search reports whether the exact word is stored in the trie.
// A prefix of a stored word is not a match unless it was inserted itself.
func (t *Trie) Search(word string) bool {
	// TODO: implement
	return false
}

// StartsWith reports whether at least one stored word begins with the prefix.
// A stored word counts as a prefix of itself.
func (t *Trie) StartsWith(prefix string) bool {
	// TODO: implement
	return false
}

// Delete removes the word from the trie and prunes every node that is no
// longer part of another word.
// It returns ErrWordNotFound when the word is not stored, ErrEmptyWord for an
// empty word and ErrNilTrie for a nil trie.
func (t *Trie) Delete(word string) error {
	// TODO: implement
	return nil
}

// WordsWithPrefix returns every stored word that begins with the prefix.
// The order of the returned words is not defined, so a caller that needs a
// stable order has to sort the result.
// It returns an empty result when no word matches.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	// TODO: implement
	return nil
}

// Size returns the number of distinct words in the trie.
func (t *Trie) Size() int {
	// TODO: implement
	return 0
}

// IsEmpty reports whether the trie holds no word at all.
func (t *Trie) IsEmpty() bool {
	// TODO: implement
	return true
}

// Clear removes every word from the trie.
func (t *Trie) Clear() {
	// TODO: implement
}
