package swgohhelp

import (
	"strconv"
	"strings"
	"time"
)

// AuthResponse represents the authentication response data.
type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Player represents the full player profile information.
type Player struct {
	Name     string `json:"name"`
	AllyCode int    `json:"allyCode"`
	Level    int    `json:"level"`

	GuildName  string `json:"guildName"`
	GuildRefID string `json:"guildRefId"`

	Titles PlayerTitle `json:"titles"`

	UpdatedAt Timestamp `json:"updated"`

	//Stats  ProfileStats `json:"stats"`
	//Roster []Units      `json:"roster"`
	Arena Arena `json:"arena"`
}

// PlayerTitle is a list of player unlocked and selected titles.
type PlayerTitle struct {
	Slected  string   `json:"selected"`
	Unlocked []string `json:"unlocked"`
}

// PlayerRequest is the payload sent to the /swgoh/player endpoint.
type PlayerRequest struct {
	AllyCodes []string `json:"allycode"`
	Lang      string   `json:"lang"`
	Enums     bool     `json:"enums,omitempty"`

	Project map[string]int `json:"project,omitempty"`
}

// Arena wraps both arena rankings for the player.
type Arena struct {
	Char ArenaRanking `json:"char"`
	Ship ArenaRanking `json:"ship"`
}

// ArenaRanking holds player arena ranking.
type ArenaRanking struct {
	Rank  int         `json:"rank"`
	Squad []SquadUnit `json:"squad"`
}

// SquadUnit represents an arena squad unit identifier set.
type SquadUnit struct {
	ID     string `json:"id"`
	UnitID string `json:"defId"`
	Type   string `json:"squadUnitType"`
}

// Timestamp is a helper unix timestamp JSON marshaller/unmarshaller.
// Source: https://gist.github.com/alexmcroberts/219127816e7a16c7bd70
type Timestamp time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Timestamp) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	return
}

// String implements the Stringer interface.
func (t Timestamp) String() string { return time.Time(t).String() }
