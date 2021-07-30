package trie

/*
A trie is a data structure for efficient information retrieval. It is a special kind of tree where a path
starting from root to a particular node can define a word that is stored in this tree. A trie can be built
for entire ASCII_SIZE, ALPHABETS, NUMBERS depending upon the use case.
*/

// Constant
const (
	ALBHABET_SIZE = 26
)

// TrieNode is a node for a trie
type TrieNode struct {
	child [ALBHABET_SIZE]*TrieNode
	end   bool
}

// A trie is a data structure for efficient information retrieval
type Trie struct {
	root *TrieNode
}

// New return a trie
func New() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

// Insert insert a word into the trie
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if cur.child[j] == nil {
			cur.child[j] = &TrieNode{}
		}
		cur = cur.child[j]
	}
	cur.end = true
}

// Find find a word in the trie
func (t *Trie) Find(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if cur.child[j] == nil {
			return false
		}
		cur = cur.child[j]
	}
	return cur.end
}

// StartWith check if there is any word in the trie starts with the given prefix
func (t *Trie) StartWith(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		j := prefix[i] - 'a'
		if cur.child[j] == nil {
			return false
		}
		cur = cur.child[j]
	}
	return true
}
