package swgohgg

type sortByStars struct {
	chars []*Char
	asc   bool
}

func ByStars(chars []*Char, ascending bool) sortByStars {
	return sortByStars{
		chars: chars,
		asc:   ascending,
	}
}

func (bs sortByStars) Len() int { return len(bs.chars) }

func (bs sortByStars) Swap(i, j int) { bs.chars[i], bs.chars[j] = bs.chars[j], bs.chars[i] }

func (bs sortByStars) Less(i, j int) bool {
	if bs.asc {
		return bs.chars[i].Stars < bs.chars[j].Stars
	}
	return bs.chars[i].Stars > bs.chars[j].Stars
}
