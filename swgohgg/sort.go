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

type sortByShipStars struct {
	ships []*Ship
	asc   bool
}

func ByShipStars(ships []*Ship, ascending bool) sortByShipStars {
	return sortByShipStars{
		ships: ships,
		asc:   ascending,
	}
}

func (bs sortByShipStars) Len() int { return len(bs.ships) }

func (bs sortByShipStars) Swap(i, j int) { bs.ships[i], bs.ships[j] = bs.ships[j], bs.ships[i] }

func (bs sortByShipStars) Less(i, j int) bool {
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
