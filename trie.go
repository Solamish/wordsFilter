package wordsfilter

const (
	INIT_TRIE_CHILDREN_NUM = 128 // Since we need to deal all kinds of language, so we use 128 instead of 26
)

type trieNode struct {
	isEndOfWord bool
	children map[rune]*trieNode
}

func newTrieNode()  *trieNode{
	return &trieNode{
		isEndOfWord: false,
		children:    make(map[rune]*trieNode,INIT_TRIE_CHILDREN_NUM),
	}
}



 

