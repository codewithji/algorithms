package datastructures

type Trie struct {
	isEnd    bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: make(map[rune]*Trie),
	}
}

func (trie *Trie) Insert(word string) {
	currNode := trie

	for _, char := range word {
		if _, found := currNode.children[char]; !found {
			// if a child trie is NOT found
			newChild := &Trie{
				isEnd:    false,
				children: make(map[rune]*Trie),
			}

			currNode.children[char] = newChild
		}
		// child node will be populated either way
		currNode = currNode.children[char]
	}

	currNode.isEnd = true
}

/** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	currNode := trie

	// go through each char
	// make sure there is a key to every given char
	// at the end, make sure the isEnd propety is true
	for _, char := range word {
		if _, found := currNode.children[char]; !found {
			return false
		}

		currNode = currNode.children[char]
	}

	return currNode.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	currNode := trie

	for _, char := range prefix {
		if _, found := currNode.children[char]; !found {
			return false
		}

		currNode = currNode.children[char]
	}

	return true
}
