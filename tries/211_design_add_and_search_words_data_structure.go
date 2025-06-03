package tries

type WordNode struct {
	children map[rune]*WordNode
	word     bool
}

type WordDictionary struct {
	root *WordNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &WordNode{
			children: make(map[rune]*WordNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, c := range word {
		if _, exists := curr.children[c]; !exists {
			curr.children[c] = &WordNode{
				children: make(map[rune]*WordNode),
			}
		}
		curr = curr.children[c]
	}
	curr.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return false
}
