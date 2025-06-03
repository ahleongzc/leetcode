package tries

type TrieNode struct {
	children map[rune]*TrieNode
	word     bool
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, r := range word {
		if _, exists := curr.children[r]; !exists {
			curr.children[r] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.children[r]
	}
	curr.word = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, r := range word {
		if _, exists := curr.children[r]; !exists {
			return false
		}
		curr = curr.children[r]
	}
	return curr.word
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, r := range prefix {
		if _, exists := curr.children[r]; !exists {
			return false
		}
		curr = curr.children[r]
	}
	return true
}
