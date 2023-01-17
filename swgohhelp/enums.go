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
