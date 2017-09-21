package swgohgg

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Collection() (roster Collection, err error) {
	for page := 1; page <= 5; page++ {
		url := fmt.Sprintf("https://swgoh.gg/u/%s/collection/?page=%d", c.profile, page)
		resp, err := c.hc.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			break
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("swgohgg: unexpected status code %d", resp.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
		doc.Find(".collection-char-list .collection-char").Each(func(i int, s *goquery.Selection) {
			char := parseChar(s)
			if !roster.Contains(char.Name) {
				roster = append(roster, char)
			}
		})
	}
	sort.Sort(ByStars(roster, false))
	return roster, nil
}

type Collection []*Char

func (r Collection) Contains(char string) bool {
	for i := range r {
		if r[i].Name == char {
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
