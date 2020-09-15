package wordsfilter

import "log"

type Filter struct {
	filter DFAUtil
	store  MemoryStore
}

func NewFilter(opts Opts) *Filter{
	memStore := NewMemoryStore(opts)
	if memStore == nil {
		log.Panic("init memory store engine error")
	}

	wordList := memStore.Read()
	dfa := NewDFAUtil(wordList)
	filter := &Filter{
		filter:	*dfa,
		store:  *memStore,
	}

	return filter
}

func (f *Filter) IsMatch(sentence string) bool {
	return f.filter.IsMatch(sentence)
}

func (f *Filter) HandledWord(sentence string, replace rune) string{
	return f.filter.HandleWord(sentence,replace)
}
