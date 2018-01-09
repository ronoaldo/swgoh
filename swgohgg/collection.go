package swgohgg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Collection() (collection Collection, err error) {
	url := fmt.Sprintf("https://swgoh.gg/u/%s/collection/", c.profile)
	doc, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	doc.Find(".collection-char-list .collection-char").Each(func(i int, s *goquery.Selection) {
		char := parseChar(s)
		if !collection.Contains(char.Name) {
			collection = append(collection, char)
		}
	})
	sort.Sort(ByStars(collection, false))
	return collection, nil
}

type Collection []*Char

func (r Collection) Contains(char string) bool {
	for i := range r {
		if strings.ToLower(r[i].Name) == strings.ToLower(char) {
			return true
		}
	}
	return false
}

func (r Collection) ContainsAll(chars ...string) bool {
	for _, char := range chars {
		if !r.Contains(char) {
			return false
		}
	}
	return true
}

type Char struct {
	Name  string
	Stars int
	Level int
	Gear  int
}

func (c *Char) String() string {
	if c == nil {
		return "nil"
	}
	return fmt.Sprintf("%s %d* G%d Lvl%d", c.Name, c.Stars, c.Gear, c.Level)
}

func parseChar(s *goquery.Selection) *Char {
	var char Char
	char.Name = s.Find(".collection-char-name-link").Text()
	char.Level, _ = strconv.Atoi(s.Find(".char-portrait-full-level").Text())
	char.Gear = gearLevel(s)
	char.Stars = stars(s)
	return &char
}

func stars(s *goquery.Selection) int {
	level := 0
	s.Find(".star").Each(func(i int, star *goquery.Selection) {
		if star.HasClass("star-inactive") {
			return
		}
		level++
	})
	return level
}

func gearLevel(s *goquery.Selection) int {
	switch s.Find(".char-portrait-full-gear-level").Text() {
	case "XII":
		return 12
	case "XI":
		return 11
	case "X":
		return 10
	case "IX":
		return 9
	case "VIII":
		return 8
	case "VII":
		return 7
	case "VI":
		return 6
	case "V":
		return 5
	case "IV":
		return 4
	case "III":
		return 3
	case "II":
		return 2
	case "I":
		return 1
	default:
		return 0
	}
}

type CharacterStats struct {
	Name      string
	Level     int
	GearLevel int
	Stars     int

	// Current character gallactic power
	GalacticPower int

	// List of skils of this character
	Skills []Skill

	// Basic Stats
	STR                int
	AGI                int
	INT                int
	StrenghGrowth      float64
	AgilityGrowth      float64
	IntelligenceGrowth float64

	// General
	Health         int
	Protection     int
	Speed          int
	CriticalDamage float64
	Potency        float64
	Tenacity       float64
	HealthSteal    float64

	PhysicalDamage     int
	PhysicalCritChance float64
	SpecialDamage      int
	SpecialCritChance  float64
}

type Skill struct {
	Name  string
	Level int
}

func (c *Client) CharacterStats(char string) (*CharacterStats, error) {
	charSlug := CharSlug(CharName(char))
	doc, err := c.Get(fmt.Sprintf("https://swgoh.gg/u/%s/collection/%s/", c.profile, charSlug))
	if err != nil {
		return nil, fmt.Errorf("swgohgg: profile %s may not have %s activated. (err=%v)", c.profile, CharName(char), err.Error())
	}

	charStats := &CharacterStats{}
	charStats.Name = doc.Find(".pc-char-overview-name").Text()
	charStats.Level = atoi(doc.Find(".char-portrait-full-level").Text())
	charStats.Stars = int(stars(doc.Find(".player-char-portrait")))
	gearInfo := strings.Split(doc.Find(".pc-gear").First().Find(".pc-heading").First().AttrOr("title", "Gear -1 "), " ")
	if len(gearInfo) > 1 {
		charStats.GearLevel = atoi(gearInfo[1])
	}
	charStats.GalacticPower = atoi(doc.Find(".unit-gp-stat-amount-current").First().Text())
	// Skills
	doc.Find(".pc-skills-list").First().Find(".pc-skill").Each(func(i int, s *goquery.Selection) {
		skill := Skill{}
		skill.Name = s.Find(".pc-skill-name").First().Text()
		skill.Level = skillLevel(s)
		charStats.Skills = append(charStats.Skills, skill)
	})
	//Stats
	doc.Find(".media-body .pc-stat").Each(func(i int, s *goquery.Selection) {
		name, value := s.Find(".pc-stat-label").Text(), s.Find(".pc-stat-value").Text()
		switch strings.TrimSpace(name) {
		case "Strength (STR)":
			charStats.STR = atoi(value)
		case "Agility (AGI)":
			charStats.AGI = atoi(value)
		case "Intelligence (INT)":
			charStats.INT = atoi(value)
		case "Strength Growth":
			charStats.StrenghGrowth = atof(value)
		case "Agility Growth":
			charStats.AgilityGrowth = atof(value)
		case "Intelligence Growth":
			charStats.IntelligenceGrowth = atof(value)
		case "Health":
			charStats.Health = atoi(value)
		case "Protection":
			charStats.Protection = atoi(value)
		case "Speed":
			charStats.Speed = atoi(value)
		case "Critical Damage":
			charStats.CriticalDamage = atof(value)
		case "Potency":
			charStats.Potency = atof(value)
		case "Tenacity":
			charStats.Tenacity = atof(value)
		case "Health Steal":
			charStats.HealthSteal = atof(value)
		case "Physical Damage":
			charStats.PhysicalDamage = atoi(value)
		case "Special Damage":
			charStats.SpecialDamage = atoi(value)
		case "Physical Critical Chance":
			charStats.PhysicalCritChance = atof(value)
		case "Special Critical Chance":
			charStats.SpecialCritChance = atof(value)
		}
	})
	return charStats, nil
}

func skillLevel(s *goquery.Selection) int {
	title := s.Find(".pc-skill-levels").First().AttrOr("data-title", "Level -1")
	// Title is in the form 'Level X of Y'
	fields := strings.Fields(title)
	if len(fields) >= 2 {
		return atoi(fields[1])
	}
	return -1
}

func atof(src string) float64 {
	src = strings.Replace(src, "%", "", -1)
	v, _ := strconv.ParseFloat(src, 64)
	return v
}

// atoi best-effort convertion to int, return 0 if unparseable
func atoi(src string) int {
	src = strings.Replace(src, ",", "", -1)
	src = strings.Replace(src, ".", "", -1)
	src = strings.Replace(src, "%", "", -1)
	v, _ := strconv.ParseInt(src, 10, 32)
	return int(v)
}
