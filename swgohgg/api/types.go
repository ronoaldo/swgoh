package api

import (
	"bytes"
	"fmt"
)

// Player represents the player profile data from the website and game.
type Player struct {
	PlayerProfile `json:"data"`
	Units         []Unit `json:"units"`
}

// PlayerProfile details the player profile information.
type PlayerProfile struct {
	Name  string `json:"name"`
	Level int    `json:"level"`

	ArenaRank         int    `json:"arena_rank"`
	ArenaLeaderBaseID string `json:"arena_leader_base_id"`

	AllyCode    int `json:"ally_code"`
	GP          int `json:"galactic_power"`
	ShipGP      int `json:"ship_galactic_power"`
	CharacterGP int `json:"character_galactic_power"`
}

// Unit is a player unit detailed info.
type Unit struct {
	UnitData `json:"data"`
}

// UnitData holds character data for each player unit.
type UnitData struct {
	BaseID    string `json:"base_id"`
	Power     int    `json:"power"`
	Level     int    `json:"level"`
	GearLevel int    `json:"gear_level"`
	Rarity    int    `json:"rarity"`
	URL       string `json:"url"`

	ZetaAbilities []string  `json:"zeta_abilities"`
	Stats         UnitStats `json:"stats"`
	Gear          []EquipedGear
}

// UnitStats contains a map of unit statistics
type UnitStats map[StatName]float64

func (u UnitStats) String() string {
	if u == nil {
		return "map[StatName]float64(nil)"
	}
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "UnitStats[")
	fmt.Fprintf(&buff, "Strength:%.0f", u[Strength])
	fmt.Fprintf(&buff, ", Agility:%.0f", u[Agility])
	fmt.Fprintf(&buff, ", Tatics:%.0f", u[Tatics])
	fmt.Fprintf(&buff, ", Speed:%.0f", u[Speed])
	fmt.Fprintf(&buff, ", Potency:%.02f%%", u[Potency]*100)
	fmt.Fprintf(&buff, ", Tenacity:%.02f%%", u[Tenaticy]*100)
	fmt.Fprintf(&buff, ", PhysicalAccuracy:%.02f", u[PhysicalAccuracy])
	fmt.Fprintf(&buff, ", SpecialAccuracy:%.02f", u[SpecialAccuracy])
	fmt.Fprintf(&buff, "]")
	return buff.String()
}

// StatName is a helper Enum to recover data from UnitStats
type StatName string

// Enumerated list of stats
const (
	Strength StatName = "2"
	Agility  StatName = "3"
	Tatics   StatName = "4"

	Health         StatName = "1"
	Protection     StatName = "28"
	Speed          StatName = "5"
	CriticalDamage StatName = "16"
	Potency        StatName = "17"
	Tenaticy       StatName = "18"
	HealthSteal    StatName = "27"

	PhysicalDamage         StatName = "6"
	PhysicalCriticalChange StatName = "14"
	ArmorPenetration       StatName = "10"
	PhysicalAccuracy       StatName = "37"

	Armor                     StatName = "8"
	DodgeChance               StatName = "12"
	PhysicalCriticalAvoidance StatName = "39"

	SpecialDamage         StatName = "7"
	SpecialCriticalChance StatName = "15"
	ResistancePenetration StatName = "11"
	SpecialAccuracy       StatName = "38"

	Resistance               StatName = "9"
	DeflectionChance         StatName = "13"
	SpecialCriticalAvoidance StatName = "40"
)

// EquipedGear represents unit equped gear details.
type EquipedGear struct {
	Slot       int    `json:"slot"`
	IsObtained bool   `json:"is_obtained"`
	BaseID     string `json:"base_id"`
}

// Character represents a character unit in the game.
type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	BaseID string `json:"base_id"`
	PK     int    `json:"pk"`

	URL   string `json:"url"`
	Image string `json:"image"`

	Power      int      `json:"power"`
	CombatType int      `json:"combat_type"`
	Alignment  string   `json:"alignment"`
	Categories []string `json:"categories"`
	Role       string   `json:"role"`

	AbilityClasses []string    `json:"ability_classes"`
	GearLevels     []GearLevel `json:"gear_levels"`
}

// GearLevel represents a character gear level item list.
type GearLevel struct {
	Tier    int      `json:"tier"`
	GearIDs []string `json:"gear"`
}
