package swgohgg

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Ships() (ships Ships, err error) {
	url := fmt.Sprintf("https://swgoh.gg/u/%s/ships/", c.profile)
	doc, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	doc.Find(".collection-char-list .collection-ship").Each(func(i int, s *goquery.Selection) {
		ship := parseShip(s)
		if !ships.Contains(ship.Name) {
			ships = append(ships, ship)
		}
	})
	sort.Sort(ByShipStars(ships, false))
	return ships, nil
}

type Ships []*Ship

func (r Ships) Contains(ship string) bool {
	for i := range r {
		if r[i].Name == ship {
			return true
		}
	}
	return false
}

func (r Ships) ContainsAll(ships ...string) bool {
	for _, ship := range ships {
		if !r.Contains(ship) {
			return false
		}
	}
	return true
}

type Ship struct {
	Name  string
	Stars int
	Level int
}

func (c *Ship) String() string {
	if c == nil {
		return "nil"
	}
	return fmt.Sprintf("%s %d* Lvl%d", c.Name, c.Stars, c.Level)
}

func parseShip(s *goquery.Selection) *Ship {
	var ship Ship
	ship.Name = s.Find(".collection-ship-name-link").Text()

	ship.Level, _ = strconv.Atoi(s.Find(".ship-portrait-full-frame-level").Text())
	ship.Stars = shipStars(s)
	return &ship
}

func shipStars(s *goquery.Selection) int {
	level := 0
	s.Find(".ship-portrait-full-star").Each(func(i int, star *goquery.Selection) {
		if star.HasClass("ship-portrait-full-star-inactive") {
			return
		}
		level++
	})
	return level
}
