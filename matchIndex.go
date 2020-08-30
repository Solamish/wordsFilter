package wordsfilter

type matchIndex struct {
	start int // start index
	end   int // end index
}

func newMatchIndex(start, end int) *matchIndex {
	return &matchIndex{
		start: start,
		end:   end,
	}
}

// Construct from existing match index object
func buildMatchIndex(m *matchIndex) *matchIndex {
	return &matchIndex{
		start: m.start,
		end:   m.end,
	}
}

