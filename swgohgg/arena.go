package swgohgg

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Arena returns basic information about the player Arena team.
// If authorized, attempts to fetch more information from character stats
func (c *Client) Arena() (team []*CharacterStats, lastUpdate time.Time, err error) {
	url := fmt.Sprintf("https://swgoh.gg/u/%s/", c.shortLink)
	doc, err := c.Get(url)
	if err != nil {
		return
	}
	doc.Find(".panel-body > p").Each(func(i int, s *goquery.Selection) {
		text := strings.ToLower(s.Text())
		if strings.Contains(text, "ally code") {
			c.SetAllyCode(nonDigits.ReplaceAllString(text, ""))
		}
		if strings.HasPrefix(text, "guild ") && c.guildName == "" {
			c.guildName = strings.TrimSpace(s.Text()[6:])
		}
	})
	doc.Find(".panel-profile .panel-body h5.panel-title").Each(func(i int, s *goquery.Selection) {
		c.playerName = strings.TrimSpace(s.Text())
	})
	order := make([]string, 0, 5)
	basicStats := make(map[string]CharacterStats)
	doc.Find(".current-rank-team").First().Find(".player-char-portrait").Each(func(i int, s *goquery.Selection) {
		charName := s.AttrOr("title", "UNKOWN")
		charBasicStats := CharacterStats{
			Name:  charName,
			Level: atoi(s.Find(".char-portrait-full-level").Text()),
			Stars: stars(s),
		}
		basicStats[charName] = charBasicStats
		order = append(order, charName)
	})
	for _, name := range order {
		basic := basicStats[name]
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
	}
	timestamp := doc.Find(".user-last-updated .datetime").First().AttrOr("data-datetime", "0000-00-00T00:00:00Z")
	lastUpdate, err = time.Parse(time.RFC3339, timestamp)
	return
}
