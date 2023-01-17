package swgohhelp

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
