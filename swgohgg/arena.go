package swgohgg

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Arena() (team []*CharacterStats, err error) {
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
	return
}
