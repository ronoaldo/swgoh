package swgohhelp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

//player.go has info on Player, as well as unit + ranking data.

type Arena struct {
	Char ArenaRanking `json:"char"`
	Ship ArenaRanking `json:"ship"`
}

// ArenaRanking holds player arena ranking.
type ArenaRanking struct {
	Rank  int         `json:"rank"`
	Squad []SquadUnit `json:"squad"`
}

type Char struct {
	Rank  int64   `json:"rank"`
	Squad []Squad `json:"squad"`
}

type Player struct {
	AllyCode   int          `json:"allyCode"`
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Level      int64        `json:"level"`
	Titles     Titles       `json:"titles"`
	GuildRefID string       `json:"guildRefId"`
	GuildName  string       `json:"guildName"`
	Stats      []Stat       `json:"stats"`
	Roster     []UnitRoster `json:"roster"`
	Arena      Arena        `json:"arena"`
	Updated    int64        `json:"updated"`
}

// Relic contains the characte relic stats, such as tier.
type Relic struct {
	CurrentTier int64 `json:"currentTier"`
}

type Skill struct {
	ID      string `json:"id"`
	Tier    int64  `json:"tier"`
	NameKey string `json:"nameKey"`
	IsZeta  bool   `json:"isZeta"`
	Tiers   int64  `json:"tiers"`
}

type SkillReferenceList struct {
	SkillID           string `json:"skillId"`
	RequiredTier      int64  `json:"requiredTier"`
	RequiredRarity    int64  `json:"requiredRarity"`
	RequiredRelicTier int64  `json:"requiredRelicTier"`
}

type Squad struct {
	ID            string `json:"id"`
	DefID         string `json:"defId"`
	SquadUnitType int64  `json:"squadUnitType"`
}

type SquadUnit struct {
	ID     string        `json:"id"`
	UnitID string        `json:"defId"`
	Type   SquadUnitType `json:"squadUnitType"`
}

type Stat struct {
	NameKey string `json:"nameKey"`
	Value   int64  `json:"value"`
	Index   int64  `json:"index"`
}

type Titles struct {
	Selected string   `json:"selected"`
	Unlocked []string `json:"unlocked"`
}

type Unit struct {
	UnitID                 string               `json:"unitId"`
	Slot                   int64                `json:"slot"`
	SkillReferenceList     []SkillReferenceList `json:"skillReferenceList"`
	SkilllessCrewAbilityID string               `json:"skilllessCrewAbilityId"`
	Gp                     int64                `json:"gp"`
	Cp                     float64              `json:"cp"`
}

type UnitEquipment struct {
	EquipmentID string `json:"equipmentId"`
	Slot        int64  `json:"slot"`
	NameKey     string `json:"nameKey"`
}

type UnitRoster struct {
	ID            string `json:"id"`
	DefID         string `json:"defId"`
	Name          string `json:"nameKey"`
	Rarity        int64  `json:"rarity"`
	Level         int64  `json:"level"`
	GalacticPower int64  `json:"gp"`
	XP            int64  `json:"xp"`

	Gear       int64           `json:"gear"`
	Equipped   []UnitEquipment `json:"equipped"`
	CombatType int64           `json:"combatType"`
	Skills     []Skill         `json:"skills"`
	Mods       []Mod           `json:"mods"`
	Crew       []Unit          `json:"crew"`

	PrimaryUnitStat interface{} `json:"primaryUnitStat"`
	Relic           *Relic      `json:"relic"`

	Stats *UnitStats `json:"stats,omitempty"`
	Data  *DataUnit  `json:"data,omitempty"`
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

// Players retrieves several player profile stats and roster details.
func (c *Client) Players(allyCodes ...string) (players []Player, err error) {
	allyCodeNumbers, err := parseAllyCodes(allyCodes...)
	if err != nil {
		return nil, fmt.Errorf("swgohhelp: error parsing ally codes: %v", err)
	}
	// Check if we have some of them in cache first
	players, missingFromCache := c.CheckAllyCodesInCache(allyCodeNumbers)
	if len(missingFromCache) == 0 {
		return players, nil
	}
	fmt.Printf("%d\n", missingFromCache)
	payload, err := json.Marshal(map[string]interface{}{
		"allycodes": missingFromCache,
		"language":  "eng_us",
		"enums":     false,
		"project": map[string]int{
			"id":         1,
			"allyCode":   1,
			"name":       1,
			"level":      1,
			"stats":      1,
			"arena":      1,
			"roster":     1,
			"guildName":  1,
			"guildRefId": 1,
			"titles":     1,
			"updated":    1,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to martial player data payload: %v", err)
	}
	resp, err := c.call("POST", "/swgoh/players", "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("unable to get player data %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("issue reading response body %s\nError: %v", c.endpoint, err)
	}

	if err := json.Unmarshal(body, &players); err != nil {
		return nil, fmt.Errorf("issue when unmarshalling json response of %s\nError: %v", c.endpoint, err)
	}

	// if *useExternalStats {
	// 	// Enrich result with related data from Crinolo's stat API
	// 	url := "https://swgoh-stat-calc.glitch.me/api/?flags=withModCalc,gameStyle,calcGP"
	// 	for _, player := range players {
	// 		if err := player.getExtendedPlayerData(c, url); err != nil {
	// 			return nil, fmt.Errorf("error when getting extended player data: %v", err)
	// 		}
	// 	}
	// }

	// Enrich result with related data from data collections
	// titles, err := c.DataPlayerTitles()
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to get title data %v", err)
	// }
	// for i := range players {
	// 	players[i].Titles.Selected = titles[players[i].Titles.Selected].Name
	// 	for j := range players[i].Titles.Unlocked {
	// 		titleKey := players[i].Titles.Unlocked[j]
	// 		players[i].Titles.Unlocked[j] = titles[titleKey].Name
	// 	}
	// }
	// unitList, err := c.DataUnits()
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to get unit data: %v", err)
	// }
	// for i := range players {
	// 	for j := range players[i].Roster {
	// 		id, defid := players[i].Roster[j].ID, players[i].Roster[j].DefID
	// 		if unitData, ok := unitList[id]; ok {
	// 			players[i].Roster[j].Data = &unitData
	// 		}
	// 		if unitData, ok := unitList[defid]; ok {
	// 			players[i].Roster[j].Data = &unitData
	// 		}
	// 	}
	// }

	// Save players missing from cache
	for i := range players {
		player := players[i]
		c.players.Put(strconv.Itoa(player.AllyCode), &player)
		log.Printf("swgohhelp: saving player %v in cache ...", player.AllyCode)
	}

	return players, nil
}

func (p *Player) getExtendedPlayerData(c *Client, url string) error {
	b, err := json.Marshal(p.Roster)
	if err != nil {
		return fmt.Errorf("unable to marshal extended player data: %v", err)
	}
	if c.debug {
		writeLogFile(b, "req", "POST", "_crinoloapi")
	}
	resp, err := c.hc.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("error when making extended player data request: %v", err)
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read player data response data: %v", err)
	}
	fmt.Printf("response data: %s\n", string(b))
	if c.debug {
		writeLogFile(b, "resp", "POST", "_crinoloapi")
	}
	err = json.Unmarshal(b, &p.Roster)
	if err != nil {
		return fmt.Errorf("unable to unmarshal extended player data resp: %v", err)
	}
	return nil
}
