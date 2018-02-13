package swgohgg

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Ability is a generic description of an ability for a given character.
type Ability struct {
	Name      string
	Character string
	IsZeta    bool
}

// Zetas fetches the current character abilities available in the
// "/characters-zeta-abilities" website pages.
func (c *Client) Zetas() (zetas []Ability, err error) {
	for page := 1; page <= 5; page++ {
		var (
			doc *goquery.Document
			url = fmt.Sprintf("https://swgoh.gg/characters/zeta-abilities/?page=%d", page)
		)
		doc, err = c.Get(url)
		if err != nil {
			return
		}
		aux := make([]Ability, 0)
		doc.Find(".media-list-stream .character h5").Each(func(i int, s *goquery.Selection) {
			// Commander Luke Skywalker · Rebel Maneuvers
			split := strings.Split(s.Text(), " · ")
			if len(split) >= 2 {
				aux = append(aux, Ability{
					Character: strings.TrimSpace(split[0]),
					Name:      strings.TrimSpace(split[1]),
					IsZeta:    true,
				})
			}
		})
		if len(aux) == 0 {
			break
		}
		zetas = append(zetas, aux...)
	}
	return
}
