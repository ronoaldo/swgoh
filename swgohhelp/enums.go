package swgohhelp

import "strconv"

// CombatType is the enum value for an Unit combat type
type CombatType int

const (
	// CombatTypeChar represents character units
	CombatTypeChar CombatType = 1
	// CombatTypeShip represents shipts units
	CombatTypeShip CombatType = 2
)

// SquadUnitType is the enum value of an arena unit type
type SquadUnitType int

const (
	// SquadUnitNormal normal arena unit
	SquadUnitNormal SquadUnitType = 1
	// SquadUnitLeader leader arena unit
	SquadUnitLeader SquadUnitType = 2
	// SquadUnitCapitalShip captial ship arena unit
	SquadUnitCapitalShip SquadUnitType = 3
	// SquadUnitReinforcement reinforcement ship arena unit
	SquadUnitReinforcement SquadUnitType = 5
)

func (s SquadUnitType) String() string {
	switch s {
	case SquadUnitNormal:
		return ""
	case SquadUnitLeader:
		return "Leader"
	case SquadUnitCapitalShip:
		return "Capital Ship"
	case SquadUnitReinforcement:
		return "Reinforcement"
	default:
		return "Unknown unit type " + strconv.Itoa(int(s))
	}
}

// MarshalText implemnts the encoding.TextMarshaler
func (s SquadUnitType) MarshalText() string {
	return s.String()
}

// ModSet is the enum for possible mod bonus set
type ModSet int

// Possible bonus set values
const (
	ModSetHealth     ModSet = 1
	ModSetOffense    ModSet = 2
	ModSetDefense    ModSet = 3
	ModSetSpeed      ModSet = 4
	ModSetCritChance ModSet = 5
	ModSetCritDamage ModSet = 6
	ModSetPotency    ModSet = 7
	ModSetTenacity   ModSet = 8
)

func (m ModSet) String() string {
	switch m {
	case ModSetHealth:
		return "Health"
	case ModSetOffense:
		return "Offense"
	case ModSetDefense:
		return "Defense"
	case ModSetSpeed:
		return "Speed"
	case ModSetCritChance:
		return "Critical Chance"
	case ModSetCritDamage:
		return "Critical Damage"
	case ModSetPotency:
		return "Potency"
	case ModSetTenacity:
		return "Tenacity"
	default:
		return "Set:" + strconv.Itoa(int(m))
	}
}

// MarshalText implements encoding.TextMarshaler
func (m ModSet) MarshalText() string {
	return m.String()
}

// MarshalYAML implements yaml.Marshaler
func (m ModSet) MarshalYAML() (interface{}, error) {
	return m.String(), nil
}

// ModUnitStat is the enum with mod unit stat values
type ModUnitStat int

// Possible values for unit stats
const (
	StatHealth                                ModUnitStat = 1
	StatStrength                              ModUnitStat = 2
	StatAgility                               ModUnitStat = 3
	StatIntelligence                          ModUnitStat = 4
	StatSpeed                                 ModUnitStat = 5
	StatAttackDamage                          ModUnitStat = 6
	StatAbilityPower                          ModUnitStat = 7
	StatArmor                                 ModUnitStat = 8
	StatSuppression                           ModUnitStat = 9
	StatArmorPenetration                      ModUnitStat = 10
	StatSuppressionPenetration                ModUnitStat = 11
	StatDodgeRating                           ModUnitStat = 12
	StatDeflectionRating                      ModUnitStat = 13
	StatAttackCriticalRating                  ModUnitStat = 14
	StatAbilityCriticalRating                 ModUnitStat = 15
	StatCriticalDamage                        ModUnitStat = 16
	StatAccuracy                              ModUnitStat = 17
	StatResistance                            ModUnitStat = 18
	StatDodgePercentAdditive                  ModUnitStat = 19
	StatDeflectionPercentAdditive             ModUnitStat = 20
	StatAttackCriticalPercentAdditive         ModUnitStat = 21
	StatAbilityCriticalPercentAdditive        ModUnitStat = 22
	StatArmorPercentAdditive                  ModUnitStat = 23
	StatSuppressionPercentAdditive            ModUnitStat = 24
	StatArmorPenetrationPercentAdditive       ModUnitStat = 25
	StatSuppressionPenetrationPercentAdditive ModUnitStat = 26
	StatHealthSteal                           ModUnitStat = 27
	StatMaxShield                             ModUnitStat = 28
	StatShieldPenetration                     ModUnitStat = 29
	StatHealthRegen                           ModUnitStat = 30
	StatAttackDamagePercentAdditive           ModUnitStat = 31
	StatAbilityPowerPercentAdditive           ModUnitStat = 32
	StatDodgeNegatePercentAdditive            ModUnitStat = 33
	StatDeflectionNegatePercentAdditive       ModUnitStat = 34
	StatAttackCriticalNegatePercentAdditive   ModUnitStat = 35
	StatAbilityCriticalNegatePercentAdditive  ModUnitStat = 36
	StatDodgeNegateRating                     ModUnitStat = 37
	StatDeflectionNegateRating                ModUnitStat = 38
	StatAttackCriticalNegateRating            ModUnitStat = 39
	StatAbilityCriticalNegateRating           ModUnitStat = 40
	StatOffense                               ModUnitStat = 41
	StatDefense                               ModUnitStat = 42
	StatDefensePenetration                    ModUnitStat = 43
	StatEvasionRating                         ModUnitStat = 44
	StatCriticalRating                        ModUnitStat = 45
	StatEvasionNegateRating                   ModUnitStat = 46
	StatCriticalNegateRating                  ModUnitStat = 47
	StatOffensePercentAdditive                ModUnitStat = 48
	StatDefensePercentAdditive                ModUnitStat = 49
	StatDefensePenetrationPercentAdditive     ModUnitStat = 50
	StatEvasionPercentAdditive                ModUnitStat = 51
	StatEvasionNegatePercentAdditive          ModUnitStat = 52
	StatCriticalChancePercentAdditive         ModUnitStat = 53
	StatCriticalNegateChancePercentAdditive   ModUnitStat = 54
	StatMaxHealthPercentAdditive              ModUnitStat = 55
	StatMaxShieldPercentAdditive              ModUnitStat = 56
	StatSpeedPercentAdditive                  ModUnitStat = 57
	StatCounterAttackRating                   ModUnitStat = 58
	StatTaunt                                 ModUnitStat = 59
)

func (s ModUnitStat) String() string {
	switch s {
	case StatHealth:
		return "Health"
	case StatMaxHealthPercentAdditive:
		return "% Health"
	case StatMaxShield:
		return "Protection"
	case StatMaxShieldPercentAdditive:
		return "% Protection"
	case StatOffense:
		return "Offense"
	case StatOffensePercentAdditive:
		return "% Offense"
	case StatDefense:
		return "Defense"
	case StatDefensePercentAdditive:
		return "% Defense"
	case StatCriticalDamage:
		return "% Crit Damage"
	case StatCriticalChancePercentAdditive:
		return "% Crit Chance"
	case StatSpeed:
		return "Speed"
	case StatSpeedPercentAdditive:
		return "% Speed"
	case StatAccuracy:
		return "% Potency"
	case StatResistance:
		return "% Tenacity"
	case StatEvasionNegatePercentAdditive:
		return "% Accuracy"
	case StatCriticalNegateChancePercentAdditive:
		return "% Crit Avoidance"
	default:
		return "Stat:" + strconv.Itoa(int(s))
	}
}

// MarshalText implements encoding.TextMarshaler
func (s ModUnitStat) MarshalText() string {
	return s.String()
}

// MarshalYAML implements yaml.Marshaler
func (s ModUnitStat) MarshalYAML() (interface{}, error) {
	return s.String(), nil
}

// ModSlot is the mod slot enum
type ModSlot int

// Available mod slots
const (
	ModSlotSquare   ModSlot = 1
	ModSlotArrow    ModSlot = 2
	ModSlotDiamond  ModSlot = 3
	ModSlotTriangle ModSlot = 4
	ModSlotCircle   ModSlot = 5
	ModSlotCross    ModSlot = 6
)

func (s ModSlot) String() string {
	switch s {
	case ModSlotSquare:
		return "Square"
	case ModSlotArrow:
		return "Arrow"
	case ModSlotDiamond:
		return "Diamond"
	case ModSlotTriangle:
		return "Triangle"
	case ModSlotCircle:
		return "Circle"
	case ModSlotCross:
		return "Cross"
	default:
		return "Slot:" + strconv.Itoa(int(s))
	}
}

// MarshalText implements encoding.TextMarshaler
func (s ModSlot) MarshalText() string {
	return s.String()
}

// MarshalYAML implements yaml.Marshaler
func (s ModSlot) MarshalYAML() (interface{}, error) {
	return s.String(), nil
}
