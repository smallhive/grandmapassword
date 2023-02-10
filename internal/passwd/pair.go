package passwd

type Pair struct {
	First  int
	Second int
	Sum    int
}

type PairSlice []Pair

func (a PairSlice) Len() int      { return len(a) }
func (a PairSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PairSlice) Less(i, j int) bool {
	return a[i].Sum < a[j].Sum
}
