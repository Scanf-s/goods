package trie_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/Scanf-s/goods/tree/trie"
)

// sampleWords is the word set used by most tests. It contains words that share
// a prefix, a word that is a prefix of another word, and a second branch under
// the root.
var sampleWords = []string{"app", "apple", "apply", "apt", "bad", "bat", "batch"}

// newTrieWith builds a trie that holds the given words.
func newTrieWith(t *testing.T, words ...string) *trie.Trie {
	t.Helper()
	tr := trie.NewTrie()
	for _, w := range words {
		if err := tr.Insert(w); err != nil {
			t.Fatalf("Insert(%q) returned unexpected error: %v", w, err)
		}
	}
	return tr
}

// checkWords compares two word lists while ignoring their order, because
// WordsWithPrefix does not promise a stable order.
func checkWords(t *testing.T, got, want []string, call string) {
	t.Helper()
	gotSorted := slices.Clone(got)
	wantSorted := slices.Clone(want)
	slices.Sort(gotSorted)
	slices.Sort(wantSorted)
	if !slices.Equal(gotSorted, wantSorted) {
		t.Errorf("%s = %v, want %v (order ignored)", call, gotSorted, wantSorted)
	}
}

// checkSearch asserts the result of Search for a single word.
func checkSearch(t *testing.T, tr *trie.Trie, word string, want bool) {
	t.Helper()
	if got := tr.Search(word); got != want {
		t.Errorf("Search(%q) = %t, want %t", word, got, want)
	}
}

// checkSize asserts the current number of stored words.
func checkSize(t *testing.T, tr *trie.Trie, want int) {
	t.Helper()
	if got := tr.Size(); got != want {
		t.Errorf("Size() = %d, want %d", got, want)
	}
	if got, wantEmpty := tr.IsEmpty(), want == 0; got != wantEmpty {
		t.Errorf("IsEmpty() = %t, want %t", got, wantEmpty)
	}
}

func TestNewTrieIsEmpty(t *testing.T) {
	tr := trie.NewTrie()

	if tr.Root == nil {
		t.Fatal("NewTrie().Root is nil, want an initialized root node")
	}
	if len(tr.Root.Children) != 0 {
		t.Errorf("new root has %d children, want 0", len(tr.Root.Children))
	}
	if tr.Root.IsEndOfWord {
		t.Error("new root should not be marked as the end of a word")
	}
	checkSize(t, tr, 0)

	if tr.Search("apple") {
		t.Error("Search on an empty trie should be false")
	}
	if tr.StartsWith("a") {
		t.Error("StartsWith on an empty trie should be false")
	}
	if words := tr.WordsWithPrefix("a"); len(words) != 0 {
		t.Errorf("WordsWithPrefix on an empty trie = %v, want no word", words)
	}
}

func TestInsertAndSearchSingleWord(t *testing.T) {
	tr := newTrieWith(t, "apple")

	checkSize(t, tr, 1)
	checkSearch(t, tr, "apple", true)
	checkSearch(t, tr, "banana", false)

	// The root must hold exactly one branch, and that branch starts with 'a'.
	if len(tr.Root.Children) != 1 {
		t.Errorf("root has %d children, want 1", len(tr.Root.Children))
	}
	if _, ok := tr.Root.Children['a']; !ok {
		t.Error("root should have a child for rune 'a'")
	}
}

// A prefix of a stored word is not a stored word by itself.
func TestSearchDistinguishesWordFromPrefix(t *testing.T) {
	tr := newTrieWith(t, "apple")

	for _, prefix := range []string{"a", "ap", "app", "appl"} {
		checkSearch(t, tr, prefix, false)
		if !tr.StartsWith(prefix) {
			t.Errorf("StartsWith(%q) = false, want true", prefix)
		}
	}
	checkSearch(t, tr, "apple", true)
	if !tr.StartsWith("apple") {
		t.Error("StartsWith(\"apple\") = false, want true because a word is a prefix of itself")
	}
	if tr.StartsWith("apples") {
		t.Error("StartsWith(\"apples\") = true, want false because the word is longer than any stored word")
	}
}

func TestStartsWith(t *testing.T) {
	tr := newTrieWith(t, sampleWords...)

	tests := []struct {
		prefix string
		want   bool
	}{
		{prefix: "a", want: true},
		{prefix: "ap", want: true},
		{prefix: "app", want: true},
		{prefix: "appl", want: true},
		{prefix: "apple", want: true},
		{prefix: "apples", want: false},
		{prefix: "b", want: true},
		{prefix: "bat", want: true},
		{prefix: "batches", want: false},
		{prefix: "c", want: false},
		{prefix: "ba d", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.prefix, func(t *testing.T) {
			if got := tr.StartsWith(tc.prefix); got != tc.want {
				t.Errorf("StartsWith(%q) = %t, want %t", tc.prefix, got, tc.want)
			}
		})
	}
}

func TestInsertDuplicateKeepsSizeStable(t *testing.T) {
	tr := newTrieWith(t, "apple", "apple", "apple")

	checkSize(t, tr, 1)
	checkSearch(t, tr, "apple", true)
}

func TestInsertRejectsEmptyWord(t *testing.T) {
	tr := newTrieWith(t, "apple")

	err := tr.Insert("")
	if !errors.Is(err, trie.ErrEmptyWord) {
		t.Errorf("Insert(\"\") returned %v, want ErrEmptyWord", err)
	}
	checkSize(t, tr, 1)
	if tr.Root.IsEndOfWord {
		t.Error("a rejected empty word must not mark the root as the end of a word")
	}
}

func TestSearchEmptyWordIsAlwaysFalse(t *testing.T) {
	empty := trie.NewTrie()
	if empty.Search("") {
		t.Error("Search(\"\") on an empty trie = true, want false")
	}

	tr := newTrieWith(t, sampleWords...)
	if tr.Search("") {
		t.Error("Search(\"\") = true, want false because an empty word is never stored")
	}
}

// The empty prefix is a prefix of every word, so it only depends on whether the
// trie holds any word at all.
func TestStartsWithEmptyPrefix(t *testing.T) {
	empty := trie.NewTrie()
	if empty.StartsWith("") {
		t.Error("StartsWith(\"\") on an empty trie = true, want false")
	}

	tr := newTrieWith(t, "apple")
	if !tr.StartsWith("") {
		t.Error("StartsWith(\"\") on a non empty trie = false, want true")
	}
}

func TestWordsWithPrefix(t *testing.T) {
	tr := newTrieWith(t, sampleWords...)

	tests := []struct {
		name   string
		prefix string
		want   []string
	}{
		{
			name:   "shared branch",
			prefix: "app",
			want:   []string{"app", "apple", "apply"},
		},
		{
			name:   "single match",
			prefix: "apt",
			want:   []string{"apt"},
		},
		{
			name:   "full word that has no longer word",
			prefix: "apple",
			want:   []string{"apple"},
		},
		{
			name:   "second branch",
			prefix: "ba",
			want:   []string{"bad", "bat", "batch"},
		},
		{
			name:   "no match",
			prefix: "cat",
			want:   []string{},
		},
		{
			name:   "prefix longer than every word",
			prefix: "applesauce",
			want:   []string{},
		},
		{
			name:   "empty prefix returns every word",
			prefix: "",
			want:   sampleWords,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tr.WordsWithPrefix(tc.prefix)
			checkWords(t, got, tc.want, fmt.Sprintf("WordsWithPrefix(%q)", tc.prefix))
		})
	}
}

func TestWordsWithPrefixOnEmptyTrie(t *testing.T) {
	tr := trie.NewTrie()

	if got := tr.WordsWithPrefix(""); len(got) != 0 {
		t.Errorf("WordsWithPrefix(\"\") = %v, want no word", got)
	}
	if got := tr.WordsWithPrefix("a"); len(got) != 0 {
		t.Errorf("WordsWithPrefix(\"a\") = %v, want no word", got)
	}
}

func TestDeleteRemovesOnlyTheTargetWord(t *testing.T) {
	tr := newTrieWith(t, sampleWords...)

	if err := tr.Delete("apple"); err != nil {
		t.Fatalf("Delete(\"apple\") returned unexpected error: %v", err)
	}

	checkSearch(t, tr, "apple", false)
	checkSize(t, tr, len(sampleWords)-1)
	for _, w := range []string{"app", "apply", "apt", "bad", "bat", "batch"} {
		checkSearch(t, tr, w, true)
	}
	if !tr.StartsWith("appl") {
		t.Error("StartsWith(\"appl\") = false, want true because \"apply\" still uses that path")
	}
}

// Deleting a word that is a prefix of another word must keep the longer word.
func TestDeleteWordThatIsPrefixOfAnother(t *testing.T) {
	tr := newTrieWith(t, "app", "apple")

	if err := tr.Delete("app"); err != nil {
		t.Fatalf("Delete(\"app\") returned unexpected error: %v", err)
	}

	checkSearch(t, tr, "app", false)
	checkSearch(t, tr, "apple", true)
	checkSize(t, tr, 1)
	if !tr.StartsWith("app") {
		t.Error("StartsWith(\"app\") = false, want true because \"apple\" still uses that path")
	}
	checkWords(t, tr.WordsWithPrefix("app"), []string{"apple"}, "WordsWithPrefix(\"app\")")
}

// Deleting the longer word must keep the shorter word that ends on the path.
func TestDeleteLongerWordKeepsShorterWord(t *testing.T) {
	tr := newTrieWith(t, "app", "apple")

	if err := tr.Delete("apple"); err != nil {
		t.Fatalf("Delete(\"apple\") returned unexpected error: %v", err)
	}

	checkSearch(t, tr, "apple", false)
	checkSearch(t, tr, "app", true)
	checkSize(t, tr, 1)
	if tr.StartsWith("appl") {
		t.Error("StartsWith(\"appl\") = true, want false because the unused nodes should be pruned")
	}
}

// After the last word is gone, the trie must not keep any leftover node.
func TestDeletePrunesUnusedNodes(t *testing.T) {
	tr := newTrieWith(t, "apple")

	if err := tr.Delete("apple"); err != nil {
		t.Fatalf("Delete(\"apple\") returned unexpected error: %v", err)
	}

	checkSize(t, tr, 0)
	if tr.StartsWith("a") {
		t.Error("StartsWith(\"a\") = true, want false after the only word was deleted")
	}
	if len(tr.Root.Children) != 0 {
		t.Errorf("root has %d children after deleting the only word, want 0", len(tr.Root.Children))
	}
}

// A branch that belongs to another word must survive the delete.
func TestDeleteKeepsOtherBranch(t *testing.T) {
	tr := newTrieWith(t, "bad", "batch")

	if err := tr.Delete("batch"); err != nil {
		t.Fatalf("Delete(\"batch\") returned unexpected error: %v", err)
	}

	checkSize(t, tr, 1)
	checkSearch(t, tr, "bad", true)
	if tr.StartsWith("bat") {
		t.Error("StartsWith(\"bat\") = true, want false because only \"bad\" is left")
	}
	if !tr.StartsWith("ba") {
		t.Error("StartsWith(\"ba\") = false, want true because \"bad\" still uses that path")
	}
}

func TestDeleteReportsMissingWord(t *testing.T) {
	tr := newTrieWith(t, "app", "apple")

	tests := []struct {
		name string
		word string
	}{
		{name: "word never inserted", word: "banana"},
		{name: "prefix that is not a word", word: "appl"},
		{name: "longer than every word", word: "applesauce"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := tr.Delete(tc.word); !errors.Is(err, trie.ErrWordNotFound) {
				t.Errorf("Delete(%q) returned %v, want ErrWordNotFound", tc.word, err)
			}
		})
	}

	// A failed delete must not change the trie.
	checkSize(t, tr, 2)
	checkSearch(t, tr, "app", true)
	checkSearch(t, tr, "apple", true)
}

func TestDeleteOnEmptyTrie(t *testing.T) {
	tr := trie.NewTrie()

	if err := tr.Delete("apple"); !errors.Is(err, trie.ErrWordNotFound) {
		t.Errorf("Delete on an empty trie returned %v, want ErrWordNotFound", err)
	}
	checkSize(t, tr, 0)
}

func TestDeleteRejectsEmptyWord(t *testing.T) {
	tr := newTrieWith(t, "apple")

	if err := tr.Delete(""); !errors.Is(err, trie.ErrEmptyWord) {
		t.Errorf("Delete(\"\") returned %v, want ErrEmptyWord", err)
	}
	checkSize(t, tr, 1)
}

func TestDeleteTwiceReportsMissingWord(t *testing.T) {
	tr := newTrieWith(t, "apple")

	if err := tr.Delete("apple"); err != nil {
		t.Fatalf("first Delete(\"apple\") returned unexpected error: %v", err)
	}
	if err := tr.Delete("apple"); !errors.Is(err, trie.ErrWordNotFound) {
		t.Errorf("second Delete(\"apple\") returned %v, want ErrWordNotFound", err)
	}
	checkSize(t, tr, 0)
}

func TestInsertAfterDelete(t *testing.T) {
	tr := newTrieWith(t, "apple")

	if err := tr.Delete("apple"); err != nil {
		t.Fatalf("Delete(\"apple\") returned unexpected error: %v", err)
	}
	if err := tr.Insert("apple"); err != nil {
		t.Fatalf("Insert(\"apple\") returned unexpected error: %v", err)
	}

	checkSize(t, tr, 1)
	checkSearch(t, tr, "apple", true)
	checkWords(t, tr.WordsWithPrefix("a"), []string{"apple"}, "WordsWithPrefix(\"a\")")
}

func TestClear(t *testing.T) {
	tr := newTrieWith(t, sampleWords...)

	tr.Clear()

	checkSize(t, tr, 0)
	if tr.Root == nil {
		t.Fatal("Root is nil after Clear, want a usable root node")
	}
	if len(tr.Root.Children) != 0 {
		t.Errorf("root has %d children after Clear, want 0", len(tr.Root.Children))
	}
	for _, w := range sampleWords {
		checkSearch(t, tr, w, false)
	}

	// The trie must still work after Clear.
	if err := tr.Insert("apple"); err != nil {
		t.Fatalf("Insert after Clear returned unexpected error: %v", err)
	}
	checkSize(t, tr, 1)
	checkSearch(t, tr, "apple", true)
}

func TestWordsAreCaseSensitive(t *testing.T) {
	tr := newTrieWith(t, "Apple", "apple")

	checkSize(t, tr, 2)
	checkSearch(t, tr, "Apple", true)
	checkSearch(t, tr, "apple", true)
	checkWords(t, tr.WordsWithPrefix("A"), []string{"Apple"}, "WordsWithPrefix(\"A\")")
	checkWords(t, tr.WordsWithPrefix("a"), []string{"apple"}, "WordsWithPrefix(\"a\")")
}

// The trie walks over runes, so a multi byte character must count as one step.
// A byte based implementation fails these cases.
func TestUnicodeWords(t *testing.T) {
	words := []string{"한국", "한국어", "한글", "가나다", "🍎", "🍎🍏"}
	tr := newTrieWith(t, words...)

	checkSize(t, tr, len(words))
	for _, w := range words {
		checkSearch(t, tr, w, true)
	}

	checkSearch(t, tr, "한", false)
	if !tr.StartsWith("한") {
		t.Error("StartsWith(\"한\") = false, want true")
	}
	checkWords(t, tr.WordsWithPrefix("한국"), []string{"한국", "한국어"}, "WordsWithPrefix(\"한국\")")
	checkWords(t, tr.WordsWithPrefix("한"), []string{"한국", "한국어", "한글"}, "WordsWithPrefix(\"한\")")
	checkWords(t, tr.WordsWithPrefix("🍎"), []string{"🍎", "🍎🍏"}, "WordsWithPrefix(\"🍎\")")

	// "한" is three bytes long, so a single child must sit under the root for it.
	if _, ok := tr.Root.Children['한']; !ok {
		t.Error("root should have one child for the rune '한'")
	}

	if err := tr.Delete("한국"); err != nil {
		t.Fatalf("Delete(\"한국\") returned unexpected error: %v", err)
	}
	checkSearch(t, tr, "한국", false)
	checkSearch(t, tr, "한국어", true)
	checkSize(t, tr, len(words)-1)
}

func TestSizeTracksInsertAndDelete(t *testing.T) {
	tr := trie.NewTrie()

	steps := []struct {
		op       string
		word     string
		wantSize int
	}{
		{op: "insert", word: "app", wantSize: 1},
		{op: "insert", word: "apple", wantSize: 2},
		{op: "insert", word: "app", wantSize: 2},
		{op: "insert", word: "bat", wantSize: 3},
		{op: "delete", word: "app", wantSize: 2},
		{op: "delete", word: "apple", wantSize: 1},
		{op: "delete", word: "bat", wantSize: 0},
	}
	for i, step := range steps {
		switch step.op {
		case "insert":
			if err := tr.Insert(step.word); err != nil {
				t.Fatalf("step %d: Insert(%q) returned unexpected error: %v", i, step.word, err)
			}
		case "delete":
			if err := tr.Delete(step.word); err != nil {
				t.Fatalf("step %d: Delete(%q) returned unexpected error: %v", i, step.word, err)
			}
		}
		if got := tr.Size(); got != step.wantSize {
			t.Errorf("step %d (%s %q): Size() = %d, want %d", i, step.op, step.word, got, step.wantSize)
		}
	}
	checkSize(t, tr, 0)
	if len(tr.Root.Children) != 0 {
		t.Errorf("root has %d children after every word was deleted, want 0", len(tr.Root.Children))
	}
}

// A larger word set checks that the trie stays consistent when many words share
// their prefixes.
func TestManyWordsStayConsistent(t *testing.T) {
	var words []string
	for a := 'a'; a <= 'e'; a++ {
		for b := 'a'; b <= 'e'; b++ {
			for c := 'a'; c <= 'e'; c++ {
				words = append(words, string([]rune{a, b, c}))
			}
		}
	}

	tr := newTrieWith(t, words...)
	checkSize(t, tr, len(words))
	for _, w := range words {
		checkSearch(t, tr, w, true)
	}
	if got := tr.WordsWithPrefix("aa"); len(got) != 5 {
		t.Errorf("WordsWithPrefix(\"aa\") returned %d words, want 5", len(got))
	}

	// Delete every second word and check both groups again.
	var deleted, kept []string
	for i, w := range words {
		if i%2 == 0 {
			deleted = append(deleted, w)
		} else {
			kept = append(kept, w)
		}
	}
	for _, w := range deleted {
		if err := tr.Delete(w); err != nil {
			t.Fatalf("Delete(%q) returned unexpected error: %v", w, err)
		}
	}

	checkSize(t, tr, len(kept))
	for _, w := range deleted {
		checkSearch(t, tr, w, false)
	}
	for _, w := range kept {
		checkSearch(t, tr, w, true)
	}
	checkWords(t, tr.WordsWithPrefix(""), kept, "WordsWithPrefix(\"\")")
}

// A nil trie must report an error instead of panicking, which matches the way
// the other structures in this repository handle a nil receiver.
func TestNilTrieDoesNotPanic(t *testing.T) {
	var tr *trie.Trie

	if err := tr.Insert("apple"); !errors.Is(err, trie.ErrNilTrie) {
		t.Errorf("Insert on a nil trie returned %v, want ErrNilTrie", err)
	}
	if err := tr.Delete("apple"); !errors.Is(err, trie.ErrNilTrie) {
		t.Errorf("Delete on a nil trie returned %v, want ErrNilTrie", err)
	}
	if tr.Search("apple") {
		t.Error("Search on a nil trie = true, want false")
	}
	if tr.StartsWith("a") {
		t.Error("StartsWith on a nil trie = true, want false")
	}
	if got := tr.WordsWithPrefix("a"); len(got) != 0 {
		t.Errorf("WordsWithPrefix on a nil trie = %v, want no word", got)
	}
	if got := tr.Size(); got != 0 {
		t.Errorf("Size on a nil trie = %d, want 0", got)
	}
	if !tr.IsEmpty() {
		t.Error("IsEmpty on a nil trie = false, want true")
	}
	tr.Clear() // must not panic
}

// benchWords builds a deterministic word list for the benchmarks.
func benchWords() []string {
	words := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		words = append(words, fmt.Sprintf("word%04d", i))
	}
	return words
}

// BenchmarkInsert measures how long it takes to fill a fresh trie with the
// whole word set, so one iteration covers len(words) inserts.
func BenchmarkInsert(b *testing.B) {
	words := benchWords()
	b.ReportAllocs()

	for b.Loop() {
		tr := trie.NewTrie()
		for _, w := range words {
			_ = tr.Insert(w)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	words := benchWords()
	tr := trie.NewTrie()
	for _, w := range words {
		_ = tr.Insert(w)
	}
	b.ReportAllocs()

	i := 0
	for b.Loop() {
		tr.Search(words[i%len(words)])
		i++
	}
}

func BenchmarkWordsWithPrefix(b *testing.B) {
	words := benchWords()
	tr := trie.NewTrie()
	for _, w := range words {
		_ = tr.Insert(w)
	}
	b.ReportAllocs()

	for b.Loop() {
		tr.WordsWithPrefix("word01")
	}
}
