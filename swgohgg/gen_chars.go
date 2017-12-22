package swgohgg

import "strings"

func CharSlug(charName string) string {
	switch strings.ToLower(charName) {
	case `aayla secura`:
		return "aayla-secura"
	case `admiral ackbar`:
		return "admiral-ackbar"
	case `ahsoka tano`:
		return "ahsoka-tano"
	case `ahsoka tano (fulcrum)`:
		return "ahsoka-tano-fulcrum"
	case `asajj ventress`:
		return "asajj-ventress"
	case `b2 super battle droid`:
		return "b2-super-battle-droid"
	case `barriss offee`:
		return "barriss-offee"
	case `baze malbus`:
		return "baze-malbus"
	case `bb-8`:
		return "bb-8"
	case `biggs darklighter`:
		return "biggs-darklighter"
	case `bistan`:
		return "bistan"
	case `boba fett`:
		return "boba-fett"
	case `bodhi rook`:
		return "bodhi-rook"
	case `cad bane`:
		return "cad-bane"
	case `captain han solo`:
		return "captain-han-solo"
	case `captain phasma`:
		return "captain-phasma"
	case `cassian andor`:
		return "cassian-andor"
	case `cc-2224 "cody"`:
		return "cc-2224-cody"
	case `chief chirpa`:
		return "chief-chirpa"
	case `chief nebit`:
		return "chief-nebit"
	case `chirrut îmwe`:
		return "chirrut-imwe"
	case `chopper`:
		return "chopper"
	case `clone sergeant - phase i`:
		return "clone-sergeant-phase-i"
	case `clone wars chewbacca`:
		return "clone-wars-chewbacca"
	case `colonel starck`:
		return "colonel-starck"
	case `commander luke skywalker`:
		return "commander-luke-skywalker"
	case `coruscant underworld police`:
		return "coruscant-underworld-police"
	case `count dooku`:
		return "count-dooku"
	case `ct-21-0408 "echo"`:
		return "ct-21-0408-echo"
	case `ct-5555 "fives"`:
		return "ct-5555-fives"
	case `ct-7567 "rex"`:
		return "ct-7567-rex"
	case `darth maul`:
		return "darth-maul"
	case `darth nihilus`:
		return "darth-nihilus"
	case `darth sidious`:
		return "darth-sidious"
	case `darth vader`:
		return "darth-vader"
	case `dathcha`:
		return "dathcha"
	case `death trooper`:
		return "death-trooper"
	case `dengar`:
		return "dengar"
	case `director krennic`:
		return "director-krennic"
	case `eeth koth`:
		return "eeth-koth"
	case `emperor palpatine`:
		return "emperor-palpatine"
	case `ewok elder`:
		return "ewok-elder"
	case `ewok scout`:
		return "ewok-scout"
	case `ezra bridger`:
		return "ezra-bridger"
	case `finn`:
		return "finn"
	case `first order officer`:
		return "first-order-officer"
	case `first order sf tie pilot`:
		return "first-order-sf-tie-pilot"
	case `first order stormtrooper`:
		return "first-order-stormtrooper"
	case `first order tie pilot`:
		return "first-order-tie-pilot"
	case `gamorrean guard`:
		return "gamorrean-guard"
	case `gar saxon`:
		return "gar-saxon"
	case `garazeb "zeb" orrelios`:
		return "garazeb-zeb-orrelios"
	case `general grievous`:
		return "general-grievous"
	case `general kenobi`:
		return "general-kenobi"
	case `general veers`:
		return "general-veers"
	case `geonosian soldier`:
		return "geonosian-soldier"
	case `geonosian spy`:
		return "geonosian-spy"
	case `grand admiral thrawn`:
		return "grand-admiral-thrawn"
	case `grand master yoda`:
		return "grand-master-yoda"
	case `grand moff tarkin`:
		return "grand-moff-tarkin"
	case `greedo`:
		return "greedo"
	case `han solo`:
		return "han-solo"
	case `hera syndulla`:
		return "hera-syndulla"
	case `hermit yoda`:
		return "hermit-yoda"
	case `hk-47`:
		return "hk-47"
	case `hoth rebel scout`:
		return "hoth-rebel-scout"
	case `hoth rebel soldier`:
		return "hoth-rebel-soldier"
	case `ig-100 magnaguard`:
		return "ig-100-magnaguard"
	case `ig-86 sentinel droid`:
		return "ig-86-sentinel-droid"
	case `ig-88`:
		return "ig-88"
	case `ima-gun di`:
		return "ima-gun-di"
	case `imperial probe droid`:
		return "imperial-probe-droid"
	case `imperial super commando`:
		return "imperial-super-commando"
	case `jawa`:
		return "jawa"
	case `jawa engineer`:
		return "jawa-engineer"
	case `jawa scavenger`:
		return "jawa-scavenger"
	case `jedi consular`:
		return "jedi-consular"
	case `jedi knight anakin`:
		return "jedi-knight-anakin"
	case `jedi knight guardian`:
		return "jedi-knight-guardian"
	case `jyn erso`:
		return "jyn-erso"
	case `k-2so`:
		return "k-2so"
	case `kanan jarrus`:
		return "kanan-jarrus"
	case `kit fisto`:
		return "kit-fisto"
	case `kylo ren`:
		return "kylo-ren"
	case `kylo ren (unmasked)`:
		return "kylo-ren-unmasked"
	case `lando calrissian`:
		return "lando-calrissian"
	case `lobot`:
		return "lobot"
	case `logray`:
		return "logray"
	case `luke skywalker (farmboy)`:
		return "luke-skywalker-farmboy"
	case `luminara unduli`:
		return "luminara-unduli"
	case `mace windu`:
		return "mace-windu"
	case `magmatrooper`:
		return "magmatrooper"
	case `mob enforcer`:
		return "mob-enforcer"
	case `mother talzin`:
		return "mother-talzin"
	case `nightsister acolyte`:
		return "nightsister-acolyte"
	case `nightsister initiate`:
		return "nightsister-initiate"
	case `nightsister spirit`:
		return "nightsister-spirit"
	case `nightsister zombie`:
		return "nightsister-zombie"
	case `nute gunray`:
		return "nute-gunray"
	case `obi-wan kenobi (old ben)`:
		return "obi-wan-kenobi-old-ben"
	case `old daka`:
		return "old-daka"
	case `pao`:
		return "pao"
	case `paploo`:
		return "paploo"
	case `plo koon`:
		return "plo-koon"
	case `poe dameron`:
		return "poe-dameron"
	case `poggle the lesser`:
		return "poggle-the-lesser"
	case `princess leia`:
		return "princess-leia"
	case `qui-gon jinn`:
		return "qui-gon-jinn"
	case `r2-d2`:
		return "r2-d2"
	case `rebel officer leia organa`:
		return "rebel-officer-leia-organa"
	case `resistance pilot`:
		return "resistance-pilot"
	case `resistance trooper`:
		return "resistance-trooper"
	case `rey`:
		return "rey"
	case `rey (jedi training)`:
		return "rey-jedi-training"
	case `rey (scavenger)`:
		return "rey-scavenger"
	case `royal guard`:
		return "royal-guard"
	case `sabine wren`:
		return "sabine-wren"
	case `savage opress`:
		return "savage-opress"
	case `scarif rebel pathfinder`:
		return "scarif-rebel-pathfinder"
	case `shoretrooper`:
		return "shoretrooper"
	case `sith assassin`:
		return "sith-assassin"
	case `sith trooper`:
		return "sith-trooper"
	case `snowtrooper`:
		return "snowtrooper"
	case `stormtrooper`:
		return "stormtrooper"
	case `stormtrooper han`:
		return "stormtrooper-han"
	case `sun fac`:
		return "sun-fac"
	case `talia`:
		return "talia"
	case `teebo`:
		return "teebo"
	case `tie fighter pilot`:
		return "tie-fighter-pilot"
	case `tusken raider`:
		return "tusken-raider"
	case `tusken shaman`:
		return "tusken-shaman"
	case `ugnaught`:
		return "ugnaught"
	case `urorrur'r'r`:
		return "urorrurrr"
	case `veteran smuggler chewbacca`:
		return "veteran-smuggler-chewbacca"
	case `veteran smuggler han solo`:
		return "veteran-smuggler-han-solo"
	case `wampa`:
		return "wampa"
	case `wedge antilles`:
		return "wedge-antilles"
	case `wicket`:
		return "wicket"
	case `zam wesell`:
		return "zam-wesell"
	default:
		return strings.ToLower(charName)
	}
}
