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
	if m == nil {
		return "nil mod"
	}
	str := fmt.Sprintf("%s %-18s L%d %d* %v %v", m.ShapeIcon(), m.BonusSet, m.Level, m.Rarity, m.PrimStat, m.SecStat)
	if m.UsingIn != "" {
		str += " (" + m.UsingIn + ")"
	}
	return str
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
		return "Quadrado "
	case "Processor":
		return "Losango  "
	case "Holo-Array":
		return "Triangulo"
	case "Data-Bus":
		return "Circulo  "
	case "Receiver":
		return "Seta     "
	case "Multiplexer":
		return "Cruz     "
	default:
		return m.Shape
	}
}

type ModStat struct {
	Type         string
	Value        float64
	ValuePercent bool
}

func (ms ModStat) String() string {
	if ms.ValuePercent {
		return fmt.Sprintf("%.02f%% %s", ms.Value, ms.Type)
	}
	return fmt.Sprintf("%.02f %s", ms.Value, ms.Type)
}

type ModCollection []*Mod

func (c *Client) Mods() (mods ModCollection, err error) {
	url := fmt.Sprintf("https://swgoh.gg/u/%s/mods/", c.profile)
	resp, err := c.hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".collection-mod").Each(func(i int, s *goquery.Selection) {
		mod := parseMod(s)
		mods = append(mods, mod)
	})
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
	stat.Type = s.Find(".statmod-stat-label").Text()

	strvalue := s.Find(".statmod-stat-value").Text()
	strvalue = strings.Replace(strvalue, "%", "", -1)
	strvalue = strings.Replace(strvalue, "+", "", -1)

	var err error
	stat.Value, err = strconv.ParseFloat(strvalue, 64)
	if err != nil {
		log.Printf("parsestat: invalid value %s", s.Find(".statmod-stat-value").Text())
	}
	stat.ValuePercent = strings.Contains(s.Find(".statmod-stat-value").Text(), "%")
	return stat
}
