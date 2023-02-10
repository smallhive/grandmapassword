package passwd

type ProcessedWord struct {
	Word       string
	Difficulty int
}

type ProcessedWordSlice []ProcessedWord

func (a ProcessedWordSlice) Len() int           { return len(a) }
func (a ProcessedWordSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ProcessedWordSlice) Less(i, j int) bool { return a[i].Difficulty < a[j].Difficulty }
