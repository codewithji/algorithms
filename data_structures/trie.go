package datastructures

// A trie is a special form of N-ary tree. It's typically used to store strings.
// The root node is associated with an empty string.

// Important property is that all descendants of a node have a common prefix
// of the string associated with that node. That's why it's also called a prefix tree.

// You can store the children nodes in the following way:
// 1) Array of size 26 that goes from 'a' to 'z'
// Reads would be incredibly fast but unnecessary waste of space may occur

// 2) Map with key, value being chars => corresponding child node
// Saves space and flexible since we are not dealing with fixed length

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
