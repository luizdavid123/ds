package trie_test

import (
	"ds/misc"
	"ds/trie"
	"testing"
)

func TestTrieBasic(t *testing.T) {
	trie := trie.New()
	trie.Insert("david")

	ok := trie.Find("david")
	misc.Equals(t, true, ok)
	ok = trie.Find("apple")
	misc.Equals(t, false, ok)

	ok = trie.StartWith("da")
	misc.Equals(t, true, ok)
	ok = trie.StartWith("app")
	misc.Equals(t, false, ok)
}

func TestTrieFind(t *testing.T) {
	tests := map[string]struct {
		word string
		ok   bool
	}{
		"case 01": {"sam", true},
		"case 02": {"john", true},
		"case 03": {"dogg", true},
		"case 04": {"dog", true},
		"case 05": {"cat", true},
		"case 06": {"doggy", false},
		"case 07": {"house", false},
		"case 08": {"", false},
		"case 09": {"sa", false},
	}
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	trie := trie.New()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ok := trie.Find(test.word)
			misc.Equals(t, test.ok, ok)
		})
	}
}

func TestTrieStartWith(t *testing.T) {
	tests := map[string]struct {
		word string
		ok   bool
	}{
		"case 01": {"dog", true},
		"case 02": {"doggy", false},
		"case 03": {"sa", true},
		"case 04": {"john", true},
		"case 05": {"dogg", true},
		"case 06": {"rose", true},
		"case 07": {"a", false},
		"case 08": {"go", false},
		"case 09": {"tam", false},
		"case 10": {"hello", false},
		"case 11": {"rosy", false},
	}
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	trie := trie.New()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ok := trie.StartWith(test.word)
			misc.Equals(t, test.ok, ok)
		})
	}
}
