package swgohgg

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Mod represents a player mod with all mod information.
type Mod struct {
	ID       string
	Level    int
	Rarity   int
	Shape    string
	BonusSet string

	PrimStat ModStat
	SecStat  []ModStat

	UsingIn string
}

func (m Mod) String() string {
	return m.Format(false)
}

// Format returns a string representation of the mod, either using Unicode emoji or not.
func (m *Mod) Format(useEmoji bool) string {
	if m == nil {
		return "nil mod"
	}
	icon := m.ShapeIcon()
	if useEmoji {
		icon = m.ShapeEmoji()
	}
	str := fmt.Sprintf("%s %-9s L%-2d %d* %v %v", icon, m.BonusShortName(), m.Level, m.Rarity, m.PrimStat, m.SecStat)
	if m.UsingIn != "" {
		str += " (" + m.UsingIn + ")"
	}
	return str
}

// BonusShortName returns the bonus set abreviated name.
func (m *Mod) BonusShortName() string {
	return statAbbrev(m.BonusSet)
}

// ShapeEmoji returns an unicode emoji that better represents the mod shape.
func (m *Mod) ShapeEmoji() string {
	switch m.Shape {
	case "Transmitter":
		return "​◼️"
	case "Processor":
		return "​♦️"
	case "Holo-Array":
		return "⚠️"
	case "Data-Bus":
		return "​⚫️"
	case "Receiver":
		return "​↗️"
	case "Multiplexer":
		return "​➕"
	default:
		return m.Shape
	}
}

// ShapeIcon returns an unicode emoji that is coloured and represents the mod shape.
func (m *Mod) ShapeIcon() string {
	switch m.Shape {
	case "Transmitter":
		return "◻"
	case "Processor":
		return "◇"
	case "Holo-Array":
		return "△"
	case "Data-Bus":
		return "○"
	case "Receiver":
		return "◹"
	case "Multiplexer":
		return "+"
	default:
		return m.Shape
	}
}

// ShapeName returns the mod shape name.
func (m *Mod) ShapeName() string {
	switch m.Shape {
	case "Transmitter":
		return "Square"
	case "Processor":
		return "Diamond"
	case "Holo-Array":
		return "Triangle"
	case "Data-Bus":
		return "Circle"
	case "Receiver":
		return "Arrow"
	case "Multiplexer":
		return "Cross"
	default:
		return m.Shape
	}
}

// HasStat checks if the mod has a primary or secondary stat.
func (m *Mod) HasStat(stat string) bool {
	return !m.GetStat(stat).IsZero()
}

// GetStat returns a stat from the mod. The return is never nil, but a Zero ModStat.
func (m *Mod) GetStat(stat string) ModStat {
	if m.PrimStat.Stat == stat || m.PrimStat.StatShortName() == stat {
		return m.PrimStat
	}
	for _, sec := range m.SecStat {
		if sec.Stat == stat || sec.StatShortName() == stat {
			return sec
		}
	}
	return ModStat{}
}

// ModStat is a single mod stat value. If a percentage value additive,
// the IsPercentage attribute will be true.
type ModStat struct {
	Stat      string
	Value     float64
	IsPercent bool
}

func (ms ModStat) String() string {
	if ms.IsPercent {
		return fmt.Sprintf("%.02f%% %s", ms.Value, ms.StatShortName())
	}
	return fmt.Sprintf("%.02f %s", ms.Value, ms.StatShortName())
}

// StatShortName returns an abbreviated stat name.
func (ms ModStat) StatShortName() string {
	return statAbbrev(ms.Stat)
}

// IsBetterThan attempts to compare two stats with each other.
func (ms ModStat) IsBetterThan(other ModStat) bool {
	switch {
	case ms.IsZero():
		return false
	case ms.IsPercent && !other.IsPercent:
		return false
	case !ms.IsPercent && other.IsPercent:
		return true
	default:
		return ms.Value > other.Value
	}
}

// IsZero checks if the stat object has a value or name, indicating it is not zero value.
func (ms ModStat) IsZero() bool {
	return ms.Stat == "" && ms.Value == 0
}

func statAbbrev(stat string) string {
	switch stat {
	case "Critical Chance":
		return "Crit Chan"
	case "Critical Damage":
		return "Crit Dam"
	case "Critical Avoidance":
		return "Crit Avoi"
	case "Protection":
		return "Prot"
	default:
		return stat
	}
}

// ModFilter allows to filter mod queries.
// Currently available filter is by character name.
type ModFilter struct {
	Char string
}

// Match checks if the provided mod matches the filter.
func (f *ModFilter) Match(mod *Mod) bool {
	if f.Char == "" {
		return true
	}
	return CharSlug(CharName(f.Char)) == CharSlug(mod.UsingIn)
}

// ModCollection is a slice of mods with extra methods for manipulation.
type ModCollection []*Mod

// ByShape filter the mod collection by a given shape.
func (m ModCollection) ByShape(shape string) (filtered ModCollection) {
	for _, mod := range m {
		if strings.ToLower(mod.Shape) == strings.ToLower(shape) {
			filtered = append(filtered, mod)
		}
	}
	log.Printf("DEBUG: Found %d %s", len(filtered), shape)
	return
}

// WithStat filters a mod collection by a given stat.
func (m ModCollection) WithStat(stat string) (filtered ModCollection) {
	for _, mod := range m {
		if mod.HasStat(stat) || mod.BonusSet == stat {
			filtered = append(filtered, mod)
		}
	}
	log.Printf("DEBUG: Found %d with %s", len(filtered), stat)
	return
}

// MinLevel filters the mod collection to one that has only the given minimum level.
func (m ModCollection) MinLevel(level int) (filtered ModCollection) {
	for _, mod := range m {
		if mod.Level >= level {
			filtered = append(filtered, mod)
		}
	}
	log.Printf("DEBUG: Found %d with level %d", len(filtered), level)
	return
}

// MinRarity filters the mod collection to one that has the minium rarity.
func (m ModCollection) MinRarity(rarity int) (filtered ModCollection) {
	for _, mod := range m {
		if mod.Rarity >= rarity {
			filtered = append(filtered, mod)
		}
	}
	log.Printf("DEBUG: Found %d with ratity %d", len(filtered), rarity)
	return
}

// Filter applies the provided mod filter and returns a filtered collection that matches it.
func (m ModCollection) Filter(filter ModFilter) (filtered ModCollection) {
	for _, mod := range m {
		if filter.Match(mod) {
			filtered = append(filtered, mod)
		}
	}
	return filtered
}

// SetWith suggests a set containing the max values of the provided stat.
func (m ModCollection) SetWith(stat string) ModSet {
	set := make(map[string]Mod)
	// Loop over all mods and find the best set for the given stat
	for i := range m {
		mod := m[i]
		// Check if the mod has the stat
		if mod.HasStat(stat) {
			curr := set[mod.Shape]
			currStat := curr.GetStat(stat)
			thisStat := mod.GetStat(stat)
			isBetter := currStat.IsBetterThan(thisStat)
			if !isBetter {
				set[mod.Shape] = *mod
			}
		}
	}
	return set
}

// Optimize searches over all your mods with level >= 12, rarity >= 4,
// and outputs a best-value for the given stat, considering
// bonus sets if the 'percentIsBetter' parameter is true.
// This is very experimental, CPU intensive and memory intensive!
func (m ModCollection) Optimize(stat string, percentIsBetter bool) ModSet {
	switch stat {
	case "Accuracy", "Critical Damage", "Critical Chance", "Tenacity", "Potency":
		percentIsBetter = true
	}
	optimized := ModSet{}
	log.Printf("DEBUG: Combining all possible mod sets for %s. This may take a while...", stat)

	transmitter := m.ByShape("Transmitter").WithStat(stat).MinLevel(12).MinRarity(4)
	receiver := m.ByShape("Receiver").WithStat(stat).MinLevel(12).MinRarity(4)
	processor := m.ByShape("Processor").WithStat(stat).MinLevel(12).MinRarity(4)
	holoArray := m.ByShape("Holo-Array").WithStat(stat).MinLevel(12).MinRarity(4)
	dataBus := m.ByShape("Data-Bus").WithStat(stat).MinLevel(12).MinRarity(4)
	multiplexer := m.ByShape("Multiplexer").WithStat(stat).MinLevel(12).MinRarity(4)
	totalSets := len(transmitter) * len(receiver) * len(processor) * len(holoArray) * len(dataBus) * len(multiplexer)
	log.Printf("DEBUG: Analysing %d sets", totalSets)

	count := 0
	for _, t := range transmitter {
		for _, r := range receiver {
			for _, p := range processor {
				for _, h := range holoArray {
					for _, d := range dataBus {
						for _, mu := range multiplexer {
							count++
							if count%1000000 == 0 {
								log.Printf("Processed %.02f%%", 100*(float64(count)/float64(totalSets)))
							}
							set := ModSet{}
							set.Add(t)
							set.Add(r)
							set.Add(p)
							set.Add(h)
							set.Add(d)
							set.Add(mu)
							if set.Sum(stat, percentIsBetter) > optimized.Sum(stat, percentIsBetter) {
								optimized = set
							}
						}
					}
				}
			}
		}
	}
	return optimized
}

// ModSet is a mod with a set of mods with unique shapes.
type ModSet map[string]Mod

// Add includes the provided mod in the set.
func (set ModSet) Add(mod *Mod) {
	if mod == nil {
		return
	}
	set[mod.Shape] = *mod
}

// AddAll include all mods in the set.
func (set ModSet) AddAll(mods []*Mod) {
	for i := range mods {
		set.Add(mods[i])
	}
}

// StatSummary returns a summary of all mods in the set.
func (set ModSet) StatSummary() (result []string) {
	for _, stat := range StatNames {
		result = append(result, fmt.Sprintf("%.0f + %.02f%% %s", set.Sum(stat, false), set.Sum(stat, true), stat))
	}
	return
}

// Sum returns a sum of all stats in the given set.
// Only mods with the same stat name and ispercent flag will be included in the sum.
func (set ModSet) Sum(stat string, isPercent bool) (res float64) {
	// First, acumulate the stat value
	for _, mod := range set {
		stat := mod.GetStat(stat)
		if stat.IsZero() || stat.IsPercent != isPercent {
			continue
		}
		res += stat.Value
	}
	// Second, acumulate the bonus set values
	if isPercent {
		bonus := set.BonusForSet(stat)
		res += bonus
	}
	return
}

// BonusForSet calculates the bonus set for the given stat.
// Takes into account mod level.
func (set ModSet) BonusForSet(stat string) float64 {
	mods := make([]*Mod, 0, len(set))
	for i := range set {
		m := set[i]
		mods = append(mods, &m)
	}
	sort.Sort(sortByLevel{mods: mods, asc: false})
	bonus := 0.0
	maxed := true
	count := 0
	for _, mod := range mods {
		if mod.BonusSet == stat {
			if mod.Level < 15 {
				maxed = false
			}
			count++

			bonusVal, required := 0.0, 0
			switch stat {
			case "Health", "Defense", "Critical Chance":
				bonusVal, required = 5, 2
			case "Tenacity", "Potency":
				bonusVal, required = 10, 2
			case "Offense", "Speed":
				bonusVal, required = 10, 4
			case "Critical Damage":
				bonusVal, required = 30, 4
			}

			if count == required {
				// We got a bonus
				if maxed {
					bonus += bonusVal
				} else {
					bonus += bonusVal / 2
				}
				count = 0
				maxed = true
			}
		}
	}
	return bonus
}

// Mods reutrns a mod collection of mods with the given filter.
func (c *Client) Mods(filter ModFilter) (mods ModCollection, err error) {
	allyCode := c.AllyCode()
	page := 1
	for {
		url := fmt.Sprintf("https://swgoh.gg/p/%s/mods/?page=%d", allyCode, page)
		resp, err := c.hc.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
		count := 0
		doc.Find(".collection-mod").Each(func(i int, s *goquery.Selection) {
			mod := parseMod(s)
			mods = append(mods, mod)
			count++
		})
		// Soft-disabled mods for more than one page while we are not fetching from the API
		if count < (12*3) || page > 2 {
			break
		}
		page++
	}
	mods = mods.Filter(filter)
	sort.Sort(sortByShape{mods: mods, asc: true})
	return mods, nil
}

func parseMod(s *goquery.Selection) *Mod {
	var err error
	mod := &Mod{}
	mod.ID = s.AttrOr("data-id", "")
	mod.Level, err = strconv.Atoi(s.Find(".statmod-level").Text())
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	mod.Rarity = s.Find(".statmod-pip").Length()
	shortname := strings.Fields(s.Find(".statmod-img").AttrOr("alt", "!Unkown!"))
	switch len(shortname) {
	case 4:
		mod.BonusSet = shortname[2]
		mod.Shape = shortname[3]
	case 5:
		mod.BonusSet = shortname[2] + " " + shortname[3]
		mod.Shape = shortname[4]
	default:
		mod.BonusSet = "?"
		mod.Shape = "?"
	}

	// Primary stat
	mod.PrimStat = parseStat(s.Find(".statmod-stats-1 .statmod-stat"))
	// Secondary stats
	s.Find(".statmod-stats-2 .statmod-stat").Each(func(i int, stat *goquery.Selection) {
		mod.SecStat = append(mod.SecStat, parseStat(stat))
	})

	mod.UsingIn = s.Find("img.char-portrait-img").AttrOr("alt", "")
	return mod
}

func parseStat(s *goquery.Selection) (stat ModStat) {
	stat.Stat = s.Find(".statmod-stat-label").Text()

	strvalue := s.Find(".statmod-stat-value").Text()
	strvalue = strings.Replace(strvalue, "%", "", -1)
	strvalue = strings.Replace(strvalue, "+", "", -1)

	var err error
	stat.Value, err = strconv.ParseFloat(strvalue, 64)
	if err != nil {
		log.Printf("parsestat: invalid value %s", s.Find(".statmod-stat-value").Text())
	}
	stat.IsPercent = strings.Contains(s.Find(".statmod-stat-value").Text(), "%")
	return stat
}
