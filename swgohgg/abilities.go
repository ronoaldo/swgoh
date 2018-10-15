package swgohgg

import "fmt"

// Ability is a generic description of an ability for a given character.
type Ability struct {
	Name            string
	Character       string
	CharacterBaseID string
	IsZeta          bool
}

func (a Ability) String() string {
	suffix := ""
	if a.IsZeta {
		suffix = " (zeta)"
	}
	return fmt.Sprintf("{%s ability \"%s\"%s}", a.Character, a.Name, suffix)
}

// Zetas fetches the current character abilities available in the
// "/characters-zeta-abilities" website pages.
func (c *Client) Zetas() (zetas []Ability, err error) {
	chars, err := c.gg.Characters()
	if err != nil {
		return nil, err
	}

	abilities, err := c.gg.Abilities()
	if err != nil {
		return nil, err
	}
	for _, a := range abilities {
		if a.IsZeta {
			char := chars.FromBaseID(a.CharacterBaseID)
			if char == nil {
				return nil, fmt.Errorf("swgohgg: unexpected character base ID from zetas: %v", a.CharacterBaseID)
			}
			zetas = append(zetas, Ability{
				Name:            a.Name,
				IsZeta:          true,
				Character:       char.Name,
				CharacterBaseID: a.CharacterBaseID,
			})
		}
	}

	return
}
