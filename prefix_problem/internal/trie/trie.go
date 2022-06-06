package trie

// trieNode represents each node of the linked list references
type trieNode struct {
	children map[rune]*trieNode
	islast   bool
}

// prefix Trie maintains the information of the
// Tries structure which is build as a cache
type prefixTrie struct {
	root *trieNode
	size int32
}

// A method that takes a string as a parameter and
// adds it to the trie. In case of existing word,
// it return true
func (t *prefixTrie) Insert(word string) bool {
	exists := true
	current := t.root
	for _, letter := range word {
		n, ok := current.children[letter]
		if !ok {
			exists = false

			n = &trieNode{
				make(map[rune]*trieNode),
				false,
			}
			current.children[letter] = n
		}

		current = n
	}
	current.islast = true

	if !exists {
		t.size++
	}

	return exists
}

// A method that takes a string as a parameter and returns the longest
// prefix that matches the input string
func (t *prefixTrie) SearchLongestPrefix(word string) (string, bool) {
	current := t.root
	prefix := []rune{}

	prev := t.root
	var prev_prefix rune

	for _, letter := range word {
		n, ok := current.children[letter]
		if ok {
			prefix = append(prefix, letter)
			prev = current
			prev_prefix = letter
			current = n
			continue
		} else if prev.children[prev_prefix].islast {
			return string(prefix), true
		} else {
			break
		}
	}
	return "", false
}

// A method that returns the size of the prefix Trie
func (t *prefixTrie) Size() int32 {
	return t.size
}

// NewPrefixTrie creates a new prefixTrie object.
func NewPrefixTrie() Trie {
	return &prefixTrie{
		&trieNode{
			children: make(map[rune]*trieNode),
			islast:   false,
		}, 0}
}
