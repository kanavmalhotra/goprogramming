package trie

type Trie interface {
	// Insert inserts a word into the trie and returns true
	// if the word already exists.
	Insert(word string) bool

	// SearchLongestPrefix returns true if the prefix or the exact
	// match of the passed string is found, otherwise return blank
	// string and false.
	SearchLongestPrefix(word string) (string, bool)

	// Size returns the total size of the Trie or rather, we can say
	// returns the list of characters in the trie
	Size() int32
}
