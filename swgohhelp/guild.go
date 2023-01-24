package swgohhelp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Guild struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Desc        string   `json:"desc"`
	Members     int64    `json:"members"`
	Status      int64    `json:"status"`
	Required    int64    `json:"required"`
	BannerColor string   `json:"bannerColor"`
	BannerLogo  string   `json:"bannerLogo"`
	Message     string   `json:"message"`
	Gp          int64    `json:"gp"`
	Raid        Raid     `json:"raid"`
	Roster      []Roster `json:"roster"`
	Updated     int64    `json:"updated"`
}

type Raid struct {
	Rancor          string `json:"rancor"`
	Aat             string `json:"aat"`
	SithRAID        string `json:"sith_raid"`
	RancorChallenge string `json:"rancor_challenge"`
}

type Roster struct {
	ID               string `json:"id"`
	GuildMemberLevel int64  `json:"guildMemberLevel"`
	Name             string `json:"name"`
	Level            int64  `json:"level"`
	AllyCode         int64  `json:"allyCode"`
	Gp               int64  `json:"gp"`
	GpChar           int64  `json:"gpChar"`
	GpShip           int64  `json:"gpShip"`
	Updated          int64  `json:"updated"`
}

// Players retrieves all the players from a specified guild including roster details.
func (c *Client) Guild(allyCode string) (*Guild, error) {
	var guildResponse = []Guild{}
	allyCodeNumbers, err := parseAllyCodes(allyCode)
	if err != nil {
		return nil, fmt.Errorf("swgohhelp: error parsing ally code: %v", err)
	}
	//Should only recieve one
	allyCodeNumber := allyCodeNumbers[0]
	// Check if we have the player's guild in the cache first.
	// var g Guild
	// if ok := c.guilds.Get(strconv.Itoa(allyCodeNumber), &g); ok {
	// 	guild = &g
	// 	return guild, nil
	// }
	body, err := c.getGuildData(allyCodeNumber)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &guildResponse); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response from %s\nError: %v", c.endpoint, err)
	}

	if len(guildResponse) == 0 {
		return nil, fmt.Errorf("no guild retrieved for ally-code %s", allyCode)
	}

	// Save the guild (indexed by the player we know about) for future use.
	c.guilds.Put(strconv.Itoa(allyCodeNumber), guildResponse[0])
	log.Printf("swgohhelp: saving guild for player %v in cache ...", allyCodeNumber)

	return &guildResponse[0], nil
}

func (c *Client) getGuildData(allyCode int) ([]byte, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"allycode": allyCode,
		"language": "eng_us",
		"enums":    false,
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.call("POST", "/swgoh/guild", "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
