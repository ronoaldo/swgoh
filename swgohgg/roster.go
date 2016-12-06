package swgohgg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func (c *Client) Roster(profile string) (roster []*Char, err error) {
	url := fmt.Sprintf("https://swgoh.gg/u/%s/collection/", profile)
	resp, err := c.hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("swgohgg: unexpected status code %d", resp.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".collection-char-list .collection-char").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("collection-char-missing") {
			return
		}
		char := parseChar(s)
		roster = append(roster, char)
	})
	return roster, nil
}

type Char struct {
	Name  string
	Stars int
	Level int
	Gear  int
}

func (c *Char) String() string {
	return fmt.Sprintf("%s %d* G%d", c.Name, c.Stars, c.Gear)
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
		return -1
	}
}
