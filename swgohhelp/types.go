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
	Name     string `json:"name,omitempty"`
	AllyCode int    `json:"allyCode,omitempty"`
	Level    int    `json:"level,omitempty"`

	GuildName  string `json:"guildName,omitempty"`
	GuildRefID string `json:"guildRefId,omitempty"`

	Titles PlayerTitle `json:"titles,omitempty"`

	Stats  []PlayerStat `json:"stats,omitempty"`
	Roster []Unit       `json:"roster,omitempty"`

	Arena Arena `json:"arena,omitempty"`

	UpdatedAt Timestamp `json:"updated,omitempty"`
}

// PlayerTitle is a list of player unlocked and selected titles.
type PlayerTitle struct {
	Selected string   `json:"selected,omitempty"`
	Unlocked []string `json:"unlocked,omitempty"`
}

// PlayerStat is a single player profile statistic info, like how many battles won.
type PlayerStat struct {
	Name  string `json:"nameKey,omitempty"`
	Value int64  `json:"value,omitempty"`
	Index int64  `json:"index,omitempty"`
}

// Unit is a game unit entity, character or ship.
type Unit struct {
	ID     string `json:"id,omitempty"`
	DefID  string `json:"defId,omitempty"`
	Name   string `json:"nameKey,omitempty"`
	Rarity int    `json:"rarity,omitempty"`
	Level  int    `json:"level,omitempty"`
	XP     int    `json:"xp,omitempty"`
	Gear   int    `json:"gear,omitempty"`

	// Either SHIP or CHARACTER
	CombatType string `json:"combatType,omitempty"`

	Skills []UnitSkill `json:"skills,omitempty"`
	Mods   []Mod       `json:"mods,omitempty"`
	Crew   []Unit      `json:"crew,omitempty"`
}

// UnitSkill is a single unit skill, with level and value
type UnitSkill struct {
	ID     string `json:"id,omitempty"`
	Tier   int    `json:"tier,omitempty"`
	Name   string `json:"nameKey,omitempty"`
	IsZeta bool   `json:"isZeta,omitempty"`
}

// Mod is a character mod detailed value.
type Mod struct {
	ID          string    `json:"id,omitempty"`
	Level       int       `json:"level,omitempty"`
	Tier        int       `json:"tier,omitempty"`
	Slot        int       `json:"slot,omitempty"`
	Pips        int       `json:"pips,omitempty"`
	Primary     ModStat   `json:"primaryStat,omitempty"`
	Secondaries []ModStat `json:"secondaryStat,omitempty"`
}

// ModStat is a single mod stat name and value.
// Unit is the key name of the modified character stat attribute,
// and value is always a floating point even when precision is zero.
type ModStat struct {
	Unit  string  `json:"unitStat,omitempty"`
	Value float64 `json:"value,omitempty"`
	Roll  int     `json:"roll,omitempty"`
}

// Arena wraps both arena rankings for the player.
type Arena struct {
	Char ArenaRanking `json:"char,omitempty"`
	Ship ArenaRanking `json:"ship,omitempty"`
}

// ArenaRanking holds player arena ranking.
type ArenaRanking struct {
	Rank  int         `json:"rank,omitempty"`
	Squad []SquadUnit `json:"squad,omitempty"`
}

// SquadUnit represents an arena squad unit identifier set.
type SquadUnit struct {
	ID     string `json:"id,omitempty"`
	UnitID string `json:"defId,omitempty"`
	Type   string `json:"squadUnitType,omitempty"`
}

// DataPlayerTitle is the data library information about player titles.
type DataPlayerTitle struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"nameKey,omitempty"`
	Desc    string `json:"descKey,omitempty"`
	Details string `json:"shortDescKey,omitempty"`
}

// Timestamp is a helper unix timestamp JSON marshaller/unmarshaller.
// Source: https://gist.github.com/alexmcroberts/219127816e7a16c7bd70
type Timestamp time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix()*1000, 10)), nil
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
