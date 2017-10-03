package swgohgg

import (
	"strings"
)

func CharName(src string) string {
	switch strings.ToLower(src) {
	case "ayla", "aayla", "aayla secura":
		return "Aayla Secura"
	case "aa", "ackbar", "admiral ackbar":
		return "Admiral Ackbar"
	case "ahsoka", "ahsoka tano", "asoka":
		return "Ahsoka Tano"
	case "ahsoka tano fulcrum", "fulcrum", "atf":
		return "Ahsoka Tano (Fulcrum)"
	case "asaj", "asajj", "ventres", "ventress", "asajj ventress":
		return "Asajj Ventress"
	case "b2", "b2 battle droid", "b2 super battle droid":
		return "B2 Super Battle Droid"
	case "barris", "baris", "barriss", "offee", "zarris", "barriss offee", "barris offee", "bariss offee", "baris offee":
		return "Barriss Offee"
	case "baze", "baze malbus":
		return "Baze Malbus"
	case "bb8", "bb-8":
		return "BB-8"
	case "biggs", "biggs darklighter":
		return "Biggs Darklighter"
	case "bistan":
		return "Bistan"
	case "bodhi", "bodhi rook":
		return "Bodhi Rook"
	case "boba", "boba fett":
		return "Boba Fett"
	case "bane", "cad", "cad bane":
		return "Cad Bane"
	case "chs", "cholo", "captain han", "captain han solo":
		return "Captain Han Solo"
	case "phasma", "captain phasma":
		return "Captain Phasma"
	case "cassian", "cassian andor":
		return "Cassian Andor"
	case "cody":
		return "CC-2224 \"Cody\""
	case "chirpa", "chief chirpa":
		return "Chief Chirpa"
	case "nebit", "chief nebit":
		return "Chief Nebit"
	case "chirrut", "chirut", "chirutt":
		return "Chirrut Îmwe"
	case "chopp", "chop", "chopper":
		return "Chopper"
	case "sarge", "sargento", "clone sergeant", "sergeant":
		return "Clone Sergeant - Phase I"
	case "chewe", "chewie", "chewbacca", "chewbaca":
		return "Clone Wars Chewbacca"
	case "cls", "commander luke", "commander luke skywalker":
		return "Commander Luke Skywalker"
	case "cup", "coruscant", "coruscant police":
		return "Coruscant Underworld Police"
	case "dooku", "dokan", "dookan", "count dooku":
		return "Count Dooku"
	case "echo":
		return "CT-21-0408 \"Echo\""
	case "fives", "5s":
		return "CT-5555 \"Fives\""
	case "rex", "zrex":
		return "CT-7567 \"Rex\""
	case "maul", "zaul", "zmaul", "darth maul":
		return "Darth Maul"
	case "dn", "nihilus", "znihilus", "darth nihilis":
		return "Darth Nihilus"
	case "sidious", "zsidious", "darth sidious", "zidious":
		return "Darth Sidious"
	case "vader", "zader", "zvader", "darth vader":
		return "Darth Vader"
	case "datcha":
		return "Dathcha"
	case "dt", "death", "death trooper":
		return "Death Trooper"
	case "dengar":
		return "dengar"
	case "dk", "director", "director krennic", "krennic":
		return "Director Krennic"
	case "eeth", "eth", "ek", "eeth koth":
		return "Eeth Koth"
	case "palpatine", "emperor", "ep", "ip", "emperor palpatine":
		return "Emperor Palpatine"
	case "ee", "elder", "anciao", "ancião", "ewok elder":
		return "Ewok Elder"
	case "es", "ewok scout":
		return "Ewok Scout"
	case "ezra", "bridger":
		return "Ezra Bridger"
	case "finn", "zinn":
		return "Finn"
	case "foo", "fo officer", "first order officer":
		return "First Order Officer"
	case "fost", "fo stormtrooper", "first order stormtrooper":
		return "First Order Stormtrooper"
	case "fotp", "fo tie", "fo pilot", "fo tie pilot", "first order pilot", "first order tie", "first order tie pilot":
		return "First Order TIE Pilot"
	case "gamorrean", "guard", "pig", "gamorrean guard", "pig guard":
		return "Gamorrean Guard"
	case "gar", "saxon", "gar saxon":
		return "Gar Saxon"
	case "zeb", "garrazeb", "garazeb", "garrazeb orrelios", "garrazeb zeb orrelios":
		return "Garrazeb \"Zeb\" Orrellios"
	case "gg", "grevous", "grievous", "grivous", "general grievous":
		return "General Grievous"
	case "kenobi", "gk", "general kenobi":
		return "General Kenobi"
	case "veers", "general veers":
		return "General Veers"
	case "sg", "gs", "geonosian", "geono", "geonosian soldier":
		return "Geonosian Soldier"
	case "spy", "gspy", "geonosian spy":
		return "Geonosian Spy"
	case "gat", "thrawn", "grand admiral", "admiral thrawn", "grand admiral thrawn":
		return "Grand Admiral Thrawn"
	case "yoda", "zoda", "gmy", "master yoda", "grand master yoda":
		return "Grand Master Yoda"
	case "gmt", "tarkin", "moff", "grand moff", "moff tarkin", "grand moff tarkin":
		return "Grand Moff Tarkin"
	case "greedo":
		return "Greedo"
	case "han", "solo", "zolo", "zsolo", "han solo", "han zolo":
		return "Han Solo"
	case "hera", "hera syndulla":
		return "Hera Syndulla"
	case "hy", "hermit", "hermit yoda":
		return "Hermit Yoda"
	case "hk", "hk47", "hk-47":
		return "HK-47"
	case "hrs", "hrscout", "hoth scout", "rebel scout", "hoth rebel scout":
		return "Hoth Rebel Scout"
	case "hrsolder", "hoth soldier", "rebel soldier", "hoth rebel soldier":
		return "Hoth Rebel Soldier"
	case "ig-100", "ig100", "ig 100", "magna", "magnaguard", "magna guard":
		return "IG-100 MagnaGuard"
	case "ig-86", "ig86", "ig 86":
		return "IG-86 Sentinel Droid"
	case "ig88", "ig-88", "ig 88":
		return "IG-88"
	case "ima", "igd", "ima gun", "gun di", "ima gun di", "ima-gun di", "ima-gun-di":
		return "Ima-Gun Di"
	case "isc", "imperial commando", "super commando", "imperial super commando":
		return "Imperial Super Commando"
	case "jawa":
		return "Jawa"
	case "je", "engineer", "jawa engineer":
		return "Jawa Engineer"
	case "js", "scavenger", "jawa scavenger":
		return "Jawa Scavenger"
	case "consul", "jc", "cj", "consular", "jedi consular":
		return "Jedi Consular"
	case "jka", "anakin", "jedi knight anakin":
		return "Jedi Knight Anakin"
	case "jkg", "jedi knight guardian":
		return "Jedi Knight Guardian"
	case "jyn", "zJyn", "zyn", "jyn erso":
		return "Jyn Erso"
	case "k2", "k2so", "k-2s0":
		return "K-2SO"
	case "kanan", "jarrus", "jarus", "kanan jarrus", "kanan jarus":
		return "Kanan Jarrus"
	case "kit", "fisto", "kit fisto":
		return "Kit Fisto"
	case "kylo", "ren", "zylo", "kylo ren":
		return "Kylo Ren"
	case "lando", "calrissian", "lando calrissian":
		return "Lando Calrissian"
	case "lobot":
		return "Lobot"
	case "logray":
		return "Logray"
	case "luke", "farmboy", "farmboy luke", "luke skywalker":
		return "Luke Skywalker (Farmboy)"
	case "lumi", "luminara", "unduli", "luminara unduli":
		return "Luminara Unduli"
	case "mace", "windu", "mace windu":
		return "Mace Windu"
	case "magma", "magmatrooper", "magma trooper":
		return "Magmatrooper"
	case "mob", "enforcer", "mob enforcer":
		return "Mob Enforcer"
	case "acolyte", "ns acolyte", "nsa", "night sister acolyte", "nightsister acolyte":
		return "Nightsister Acolyte"
	case "initiate", "ns initiate", "nsi", "night sister initiate", "nightsister initiate":
		return "Nightsister Initiate"
	case "nute", "gunray", "nute gunray":
		return "Nute Gunray"
	case "old ben", "obi", "obi wan", "obi wan kenobi":
		return "Obi-Wan Kenobi (Old Ben)"
	case "daka":
		return "Old Daka"
	case "plo", "koon":
		return "Plo Koon"
	case "poe":
		return "Poe Dameron"
	case "poggle", "pogle", "pogle the lesser":
		return "Poggle the Lesser"
	case "leia", "zleia", "léia":
		return "Princess Leia"
	case "qgj", "quigon", "qui-gon", "qui gon jin", "qui-gon-jin":
		return "Qui-Gon Jin"
	case "r2d2", "r2":
		return "R2-D2"
	case "rp":
		return "Resistance Pilot"
	case "rt":
		return "Resistance Trooper"
	case "rolo", "rebel leia":
		return "Rebel Office Leia Organa"
	case "rey":
		return "Rey"
	case "rg", "royal":
		return "Royal Guard"
	case "sabine", "wren":
		return "Sabine Wren"
	case "savage", "so", "zavage":
		return "Savage Opress"
	case "scarif", "srp", "pathfinder":
		return "Scarif Rebel Pathfinder"
	case "shore":
		return "Shoretrooper"
	case "sass", "assassin", "sassassin":
		return "Sith Assassin"
	case "strooper", "trooper":
		return "Sith Trooper"
	case "snow":
		return "Snowtrooper"
	case "st", "storm":
		return "Stormtrooper"
	case "sth", "sthan":
		return "Stormtrooper Han"
	case "sf":
		return "Sun Fac"
	case "talia":
		return "Talia"
	case "teebo":
		return "teebo"
	case "tfp", "tie", "tie pilot", "tie fighter":
		return "TIE Fighter Pilot"
	case "tusken", "Raider":
		return "Tusken Raider"
	case "shaman":
		return "Tusken Shaman"
	case "ug", "ugg":
		return "Ugnaught"
	case "uror", "urorr":
		return "URoRRuR'R'R"
	case "vhan", "vet han", "veteran han", "veteran han solo", "smuggler han":
		return "Veteran Smuggler Chewbacca"
	case "vchewie", "vchewbacca", "vet chewie", "vet chewbacca", "veteran chewie", "veteran chewbacca", "smuggler chewie", "smuggler chewbacca":
		return "Veteran Smuggler Han Solo"
	case "wedge", "wedge antilles":
		return "Wedge Antilles"
	case "wicket":
		return "Wicket"
	case "zam", "zw", "zam wesell":
		return "Zam Wesell"
	}
	return src
}
