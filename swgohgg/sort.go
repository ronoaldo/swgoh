package swgohgg

// SortByStars is a sorting criteria to sort a character collection by star level.
type SortByStars struct {
	chars []*Char
	asc   bool
}

// ByStars returns a sorting criteria to sort characters by stars.
func ByStars(chars []*Char, ascending bool) SortByStars {
	return SortByStars{
		chars: chars,
		asc:   ascending,
	}
}

func (bs SortByStars) Len() int { return len(bs.chars) }

func (bs SortByStars) Swap(i, j int) { bs.chars[i], bs.chars[j] = bs.chars[j], bs.chars[i] }

func (bs SortByStars) Less(i, j int) bool {
	if bs.asc {
		return bs.chars[i].Stars < bs.chars[j].Stars
	}
	return bs.chars[i].Stars > bs.chars[j].Stars
}

// SortByShipStars is a sorting criteria to sort ship collection by stars.
type SortByShipStars struct {
	ships []*Ship
	asc   bool
}

// ByShipStars returns a sorting criteria to sort ships by stars.
func ByShipStars(ships []*Ship, ascending bool) SortByShipStars {
	return SortByShipStars{
		ships: ships,
		asc:   ascending,
	}
}

func (bs SortByShipStars) Len() int { return len(bs.ships) }

func (bs SortByShipStars) Swap(i, j int) { bs.ships[i], bs.ships[j] = bs.ships[j], bs.ships[i] }

func (bs SortByShipStars) Less(i, j int) bool {
	if bs.asc {
		return bs.ships[i].Stars < bs.ships[j].Stars
	}
	return bs.ships[i].Stars > bs.ships[j].Stars
}

type sortByShape struct {
	mods []*Mod
	asc  bool
}

func (s sortByShape) Len() int      { return len(s.mods) }
func (s sortByShape) Swap(i, j int) { s.mods[i], s.mods[j] = s.mods[j], s.mods[i] }

func (s sortByShape) Less(i, j int) bool {
	var less bool
	a, b := s.mods[i], s.mods[j]
	if a.Shape == b.Shape {
		// If shapes are equal, compare ids
		less = a.ID < b.ID
	} else {
		less = shapes[a.Shape] < shapes[b.Shape]
	}

	if s.asc {
		return less
	}
	return !less
}

type sortByLevel struct {
	mods []*Mod
	asc  bool
}

func (s sortByLevel) Len() int      { return len(s.mods) }
func (s sortByLevel) Swap(i, j int) { s.mods[i], s.mods[j] = s.mods[j], s.mods[i] }

func (s sortByLevel) Less(i, j int) bool {
	var less bool
	a, b := s.mods[i], s.mods[j]
	less = a.Level < b.Level
	if s.asc {
		return less
	}
	return !less
}
