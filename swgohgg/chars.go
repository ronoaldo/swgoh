package swgohgg

import (
	"strings"
)

func CharName(src string) string {
	switch strings.ToLower(src) {
	case "aa", "ackbar":
		return "Admiral Ackbar"
	case "ayla", "aayla":
		return "Aayla Secura"
	case "ahsoka", "asoka", "ahsoka tano":
		return "Ahsoka Tano"
	case "asaj", "asajj", "ventres", "ventress", "asajj ventress":
		return "Asajj Ventress"
	case "b2", "b2 battle droid", "b2 super battle droid":
		return "B2 Super Battle Droid"
	case "barris", "baris", "barriss", "offee", "zarris", "barriss offee":
		return "Barriss Offee"
	case "baze", "baze malbus":
		return "Baze Malbus"
	case "biggs", "biggs darklighter":
		return "Biggs Darklighter"
	case "boba", "boba fett":
		return "Boba Fett"
	case "bane", "cad", "cad bane":
		return "Cad Bane"
	case "cholo", "captain han solo":
		return "Captain Han Solo"
	case "phasma", "captain phasma":
		return "Captain Phasma"
	case "cassian", "cassian andor":
		return "Cassian Andor"
	case "cls", "commander luke", "commander luke skywalker":
		return "Commander Luke Skywalker"
	case "cody":
		return "CT-2224 \"Cody\""
	case "chirpa", "chief chirpa":
		return "Chief Chirpa"
	case "nebit", "chief nebit":
		return "Chief Nebit"
	case "chirrut":
		return "Chirrut Îmwe"
	case "sarge", "sargento":
		return "Clone Sergeant - Phase I"
	case "chewe", "chewbacca", "chewbaca":
		return "Clone Wars Chewbacca"
	case "cup", "coruscant":
		return "Coruscant Underworld Police"
	case "dooku", "dokan", "dookan", "count dooku":
		return "Count Dooku"
	case "echo":
		return "CT-21-0408 \"Echo\""
	case "fives", "5s":
		return "CT-5555 \"Fives\""
	case "rex":
		return "CT-7567 \"Rex\""
	case "maul", "darth maul":
		return "Darth Maul"
	case "sidious", "sid", "darth sidious":
		return "Darth Sidious"
	case "vader", "darth vader":
		return "Darth Vader"
	case "datcha":
		return "Dathcha"
	case "dt", "death", "death trooper", "deathtrooper":
		return "Death Trooper"
	case "dk", "krennic", "krenic", "director krenic", "director krennic":
		return "Director Krennic"
	case "eeth", "eth", "ek":
		return "Eeth Koth"
	case "palpatine", "emperor", "ep", "ip", "emperor palpatine":
		return "Emperor Palpatine"
	case "ee", "elder", "anciao", "ancião", "ewok elder":
		return "Ewok Elder"
	case "es", "ewok scout":
		return "Ewok Scout"
	case "foo":
		return "First Order Officer"
	case "fost":
		return "First Order Stormtrooper"
	case "fotp":
		return "First Order TIE Pilot"
	case "gar", "saxon", "gar saxon":
		return "Gar Saxon"
	case "gamorrean", "pig":
		return "Gamorrean Guard"
	case "gg", "grevous", "grievous", "grivous":
		return "General Grievous"
	case "kenobi", "gk", "general kenobi":
		return "General Kenobi"
	case "veers":
		return "General Veers"
	case "sg", "gs", "geonosian", "geono":
		return "Geonosian Soldier"
	case "spy", "gspy":
		return "Geonosian Spy"
	case "gat", "thrawn":
		return "Grand Admiral Thrawn"
	case "yoda", "gmy":
		return "Grand Master Yoda"
	case "tarkin", "moff":
		return "Grand Moff Tarkin"
	case "han", "solo":
		return "Han Solo"
	case "hrscout":
		return "Hoth Rebel Scount"
	case "hrsolder":
		return "Hoth Rebel Soldier"
	case "ig-100", "ig100", "ig 100":
		return "IG-100 MagnaGuard"
	case "ig-86", "ig86", "ig 86":
		return "IG-86 Sentinel Droid"
	case "ig88", "ig-88", "ig 88":
		return "IG-88"
	case "ima", "igd":
		return "Ima-Gun Di"
	case "isc":
		return "Imperial Super Commando"
	case "je", "engineer":
		return "Jawa Engineer"
	case "scavenger":
		return "Jawa Scavenger"
	case "consul", "jc", "cj":
		return "Jedi Consular"
	case "jka", "anakin":
		return "Jedi Knight Anakin"
	case "jkg":
		return "Jedi Knight Guardian"
	case "jyn":
		return "Jyn Erso"
	case "k2", "k2so":
		return "K-2SO"
	case "kit", "fisto":
		return "Kit Fisto"
	case "kylo", "ren":
		return "Kylo Ren"
	case "lando":
		return "Lando Calrissian"
	case "luke", "luke skywalker":
		return "Luke Skywalker (Farmboy)"
	case "lumi", "luminara":
		return "Luminara Unduli"
	case "mace", "windu":
		return "Mace Windu"
	case "magma":
		return "Magmatrooper"
	case "mob":
		return "Mob Enforcer"
	case "acolyte":
		return "Nightsister Acolyte"
	case "initiate":
		return "Nightsister Initiate"
	case "nute":
		return "Nute Gunray"
	case "old ben", "obi":
		return "Obi-Wan Kenobi (Old Ben)"
	case "daka":
		return "Old Daka"
	case "plo":
		return "Plo Koon"
	case "poe":
		return "Poe Dameron"
	case "poggle", "pogle":
		return "Poggle the Lesser"
	case "leia", "léia":
		return "Princess Leia"
	case "qgj", "quigon", "qui-gon":
		return "Qui-Gon Jin"
	case "r2d2", "r2":
		return "R2-D2"
	case "rp":
		return "Resistance Pilot"
	case "rt":
		return "Resistance Trooper"
	case "rg", "royal":
		return "Royal Guard"
	case "savage", "so":
		return "Savage Opress"
	case "scarif", "srp":
		return "Scarif Rebel Pathfinder"
	case "shore":
		return "Shoretrooper"
	case "snow":
		return "Snowtrooper"
	case "st":
		return "Stormtrooper"
	case "sthan", "stormtrooper han":
		return "Stormtrooper Han"
	case "sf":
		return "Sun Fac"
	case "tfp", "tie":
		return "TIE Fighter Pilot"
	case "tusken":
		return "Tusken Rider"
	case "shaman":
		return "Tusken Shaman"
	case "uror":
		return "URoRRuR'R'R"
	case "wedge":
		return "Wedge Antilles"
	case "zam", "zw":
		return "Zam Wesell"
	}
	return src
}
