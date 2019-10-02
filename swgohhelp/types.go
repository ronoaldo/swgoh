package swgohhelp

import (
	"bytes"
	"fmt"
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

	Stats  []PlayerStat `json:"stats"`
	Roster Roster       `json:"roster"`

	Arena Arena `json:"arena"`

	UpdatedAt Timestamp `json:"updated"`
}

// Raid contains information about the last performed raid for a guild.
type Raid struct {
	Rancor   string `json:"rancor"`
	AAT      string `json:"aat"`
	SithRaid string `json:"sith_raid"`
}

// Guild represents the full guild profile with abbreviated roster info.
type Guild struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Members  int    `json:"members"`
	Status   int    `json:"status"`
	Required int    `json:"required"`

	BannerColor string `json:"bannerColor"`
	BannerLogo  string `json:"bannerLogo"`

	Message string `json:"message"`
	GP      int    `json:"gp"`

	Raid Raid `json:"raid"`

	Roster []Player `json:"roster"`

	UpdatedAt Timestamp `json:"updated"`
}

// PlayerTitle is a list of player unlocked and selected titles.
type PlayerTitle struct {
	Selected string   `json:"selected"`
	Unlocked []string `json:"unlocked"`
}

// PlayerStat is a single player profile statistic info, like how many battles won.
type PlayerStat struct {
	Name  string `json:"nameKey"`
	Value int64  `json:"value"`
	Index int64  `json:"index"`
}

// Roster is a helper collection to manipulate player roster
type Roster []Unit

// FindByID filter the unit collection by the DefID attribute
func (r Roster) FindByID(defID string) (*Unit, bool) {
	for _, unit := range r {
		if unit.DefID == defID {
			return &unit, true
		}
	}
	return nil, false
}

// FindByName filter the unit collection by name attribute
func (r Roster) FindByName(unitName string) (*Unit, bool) {
	for _, unit := range r {
		if strings.ToLower(unit.Name) == strings.ToLower(unitName) {
			return &unit, true
		}
	}
	return nil, false
}

// Mods returns the roster equiped mods
func (r Roster) Mods() (mods []Mod) {
	for i := range r {
		for j := range r[i].Mods {
			if r[i].Mods[j].UnitEquiped == "" {
				r[i].Mods[j].UnitEquiped = r[i].Name
			}
			mods = append(mods, r[i].Mods[j])
		}
	}
	return
}

// Unit is a game unit entity, character or ship.
type Unit struct {
	ID            string          `json:"id"`
	DefID         string          `json:"defId"`
	Name          string          `json:"nameKey"`
	Rarity        int             `json:"rarity"`
	Level         int             `json:"level"`
	XP            int             `json:"xp"`
	GalacticPower int             `json:"gp"`
	Gear          int             `json:"gear"`
	Equiped       []UnitEquipment `json:"equipped"`

	CombatType CombatType `json:"combatType"`

	Skills []UnitSkill `json:"skills"`
	Mods   []Mod       `json:"mods"`
	Crew   []Unit      `json:"crew"`
	Relic  Relic       `json:"relic"`

	Stats *UnitStats `json:"stats,omitempty"`
}

// Relic contains the characte relic stats, such as tier.
type Relic struct {
	Tier int `json:"currentTier"`
}

// UnitStats unit statis information split by Final and FromMods
type UnitStats struct {
	Final    UnitStatItems `json:"final"`
	FromMods UnitStatItems `json:"mods"`
}

// UnitStatItems is a set of character statistics such as health, speed, etc.
type UnitStatItems struct {
	// Primary attributes
	Strength int `json:"Strength"`
	Agility  int `json:"Agility"`
	Tactics  int `json:"Tactics"`

	// General
	Health         int     `json:"Health"`
	Protection     int     `json:"Protection"`
	Speed          int     `json:"Speed"`
	CriticalDamage float64 `json:"Critical Damage"`
	Potency        float64 `json:"Potency"`
	Tenacity       float64 `json:"Tenacity"`
	HealthSteal    float64 `json:"Helth Steal"`

	// Physical Offense
	PhysicalDamage         int     `json:"Physical Damage"`
	PhysicalCriticalChance float64 `json:"Physical Critical Chance"`
	ArmorPenetration       int     `json:"Armor Penetration"`
	PhysicalAccuracy       float64 `json:"Physical Accuracy"`

	// Physical Survivability
	Armor                     float64 `json:"Armor"`
	DodgeChance               float64 `json:"Dodge Chance"`
	PhysicalCriticalAvoidance float64 `json:"Physical Critical Avoidance"`

	// Special Offense
	SpecialDamage         int     `json:"Special Damage"`
	SpecialCriticalChance float64 `json:"Special Critical Chance"`
	ResistancePenetration int     `json:"Resistance Penetration"`
	SpecialAccuracy       float64 `json:"Special Accuracy"`

	// Special Survivability
	Resistance               float64 `json:"Resistance"`
	DeflectionChance         float64 `json:"Deflection Chance"`
	SpecialCriticalAvoidance float64 `json:"Special Critical Avoidance"`
}

// UnitSkill is a single unit skill, with level and value
type UnitSkill struct {
	ID     string `json:"id"`
	Tier   int    `json:"tier"`
	Name   string `json:"nameKey"`
	IsZeta bool   `json:"isZeta"`
}

// Mod is a character mod detailed value.
type Mod struct {
	ID          string    `json:"id"`
	Level       int       `json:"level"`
	Set         ModSet    `json:"set"`
	Tier        int       `json:"tier"`
	Pips        int       `json:"pips"`
	Slot        ModSlot   `json:"slot"`
	UnitEquiped string    `json:"unit_equiped"`
	Primary     ModStat   `json:"primaryStat"`
	Secondaries []ModStat `json:"secondaryStat"`
}

func (m Mod) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%d* Lv%d %s %s '%s' ", m.Pips, m.Level, m.Set, m.Slot, m.UnitEquiped)
	fmt.Fprintf(&buff, "%s ", m.Primary)
	for _, s := range m.Secondaries {
		fmt.Fprintf(&buff, "%s", s)
	}
	return buff.String()
}

// ModStat is a single mod stat name and value.
// Unit is the key name of the modified character stat attribute,
// and value is always a floating point even when precision is zero.
type ModStat struct {
	Unit  ModUnitStat `json:"unitStat"`
	Value float64     `json:"value"`
	Roll  int         `json:"roll"`
}

func (s ModStat) String() string {
	return fmt.Sprintf("%.02f %s (%d)", s.Value, s.Unit, s.Roll)
}

// UnitEquipment is the unit equiped gear at the current level
type UnitEquipment struct {
	EquipmentID string `json:"equipmentId"`
	Slot        int    `json:"slot"`
	NameKey     string `json:"nameKey"`
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
	ID     string        `json:"id"`
	UnitID string        `json:"defId"`
	Type   SquadUnitType `json:"squadUnitType"`
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
