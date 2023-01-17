package swgohhelp

type Guild struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Desc        string   `json:"desc"`
	Members     int64    `json:"members"`
	Status      int64    `json:"status"`
	Required    int64    `json:"required"`
	BannerColor string   `json:"bannerColor"`
	BannerLogo  string   `json:"bannerLogo"`
	Message     string   `json:"message"`
	Gp          int64    `json:"gp"`
	Raid        Raid     `json:"raid"`
	Roster      []Roster `json:"roster"`
	Updated     int64    `json:"updated"`
}

type Raid struct {
	Rancor          string `json:"rancor"`
	Aat             string `json:"aat"`
	SithRAID        string `json:"sith_raid"`
	RancorChallenge string `json:"rancor_challenge"`
}

type Roster struct {
	ID               string `json:"id"`
	GuildMemberLevel int64  `json:"guildMemberLevel"`
	Name             string `json:"name"`
	Level            int64  `json:"level"`
	AllyCode         int64  `json:"allyCode"`
	Gp               int64  `json:"gp"`
	GpChar           int64  `json:"gpChar"`
	GpShip           int64  `json:"gpShip"`
	Updated          int64  `json:"updated"`
}
