package swgohgg

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Arena returns basic information about the player Arena team.
// If authorized, attempts to fetch more information from character stats
func (c *Client) Arena() (team []*CharacterStats, lastUpdate time.Time, err error) {
	doc, err := c.Get(fmt.Sprintf("https://swgoh.gg/u/%s/", c.profile))
	if err != nil {
		return
	}
	basicStats := make(map[string]CharacterStats)
	doc.Find(".current-rank-team").First().Find(".static-char-portrait").Each(func(i int, s *goquery.Selection) {
		charName := s.AttrOr("title", "UNKOWN")
		charBasicStats := CharacterStats{
			Name:  charName,
			Level: atoi(s.Find(".char-portrait-full-level").Text()),
			Stars: stars(s),
		}
		basicStats[charName] = charBasicStats
	})
	fmt.Printf("basicStats => %#v", basicStats)
	for name := range basicStats {
		basic := basicStats[name]
		if c.authorized {
			var stat *CharacterStats
			stat, err = c.CharacterStats(name)
			if err != nil {
				return
			}
			if stat.GearLevel < 0 {
				stat.Name = basic.Name
				stat.Level = basic.GearLevel
				stat.Stars = basic.Stars
			}
			team = append(team, stat)
		} else {
			team = append(team, &basic)
		}
	}
	timestamp := doc.Find(".user-last-updated .datetime").First().AttrOr("data-datetime", "0000-00-00T00:00:00Z")
	lastUpdate, err = time.Parse(time.RFC3339, timestamp)
	return
}
