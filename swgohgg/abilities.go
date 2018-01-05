package swgohgg

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Ability struct {
	Name      string
	Character string
	IsZeta    bool
}

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
			} else {
				log.Printf("> Skipped (len=%d)", len(split))
			}
		})
		if len(aux) == 0 {
			break
		}
		zetas = append(zetas, aux...)
	}
	return
}
