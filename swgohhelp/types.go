package swgohhelp

// AuthResponse represents the authentication response data.
type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Player represents the full player profile information.
type Player struct {
	Name       string `json:"name"`
	AllyCode   int    `json:"allycode"`
	Level      int    `json:"level"`
	GuildName  string `json:"guildName"`
	GuildRefID string `json:"guildRefId"`

	UpdateTimestamp int `json:"updated"`

	// Stats ProfileStats `json:"stats"`
	// Roster []Units      `json:"roster"`
	Arena Arena `json:"arena"`
}

// Arena wraps both arena rankings for the player.
type Arena struct {
	Char ArenaRanking `json:"char"`
	Ship ArenaRanking `json:"ship"`
}

// ArenaRanking holds player arena ranking.
type ArenaRanking struct {
	Rank int `json:"rank"`

	Squad []SquadUnit `json:"char"`
}

// SquadUnit represents an arena squad unit identifier set.
type SquadUnit struct {
	ID     string `json:"id"`
	UnitID string `json:"defId"`
	Type   string `json:"type"`
}
