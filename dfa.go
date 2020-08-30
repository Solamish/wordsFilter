package wordsfilter

// dfa util
type DFAUtil struct {
	// The root node
	root *trieNode
}

func (dfa *DFAUtil) insertWord(word []rune) {
	currNode := dfa.root
	for _, c := range word {
		if childNode, exist := currNode.children[c]; !exist {
			childNode = newTrieNode()
			currNode.children[c] = childNode
			currNode = childNode
		} else {
			currNode = childNode
		}
	}

	currNode.isEndOfWord = true
}

func (dfa *DFAUtil) startsWith(prefix []rune) bool {
	currNode := dfa.root
	for _, c := range prefix {
		if childNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = childNode
		}
	}
	return true
}

func (dfa *DFAUtil) searchWord(word []rune) bool {
	currNode := dfa.root
	for _,c := range word{
		if childNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = childNode
		}
	}
	return currNode.isEndOfWord
}

func (dfa *DFAUtil) searchSentence(sentence string) (matchIndexList []*matchIndex) {
	start, end := 0,1
	sentenceRuneList := []rune(sentence)

	startsWith := false
	for end <= len(sentenceRuneList) {
		// Check if a sensitive word starts with word range from [start:end)
		if dfa.startsWith(sentenceRuneList[start:end]) {
			startsWith = true
			end += 1
		} else {
			if startsWith == true {
				// Check any sub word is the sensitive word from long to short
				for index := end - 1; index > start; index-- {
					if dfa.searchWord(sentenceRuneList[start:index]) {
						matchIndexList = append(matchIndexList, newMatchIndex(start, index-1))
						break
					}
				}
			}
			start, end = end-1, end+1
			startsWith = false
		}
	}

	// If finishing not because of unmatching, but reaching the end, we need to
	// check if the previous startsWith is true or not.
	// If it's true, we need to check if there is any candidate?
	if startsWith {
		for index := end - 1; index > start; index-- {
			if dfa.searchWord(sentenceRuneList[start:index]) {
				matchIndexList = append(matchIndexList, newMatchIndex(start, index-1))
				break
			}
		}
	}

	return
}

func (dfa *DFAUtil) IsMatch(sentence string) bool {
	return len(dfa.searchSentence(sentence)) > 0
}

func (dfa *DFAUtil) HandleWord(sentence string, repalceCh rune) string {
	matchIndexList := dfa.searchSentence(sentence)
	if len(matchIndexList) == 0 {
		return sentence
	}


	sentenceList := []rune(sentence)
	for _, matchIndexObj := range matchIndexList {
		for index := matchIndexObj.start; index <= matchIndexObj.end; index++ {
			sentenceList[index] = repalceCh
		}
	}
	return string(sentenceList)
}

func NewDFAUtil(wordList []string) *DFAUtil {
	dfa := &DFAUtil{
		newTrieNode(),
	}
	for _,word := range wordList {
		wordRuneList := []rune(word)
		if len(wordRuneList) > 0 {
			dfa.insertWord(wordRuneList)
		}
	}
	return dfa
}


