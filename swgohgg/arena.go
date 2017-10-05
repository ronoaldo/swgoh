package swgohgg

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Arena() (team []*CharacterStats, lastUpdate time.Time, err error) {
	doc, err := c.Get(fmt.Sprintf("https://swgoh.gg/u/%s/", c.profile))
	if err != nil {
		return
	}
	teamName := make([]string, 0, 5)
	doc.Find(".current-rank-team").First().Find(".static-char-portrait").Each(func(i int, s *goquery.Selection) {
		teamName = append(teamName, s.AttrOr("title", "UNKOWN"))
	})
	for _, name := range teamName {
		var stat *CharacterStats
		stat, err = c.CharacterStats(name)
		if err != nil {
			return
		}
		team = append(team, stat)
	}
	timestamp := doc.Find(".user-last-updated .datetime").First().AttrOr("data-datetime", "0000-00-00T00:00:00Z")
	lastUpdate, err = time.Parse(time.RFC3339, timestamp)
	return
}
