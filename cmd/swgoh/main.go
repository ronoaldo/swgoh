package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"text/template"

	"github.com/ronoaldo/swgoh"
	"github.com/ronoaldo/swgoh/swgohhelp"
)

var (
	// Auth
	allyCode string
	username string
	password string

	// Filter
	starLevel  int
	unitFilter string

	// Actions
	showCharacters bool
	showShips      bool
	showMods       bool
	showStats      bool
	showArena      bool

	// Debug
	debug bool
)

func init() {
	flag.StringVar(&allyCode, "a", "", "The user `ally code` in game")
	flag.StringVar(&username, "u", "", "The `username` to authenticate")
	flag.StringVar(&password, "p", "", "The `password` to authenticate")

	// Operation flags
	flag.BoolVar(&showCharacters, "characters", false, "Show user character collection")
	flag.BoolVar(&showShips, "ships", false, "Show user ships collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection as a CSV file to standard output")
	flag.BoolVar(&showStats, "stats", false, "Show a single character stats (requires -char)")
	flag.BoolVar(&showArena, "arena", false, "Show stats for your current arena team")

	// Debug info
	flag.BoolVar(&debug, "debug", false, "Debug request and response to temporary folder")

	// Filter flags
	flag.IntVar(&starLevel, "stars", 0, "The minimal character or mod `stars` to display")
	flag.StringVar(&unitFilter, "unit", "", "Restrict mods used by this `character` or `ship`")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Authenticate
	client := swgohhelp.New(ctx).SetDebug(debug)
	if _, err := client.SignIn(username, password); err != nil {
		log.Fatalf("swgoh: error authenticating with API backend: %v", err)
	}

	players, err := client.Players(allyCode)
	if err != nil {
		log.Fatalf("swgoh: error loading player profile: %v", err)
	}
	if len(players) == 0 {
		log.Fatalf("swgoh: no players found for ally code %v", allyCode)
	}

	player := players[0]
	unitFilter = swgoh.CharName(unitFilter)

	if showStats {
		for _, unit := range player.Roster {
			if unit.Name == unitFilter {
				s := unit.Stats.Final
				m := unit.Stats.FromMods

				fn := template.FuncMap{
					"perc": func(val float64) string {
						return fmt.Sprintf("%.2f", val*100)
					},
				}
				t := template.Must(template.New("unitTemplate").Funcs(fn).Parse(unitTemplate))
				t.Execute(os.Stdout, map[string]interface{}{
					"s":    s,
					"m":    m,
					"unit": unit,
				})
				break
			}
		}
	}

	if showCharacters {
		for _, u := range player.Roster {
			if u.Rarity >= starLevel && u.CombatType == swgohhelp.CombatTypeChar {
				fmt.Println(fmt.Sprintf("%s %d* G%d Lvl%d", u.Name, u.Rarity, u.Gear, u.Level))
			}
		}
	}

	if showShips {
		for _, u := range player.Roster {
			if u.Rarity >= starLevel && u.CombatType == swgohhelp.CombatTypeShip {
				fmt.Println(fmt.Sprintf("%s %d* G%d Lvl%d", u.Name, u.Rarity, u.Gear, u.Level))
			}
		}
	}

	if showMods {
		mods := player.Roster.Mods()
		w := csv.NewWriter(os.Stdout)
		w.Write([]string{"ID", "Pips", "Level", "Set", "Slot", "Character",
			"PrimStatVal", "PrimStatName", "SecStatVal1", "SecStatName1", "SecStatVal2", "SecStatName2",
			"SecStatVal3", "SecStatName3", "SecStatVal4", "SecStatName4"})
		for _, m := range mods {
			row := []string{
				m.ID,
				strconv.Itoa(m.Pips),
				strconv.Itoa(m.Level),
				m.Set.String(),
				m.Slot.String(),
				m.UnitEquiped,
				fmt.Sprintf("%.02f", m.Primary.Value),
				m.Primary.Unit.String(),
			}
			for _, stat := range m.Secondaries {
				row = append(row, fmt.Sprintf("%.02f", stat.Value), stat.Unit.String())
			}
			w.Write(row)
			w.Flush()
		}
	}

	if showArena {
		fmt.Printf("%s's Arena Teams (%s)\n", player.Name, player.Titles.Selected)
		fmt.Printf("\nCharacter Arena (Ranking %d)\n\n", player.Arena.Char.Rank)
		for _, unit := range player.Arena.Char.Squad {
			suffix := ""
			if unit.Type.String() != "" {
				suffix = "(" + unit.Type.String() + ")"
			}
			if char, ok := player.Roster.FindByID(unit.UnitID); ok {
				fmt.Printf("- %d* %s G%d Lvl%d %v\n", char.Rarity, char.Name, char.Gear, char.Level, suffix)
			}
		}
		fmt.Printf("\nShip Arena (Ranking %d)\n\n", player.Arena.Ship.Rank)
		for _, unit := range player.Arena.Ship.Squad {
			if ship, ok := player.Roster.FindByID(unit.UnitID); ok {
				suffix := ""
				if unit.Type.String() != "" {
					suffix = "(" + unit.Type.String() + ")"
				}
				fmt.Printf("- %d* %s Lvl%d %v\n", ship.Rarity, ship.Name, ship.Level, suffix)
			}
		}
	}
}

var cleanAllyRegexp = regexp.MustCompile("[^0-9]+")

var unitTemplate = `{{.unit.Rarity}}* Lvl{{.unit.Level}} G{{.unit.Gear}} {{.unit.Name}}

Primary Attributes
- Strength: {{.s.Strength}}
- Agility: {{.s.Agility}}
- Tactics: {{.s.Tactics}}

General
- Health: {{.s.Health}} ({{.m.Health}})
- Protection: {{.s.Protection}} ({{.m.Protection}})
- Speed: {{.s.Speed}} ({{.m.Speed}})
- Critical Damage: {{perc .s.CriticalDamage}} ({{perc .m.CriticalDamage}})
- Potency: {{perc .s.Potency}} ({{perc .m.Potency}})
- Tenacity: {{perc .s.Tenacity}} ({{perc .m.Tenacity}})

Physical Offense
- Physical Damage: {{.s.PhysicalDamage}} ({{.m.PhysicalDamage}})
- Physical Critical Chance: {{perc .s.PhysicalCriticalChance}} ({{perc .m.PhysicalCriticalChance}})
- Armor Penetration: {{.s.ArmorPenetration}}
- Physical Accuracy: {{.s.PhysicalAccuracy}} ({{.m.PhysicalAccuracy}})

Physical Survivability
- Armor: {{perc .s.Armor}} ({{perc .m.Armor}})
- Dodge Chance: {{perc .s.DodgeChance}} ({{perc .m.DodgeChance}})
- Physical Critical Avoidance: {{perc .s.PhysicalCriticalAvoidance}} ({{perc .m.PhysicalCriticalAvoidance}})

Special Offense
- Special Damage: {{.s.SpecialDamage}} ({{.m.SpecialDamage}})
- Special Critical Chance: {{perc .s.SpecialCriticalChance}} ({{perc .m.SpecialCriticalChance}})
- Resistance Penetration: {{.s.ResistancePenetration}} ({{.m.ResistancePenetration}})
- Special Accuracy: {{.s.SpecialAccuracy}} ({{.m.SpecialAccuracy}})

Special Survavibility
- Resistance: {{perc .s.Resistance}} ({{perc .m.Resistance}})
- Deflection Chance: {{.s.DeflectionChance}} ({{.m.DeflectionChance}})
- Special Critical Avoidance: {{.s.SpecialCriticalAvoidance}} ({{.m.SpecialCriticalAvoidance}})

Mods
{{range .unit.Mods }}[{{.Slot}}] {{.Pips}}* Lvl{{.Level}} {{.Set}} set
- {{.Primary}}
{{range .Secondaries }}- {{.}}
{{end}}{{end}}
`
