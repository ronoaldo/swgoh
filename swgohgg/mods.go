package swgohgg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

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

func (m *Mod) String() string {
	return m.Format(false)
}

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

func (m *Mod) BonusShortName() string {
	return statAbbrev(m.BonusSet)
}

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

func (m *Mod) HasStat(stat string) bool {
	if m.PrimStat.Stat == stat || m.PrimStat.StatShortName() == stat {
		return true
	}
	for _, sec := range m.SecStat {
		if sec.Stat == stat || sec.StatShortName() == stat {
			return true
		}
	}
	return false
}

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

func (ms ModStat) StatShortName() string {
	return statAbbrev(ms.Stat)
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

type ModFilter struct {
	Char string
}

func (f *ModFilter) Match(mod *Mod) bool {
	if f.Char == "" {
		return true
	}
	return f.Char == mod.UsingIn
}

type ModCollection []*Mod

func (c *Client) Mods(filter ModFilter) (mods ModCollection, err error) {
	page := 1
	for {
		url := fmt.Sprintf("https://swgoh.gg/u/%s/mods/?page=%d", c.profile, page)
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
			if filter.Match(mod) {
				mods = append(mods, mod)
			}
			count++
		})
		if count < 60 {
			break
		}
		page++
	}
	return mods, nil
}

func parseMod(s *goquery.Selection) *Mod {
	var err error
	mod := &Mod{}
	mod.ID = s.AttrOr("data-id", "")
	mod.Level, err = strconv.Atoi(s.Find(".statmod-level").Text())
	if err != nil {
		log.Println("Error: %v", err)
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
