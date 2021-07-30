package trie_test

import (
	"ds/misc"
	"ds/trie"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := trie.New()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	prefixs := []string{"s", "d", "jo", "dog", "rose"}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	for i := 0; i < len(words); i++ {
		misc.Equals(t, true, trie.Find(words[i]))
	}
	for i := 0; i < len(words); i++ {
		misc.Equals(t, true, trie.StartWith(words[i]))
	}
	for i := 0; i < len(prefixs); i++ {
		misc.Equals(t, true, trie.StartWith(words[i]))
	}

	nwords := []string{"sa", "t", "", "doggy", "house"}
	nprefixs := []string{"z", "go", "tam", "doggy"}
	for i := 0; i < len(nwords); i++ {
		misc.Equals(t, false, trie.Find(nwords[i]))
	}
	for i := 0; i < len(nprefixs); i++ {
		misc.Equals(t, false, trie.StartWith(nprefixs[i]))
	}
}
